package web

import (
	"github.com/badThug/otus-social-network/app/handlers"
	"github.com/badThug/otus-social-network/app/params"
	"net/http"
)

func initRoutes(d Dispatcher) {
	d.Get(params.ViewIndexRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ViewIndexHandler()
	}))
	d.Get(params.ViewFlowRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ViewFlowHandler()
	}))
	d.Get(params.ViewLoginRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ViewLoginHandler()
	}))
	d.Get(params.ViewRegisterRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ViewRegisterHandler()
	}))

	d.Post(params.ApiFriendRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ApiAddFriendHandler()
	}))
	d.Post(params.ApiLoginRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ApiLoginHandler()
	}))
	d.Post(params.ApiRegisterRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ApiRegisterHandler()
	}))
	d.Post(params.ApiLogoutRoute, d.handleRequest(func(h *handlers.Handler) {
		h.ApiLogoutHandler()
	}))

	d.Router.Use(SessionAuthentication)
}

func (d *Dispatcher) handleRequest(handlerMethod func(h *handlers.Handler)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := handlers.InitHandler(d.db, d.SessionStorage)
		h.InitHandle(w, r)
		handlerMethod(h)
	}
}
