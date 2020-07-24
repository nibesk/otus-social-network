package handlers

import (
	"github.com/badThug/otus-social-network/app/components/storage"
	"github.com/badThug/otus-social-network/app/components/utils"
	"log"
	"math/rand"
)

func (h *Handler) ViewLoginHandler() {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))
}

func (h *Handler) ViewRegisterHandler() {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))
}

func (h *Handler) ApiLoginHandler() {
	userId := rand.Intn(1000000)

	h.session.Values[storage.SessionUserIdKey] = userId
	h.session.Save(h.request, h.writer)

	log.Printf("userId = %d", userId)

	utils.SendResponseJson(h.writer, utils.ResponseMessage(true, "Login Success!"))
}

func (h *Handler) ApiLogoutHandler() {
	h.session.Values[storage.SessionUserIdKey] = nil
	h.session.Save(h.request, h.writer)

	utils.SendResponseJson(h.writer, utils.ResponseMessage(false, "Login Success!"))
}

func (h *Handler) ApiRegisterHandler() {
	h.writer.Write([]byte("<h1>Hello from Login!</h1>"))
}
