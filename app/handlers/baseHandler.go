package handlers

import "github.com/badThug/otus-social-network/app/components/storage"

type Handler struct {
	db             *storage.Connection
	sessionStorage storage.SessionStorage
}

func InitHandler(db *storage.Connection, session storage.SessionStorage) *Handler {
	return &Handler{db: db, sessionStorage: session}
}
