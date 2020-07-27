package handlers

import (
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/components/utils"
	"github.com/gorilla/sessions"
	"net/http"
)

type Handler struct {
	db             *storage.DbConnection
	sessionStorage storage.SessionStorage
	session        *sessions.Session
	writer         http.ResponseWriter
	request        *http.Request
}

func InitHandler(db *storage.DbConnection, session storage.SessionStorage) *Handler {
	return &Handler{db: db, sessionStorage: session}
}

func (h *Handler) InitHandle(w http.ResponseWriter, r *http.Request) {
	h.writer = w
	h.request = r
	h.session = h.sessionStorage.GetSession(r)
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

func (h *Handler) success(data interface{}) {
	var message interface{}

	switch v := data.(type) {
	case string:
		message = utils.ResponseMessage(true, v)
	default:
		message = utils.ResponseData(true, v)
	}

	utils.SendResponseJson(h.writer, message)
}

func (h *Handler) error(data interface{}) {
	var message interface{}

	switch v := data.(type) {
	case string:
		message = utils.ResponseMessage(false, v)
	default:
		message = utils.ResponseData(false, v)
	}

	utils.SendResponseJsonWithStatusCode(h.writer, message, http.StatusBadRequest)
}
