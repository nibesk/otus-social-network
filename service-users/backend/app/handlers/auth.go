package handlers

import (
	"database/sql"
	"github.com/pkg/errors"
	"service-users/app/handlers/requests"
	"service-users/app/models"
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

	user, err := models.UserFindByEmail(h.db.GetDb(), request.Email)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error("Password or email incorrect")
	case nil != err:
		return err
	}

	if err := utils.CheckPassword(request.Password, user.Password); nil != err {
		return h.error("Password or email incorrect")
	}

	if "" == user.Token.String {
		token := utils.CreateJWTToken(uint(user.User_id))
		err = user.UpdateToken(h.db.GetDb(), token)
		if nil != err {
			return errors.Wrapf(err, "Can't update token")
		}
	}

	return h.success(map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) ApiLogoutHandler() error {
	userId, _ := h.getAuthUserId()

	user, err := models.UserFindById(h.db.GetDb(), userId)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.success(map[string]interface{}{"user": nil})
	case nil != err:
		return err
	}

	err = user.UpdateToken(h.db.GetDb(), nil)
	if nil != err {
		return errors.Wrapf(err, "Can't update token")
	}

	return h.success("Logout Success!")
}

func (h *Handler) ApiGetUserHandler() error {
	userId, ok := h.getAuthUserId()
	if !ok {
		return h.success(map[string]interface{}{"user": nil})
	}

	user, err := models.UserFindById(h.db.GetDb(), userId)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.success(map[string]interface{}{"user": nil})
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

	if _, ok := h.getAuthUserId(); ok {
		return h.error("You are already logged in")
	}

	hashedPwd, err := utils.HashPassword(request.Password)
	if nil != err {
		return err
	}

	tr, err := h.db.GetDb().Begin()
	if nil != err {
		return errors.WithStack(err)
	}
	defer tr.Rollback()

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

	_, err = models.UserCreate(tr, user)
	if nil != err {
		return err
	}

	token := utils.CreateJWTToken(uint(user.User_id))
	if err = user.UpdateToken(tr, token); nil != err {
		return err
	}

	if err = tr.Commit(); nil != err {
		return err
	}

	return h.success(map[string]interface{}{
		"user": user,
	})
}
