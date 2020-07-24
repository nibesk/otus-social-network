package web

import (
	"github.com/badThug/otus-social-network/app/components/config"
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Dispatcher struct {
	Router         *mux.Router
	db             *storage.Connection
	SessionStorage storage.SessionStorage
}

var sessionStorage storage.SessionStorage

func InitDispatcher(db *storage.Connection, config *config.Config) Dispatcher {
	router := mux.NewRouter()
	sessionStorage = storage.InitSession(config)

	dispatcher := Dispatcher{
		Router:         router,
		SessionStorage: sessionStorage,
	}

	initRoutes(dispatcher)

	return dispatcher
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

func IsJsonRequest(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")

	return "application/json" == contentType
}
