package handlers

import (
	"net/http"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from index!</h1>"))
}
