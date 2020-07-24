package handlers

import (
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/gorilla/sessions"
	"net/http"
)

type Handler struct {
	db             *storage.Connection
	sessionStorage storage.SessionStorage
	session        *sessions.Session
	writer         http.ResponseWriter
	request        *http.Request
}

func InitHandler(db *storage.Connection, session storage.SessionStorage) *Handler {
	return &Handler{db: db, sessionStorage: session}
}

func (h *Handler) InitHandle(w http.ResponseWriter, r *http.Request) {
	h.writer = w
	h.request = r
	h.session = h.sessionStorage.GetSession(r)
}
