package utils

import (
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendResponseJson(w, ResponseMessage(false, "This resources was not found on our server"))
		w.WriteHeader(http.StatusNotFound)
		next.ServeHTTP(w, r)
	})
}
