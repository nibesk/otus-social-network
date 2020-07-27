package web

import (
	"github.com/badThug/otus-social-network/app/components/config"
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/components/utils"
	"github.com/badThug/otus-social-network/app/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type Dispatcher struct {
	Router         *mux.Router
	db             *storage.DbConnection
	SessionStorage storage.SessionStorage
}

var sessionStorage storage.SessionStorage

func InitDispatcher(db *storage.DbConnection, config *config.Config) Dispatcher {
	router := mux.NewRouter()
	sessionStorage = storage.InitSession(config)

	dispatcher := Dispatcher{
		Router:         router,
		SessionStorage: sessionStorage,
		db:             db,
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

func (d *Dispatcher) handleRequest(handlerMethod func(h *handlers.Handler) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d.db.Connect()
		defer d.db.Close()

		h := handlers.InitHandler(d.db, d.SessionStorage)
		h.InitHandle(w, r)

		var malformedRequest *utils.MalformedRequest

		if err := handlerMethod(h); nil != err {
			switch {
			case errors.As(err, &malformedRequest):
				if malformed, ok := err.(*utils.MalformedRequest); ok {
					h.ResponseWithError(malformed.Msg, malformed.Status)
				}
			default:
				h.ResponseWithError("Internal server error", http.StatusInternalServerError)
			}

			log.Printf("Error occured: %w", err)
		}
	}
}
