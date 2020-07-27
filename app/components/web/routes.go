package web

import (
	"github.com/badThug/otus-social-network/app/handlers"
	"github.com/badThug/otus-social-network/app/params"
)

func initRoutes(d Dispatcher) {
	d.Get(params.ViewIndexRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewIndexHandler()
	}))
	d.Get(params.ViewFlowRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewFlowHandler()
	}))
	d.Get(params.ViewLoginRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewLoginHandler()
	}))
	d.Get(params.ViewRegisterRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ViewRegisterHandler()
	}))

	d.Post(params.ApiFriendRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiAddFriendHandler()
	}))
	d.Post(params.ApiLoginRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiLoginHandler()
	}))
	d.Post(params.ApiRegisterRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiRegisterHandler()
	}))
	d.Post(params.ApiLogoutRoute, d.handleRequest(func(h *handlers.Handler) error {
		return h.ApiLogoutHandler()
	}))

	d.Router.Use(SessionAuthentication)
}
