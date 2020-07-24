package handlers

import (
	"fmt"
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/components/utils"
	"math/rand"
	"net/http"
)

func (h *Handler) ViewLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from Login!</h1>"))
}

func (h *Handler) ViewRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from Login!</h1>"))
}

func (h *Handler) ApiLoginHandler(w http.ResponseWriter, r *http.Request) {
	session := h.sessionStorage.GetSession(r)
	session.Values[storage.SessionUserIdKey] = rand.Intn(1000000)
	session.Save(r, w)

	// Print secret message
	fmt.Fprintln(w, fmt.Sprintf("userId = %d", session.Values["userId"].(int)))

	utils.SendResponseJson(w, utils.ResponseMessage(false, "Login Success!"))
}

func (h *Handler) ApiLogoutHandler(w http.ResponseWriter, r *http.Request) {
	session := h.sessionStorage.GetSession(r)
	session.Values[storage.SessionUserIdKey] = nil
	session.Save(r, w)

	// Print secret message
	fmt.Fprintln(w, fmt.Sprintf("userId = %d", session.Values["userId"].(int)))

	utils.SendResponseJson(w, utils.ResponseMessage(false, "Login Success!"))
}

func (h *Handler) ApiRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from Login!</h1>"))
}
