package web

import (
	"service-chat/app/globals"
	"service-chat/app/handlers"
)

func initRoutes(d Dispatcher) {
	d.Get(globals.ViewIndexRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewIndexHandler()
	}))

	d.Get(globals.WebSocketChatRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ConnectToSocket()
	}))

	d.Get(globals.MessagesRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.GetMessages()
	}))
}
