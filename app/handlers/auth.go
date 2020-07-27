package handlers

import (
	"github.com/badThug/otus-social-network/app/handlers/requests"
	"github.com/badThug/otus-social-network/app/storage"
	"log"
	"math/rand"
)

func (h *Handler) ViewLoginHandler() error {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))

	return nil
}

func (h *Handler) ViewRegisterHandler() error {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))

	return nil
}

func (h *Handler) ApiLoginHandler() error {
	var request *requests.RegisterRequest
	if err := h.decodeJson(&request); nil != err {
		return err
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	userId := rand.Intn(1000000)

	h.session.Values[storage.SessionUserIdKey] = userId
	h.session.Save(h.request, h.writer)

	log.Printf("userId = %d", userId)

	h.success("Login Success!")

	return nil
}

func (h *Handler) ApiLogoutHandler() error {
	h.session.Values[storage.SessionUserIdKey] = nil
	h.session.Save(h.request, h.writer)

	h.success("Logout Success!")

	return nil
}

func (h *Handler) ApiRegisterHandler() error {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))

	return nil
}
