package web

import (
	"github.com/badThug/otus-social-network/app/components/config"
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/handlers"
	"github.com/badThug/otus-social-network/app/params"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Dispatcher struct {
	Router         *mux.Router
	Handler        *handlers.Handler
	SessionStorage storage.SessionStorage
}

var sessionStorage storage.SessionStorage

func InitDispatcher(db *storage.Connection, config *config.Config) Dispatcher {
	router := mux.NewRouter()
	sessionStorage = storage.InitSession(config)
	handler := handlers.InitHandler(db, sessionStorage)

	dispatcher := Dispatcher{
		Router:         router,
		Handler:        handler,
		SessionStorage: sessionStorage,
	}
	dispatcher.initRoutes()

	return dispatcher
}

func (d *Dispatcher) initRoutes() {
	d.Get(params.ViewIndexRoute, d.handleRequest(d.Handler.ViewIndexHandler))
	d.Get(params.ViewFlowRoute, d.handleRequest(d.Handler.ViewFlowHandler))
	d.Get(params.ViewLoginRoute, d.handleRequest(d.Handler.ViewLoginHandler))
	d.Get(params.ViewRegisterRoute, d.handleRequest(d.Handler.ViewRegisterHandler))

	d.Post(params.ApiFriendRoute, d.handleRequest(d.Handler.ApiAddFriendHandler))
	d.Post(params.ApiLoginRoute, d.handleRequest(d.Handler.ApiLoginHandler))
	d.Post(params.ApiRegisterRoute, d.handleRequest(d.Handler.ApiRegisterHandler))
	d.Post(params.ApiLogoutRoute, d.handleRequest(d.Handler.ApiLogoutHandler))

	d.Router.Use(SessionAuthentication)
}

// Get wraps the router for GET method
func (d *Dispatcher) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	d.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (d *Dispatcher) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	d.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (d *Dispatcher) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	d.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (d *Dispatcher) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	d.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (d *Dispatcher) Run(host string) {
	log.Fatal(http.ListenAndServe(host, d.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (d *Dispatcher) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func IsJsonRequest(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")

	return "application/json" == contentType
}
