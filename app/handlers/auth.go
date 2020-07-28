package handlers

import (
	"database/sql"
	"github.com/badThug/otus-social-network/app/handlers/requests"
	"github.com/badThug/otus-social-network/app/models"
	"github.com/badThug/otus-social-network/app/storage"
	"github.com/badThug/otus-social-network/app/utils"
	"github.com/pkg/errors"
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
	var request *requests.LoginRequest
	if err := h.decodeJson(&request); nil != err {
		return err
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	user, err := models.UserFindByEmail(h.db, request.Email)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error("Password or email incorrect")
	case nil != err:
		return err
	}

	if err := utils.CheckPassword(request.Password, user.Password); nil != err {
		return h.error("Password or email incorrect")
	}

	h.session.Values[storage.SessionUserIdKey] = user.User_id
	h.session.Save(h.request, h.writer)

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}

func (h *Handler) ApiLogoutHandler() error {
	h.session.Values[storage.SessionUserIdKey] = nil
	h.session.Save(h.request, h.writer)

	return h.success("Logout Success!")
}

func (h *Handler) ApiIsAuthHandler() error {
	userId, ok := h.session.Values[storage.SessionUserIdKey].(int)
	if !ok {
		return h.success(map[string]interface{}{
			"user": nil,
		})
	}

	user, err := models.UserFindById(h.db, userId)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		h.session.Values[storage.SessionUserIdKey] = nil
		h.session.Save(h.request, h.writer)

		return h.success(map[string]interface{}{
			"user": nil,
		})
	case nil != err:
		return err
	}

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}

func (h *Handler) ApiRegisterHandler() error {
	var request *requests.RegisterRequest
	if err := h.decodeJson(&request); nil != err {
		return err
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	_, ok := h.session.Values[storage.SessionUserIdKey].(int)
	if ok {
		return h.error("You are already logged in")
	}

	hashedPwd, err := utils.HashPassword(request.Password)
	if nil != err {
		return err
	}

	user, err := models.UserCreate(h.db, request.Name, request.Email, hashedPwd)
	if nil != err {
		return err
	}

	h.session.Values[storage.SessionUserIdKey] = user.User_id
	h.session.Save(h.request, h.writer)

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}
