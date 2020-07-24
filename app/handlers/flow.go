package handlers

import (
	"net/http"
)

func (h *Handler) ViewFlowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from Flow!</h1>"))
}
