package handlers

import (
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"service-chat/app/utils"
)

type Handler struct {
	writer          http.ResponseWriter
	request         *http.Request
	isBlockResponse bool
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func CreateValidator() {
	validate = validator.New()
}

func InitHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{
		writer:          w,
		request:         r,
		isBlockResponse: false,
	}
}

func (h *Handler) disableResponse() {
	h.isBlockResponse = true
}

func (h *Handler) ResponseWithError(msg string, statusCode int) {
	if h.isBlockResponse {
		return
	}

	if utils.IsJsonRequest(h.request) {
		utils.SendResponseJsonWithStatusCode(h.writer, utils.ResponseMessage(false, msg), statusCode)
	} else {
		http.Error(h.writer, msg, statusCode)
	}
}

func (h *Handler) decodeRequest(obj interface{}) error {
	return utils.DecodeJSONBody(&obj, h.writer, h.request)
}

func (h *Handler) success(data interface{}) error {
	if h.isBlockResponse {
		return nil
	}

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
	if h.isBlockResponse {
		return nil
	}

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
