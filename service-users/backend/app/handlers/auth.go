package handlers

import (
	"database/sql"
	"github.com/pkg/errors"
	"service-users/app/handlers/requests"
	"service-users/app/models"
	"service-users/app/storage"
	"service-users/app/utils"
)

func (h *Handler) ApiLoginHandler() error {
	var request *requests.LoginRequest
	if err := h.decodeRequest(&request); nil != err {
		return errors.WithStack(err)
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
		"user": user,
	})
}

func (h *Handler) ApiLogoutHandler() error {
	h.session.Values[storage.SessionUserIdKey] = nil
	h.session.Save(h.request, h.writer)

	return h.success("Logout Success!")
}

func (h *Handler) ApiGetUserHandler() error {
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
	if err := h.decodeRequest(&request); nil != err {
		return errors.WithStack(err)
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	_, err := h.getSessionUserId()
	if nil == err {
		return h.error("You are already logged in")
	}

	hashedPwd, err := utils.HashPassword(request.Password)
	if nil != err {
		return err
	}

	user := &models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashedPwd,
		Surname:   request.Surname,
		Age:       request.Age,
		City:      request.City,
		Interests: request.Interests,
		Sex:       request.Sex,
	}

	_, err = models.UserCreate(h.db, user)
	if nil != err {
		return err
	}

	h.session.Values[storage.SessionUserIdKey] = user.User_id
	h.session.Save(h.request, h.writer)

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}
