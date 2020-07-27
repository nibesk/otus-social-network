package web

import (
	"github.com/badThug/otus-social-network/app/globals"
	"github.com/badThug/otus-social-network/app/handlers"
)

func initRoutes(d Dispatcher) {
	d.Get(globals.ViewIndexRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewIndexHandler()
	}))
	d.Get(globals.ViewFlowRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewFlowHandler()
	}))
	d.Get(globals.ViewLoginRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewLoginHandler()
	}))
	d.Get(globals.ViewRegisterRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewRegisterHandler()
	}))

	d.Post(globals.ApiFriendRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiAddFriendHandler()
	}))
	d.Post(globals.ApiLoginRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiLoginHandler()
	}))
	d.Post(globals.ApiRegisterRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiRegisterHandler()
	}))
	d.Post(globals.ApiLogoutRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiLogoutHandler()
	}))

	d.Router.Use(SessionAuthentication)
}
