package handlers

import (
	"fmt"
	"github.com/badThug/otus-social-network/app/storage"
	"github.com/badThug/otus-social-network/app/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"net/http"
)

type Handler struct {
	db      *storage.DbConnection
	session *sessions.Session
	writer  http.ResponseWriter
	request *http.Request
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func CreateValidator() {
	validate = validator.New()
}

func InitHandler(db *storage.DbConnection, sessionStorage storage.SessionStorage, w http.ResponseWriter, r *http.Request) *Handler {
	session := sessionStorage.GetSession(r)

	return &Handler{
		db:      db,
		writer:  w,
		request: r,
		session: session}
}

func (h *Handler) ResponseWithError(msg string, statusCode int) {
	if utils.IsJsonRequest(h.request) {
		utils.SendResponseJsonWithStatusCode(h.writer, utils.ResponseMessage(false, msg), statusCode)
	} else {
		http.Error(h.writer, msg, statusCode)
	}
}

func (h *Handler) decodeJson(obj interface{}) error {
	return utils.DecodeJSONBody(&obj, h.writer, h.request)
}

func (h *Handler) success(data interface{}) error {
	var message interface{}

	switch v := data.(type) {
	case string:
		message = utils.ResponseMessage(true, v)
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
		message = utils.ResponseMessage(false, v)
	default:
		message = utils.ResponseErrors(v)
	}

	utils.SendResponseJsonWithStatusCode(h.writer, message, http.StatusBadRequest)

	return nil
}

func (h *Handler) checkValidations(s interface{}) []string {
	err := validate.Struct(s)

	if err != nil {
		violations := make([]string, 0, 5)
		for _, err := range err.(validator.ValidationErrors) {
			violations = append(violations, fmt.Sprintf("Field validation for '%s' failed. %s", err.Field(), err.Tag()))
		}

		return violations
	}

	return nil
}
