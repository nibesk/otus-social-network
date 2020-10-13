package handlers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"net/http"
	"service-users/app/globals"
	"service-users/app/models"
	"service-users/app/storage"
	"service-users/app/utils"
)

type Handler struct {
	db      *storage.DbConnection
	writer  http.ResponseWriter
	request *http.Request
}

func InitHandler(db *storage.DbConnection, w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{
		db:      db,
		writer:  w,
		request: r,
	}
}

func (h *Handler) ResponseWithError(msg string, statusCode int) {
	if utils.IsJsonRequest(h.request) {
		utils.SendResponseJsonWithStatusCode(h.writer, utils.ResponseErrorMessage(msg), statusCode)
	} else {
		http.Error(h.writer, msg, statusCode)
	}
}

func (h *Handler) decodeRequest(obj interface{}) error {
	return utils.DecodeJSONBody(&obj, h.writer, h.request)
}

func (h *Handler) success(data interface{}) error {
	var message interface{}

	switch v := data.(type) {
	case string:
		message = utils.ResponseSuccessMessage(v)
	default:
		message = utils.ResponseData(v)
	}

	utils.SendResponseJson(h.writer, message)

	return nil
}

func (h *Handler) error(data interface{}) error {
	var message interface{}

	switch v := data.(type) {
	case string:
		message = utils.ResponseErrorMessage(v)
	default:
		message = utils.ResponseErrors(v)
	}

	utils.SendResponseJsonWithStatusCode(h.writer, message, http.StatusBadRequest)

	return nil
}

func (h *Handler) checkValidations(s interface{}) []string {
	err := utils.GetValidator().Struct(s)

	if err != nil {
		violations := make([]string, 0, 5)
		for _, err := range err.(validator.ValidationErrors) {
			violations = append(violations, fmt.Sprintf("Field validation for '%s' failed. %s", err.Field(), err.Tag()))
		}

		return violations
	}

	return nil
}

func (h *Handler) getAuthUserId() (int, bool) {
	userId, ok := h.request.Context().Value(globals.AuthUserIdKey).(uint)

	return int(userId), ok
}

var ErrNoAuthUserId = errors.New("UserId doesn't persist in account")

func (h *Handler) getAuthUser() (*models.User, error) {
	userId, ok := h.request.Context().Value(globals.AuthUserIdKey).(uint)
	if !ok {
		return nil, ErrNoAuthUserId
	}

	user, err := models.UserFindById(h.db.GetDb(), int(userId))
	if nil != err {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
