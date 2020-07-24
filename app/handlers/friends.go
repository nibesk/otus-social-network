package handlers

import (
	"net/http"
)

func (h *Handler) ApiAddFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from ApiAddFriendHandler!</h1>"))
}
