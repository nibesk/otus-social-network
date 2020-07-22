package dispatcher

import (
	"github.com/badThug/otus-social-network/app/database"
	"github.com/badThug/otus-social-network/app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// App has router and db instances
type Dispatcher struct {
	Router  *mux.Router
	Handler *handlers.Handler
}

func Init(db *database.Connection) *Dispatcher {
	router := mux.NewRouter()
	handler := &handlers.Handler{Db: db}
	dispatcher := &Dispatcher{Router: router, Handler: handler}
	dispatcher.initRoutes()

	return dispatcher
}

func (d *Dispatcher) initRoutes() {
	d.Get("/", d.handleRequest(d.Handler.IndexHandler))
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
