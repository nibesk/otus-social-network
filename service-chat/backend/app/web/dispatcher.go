package web

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"runtime/debug"
	"service-chat/app/customErrors"
	"service-chat/app/handlers"
	"service-chat/app/utils"
)

type Dispatcher struct {
	Router *mux.Router
}

func InitDispatcher() Dispatcher {
	router := mux.NewRouter()
	dispatcher := Dispatcher{Router: router}
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

func (d *Dispatcher) handleRequest(handlerMethod func(h *handlers.Handler) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := handlers.InitHandler(w, r)
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC] %w; Stack trace %s", r, string(debug.Stack()))
				h.ResponseWithError("Internal server error", http.StatusInternalServerError)
			}
		}()

		log.Printf("%s %s", r.Method, r.URL)

		err := handlerMethod(h)
		if nil != err {
			switch causedErr := errors.Cause(err).(type) {
			case *utils.MalformedRequest:
				h.ResponseWithError(causedErr.Msg, causedErr.Status)
			case *customErrors.TypedError:
				h.ResponseWithError(causedErr.Msg, http.StatusBadRequest)
			case *customErrors.TypedStatusError:
				h.ResponseWithError(causedErr.Msg, causedErr.Status)
			default:
				h.ResponseWithError("Internal server error", http.StatusInternalServerError)
			}

			log.Printf("[ERROR] %+v\n", err)
		}
	}
}
