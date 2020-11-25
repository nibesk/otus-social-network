package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"service-users/app/handlers/requests"
	"service-users/app/models"
	"service-users/app/storage"
	"strconv"
)

func (h *Handler) ApiGetAvailableFriendsHandler() error {
	userId, _ := h.getAuthUserId()

	request := requests.AvailableFriendsRequest{}
	request.Parse(h.request)
	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	users, err := models.UserFindAllExceptUserId(storage.GetCDb(), userId, request)
	if nil != err {
		return err
	}

	userRelations, err := models.UserRelationFindByUserId(storage.GetCDb(), userId)
	if nil != err {
		return err
	}

	userRelationsMap := make(map[int]bool)
	for _, relation := range userRelations {
		if relation.User_id == userId {
			userRelationsMap[relation.Friend_user_id] = true
		} else {
			userRelationsMap[relation.User_id] = true
		}
	}

	usersPublic := make([]interface{}, len(users))
	for i, user := range users {
		publicData := user.Public()

		_, ok := userRelationsMap[user.User_id]
		publicData["is_friend"] = ok

		usersPublic[i] = publicData
	}

	return h.success(map[string]interface{}{
		"users": usersPublic,
	})
}

func (h *Handler) ApiGetFriendsHandler() error {
	userId, _ := h.getAuthUserId()
	users, err := models.UserFindFriendsForUser(storage.GetCDb(), userId)
	if nil != err {
		return err
	}

	usersPublic := make([]interface{}, len(users))
	for i, user := range users {
		usersPublic[i] = user.Public()
	}

	return h.success(map[string]interface{}{
		"users": usersPublic,
	})
}

func (h *Handler) ApiGetUserByIdHandler() error {
	userIdString, ok := mux.Vars(h.request)["userId"]
	if !ok {
		return h.error("userId is required")
	}

	userId, err := strconv.Atoi(userIdString)
	if nil != err {
		return h.error("userId must be a number")
	}

	user, err := models.UserFindById(storage.GetDb(), userId)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.success(map[string]interface{}{
			"users": nil,
		})
	case nil != err:
		return err
	}

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}

func (h *Handler) ApiAddFriendHandler() error {
	var request *requests.UserRelationRequest
	if err := h.decodeRequest(&request); nil != err {
		return errors.WithStack(err)
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	userId, _ := h.getAuthUserId()
	if userId == request.Friend_user_id {
		return h.error("You can't add yourself as a friend")
	}

	user, err := models.UserFindById(storage.GetDb(), request.Friend_user_id)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error(fmt.Sprintf("User with id %d is not found", request.Friend_user_id))
	case nil != err:
		return err
	}

	_, err = models.UserRelationCreate(storage.GetDb(), userId, request.Friend_user_id)
	if nil != err {
		return err
	}

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}

func (h *Handler) ApiDeleteFriendHandler() error {
	var request *requests.UserRelationRequest
	if err := h.decodeRequest(&request); nil != err {
		return errors.WithStack(err)
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	userId, _ := h.getAuthUserId()
	user, err := models.UserFindById(storage.GetDb(), request.Friend_user_id)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error(fmt.Sprintf("User with id %d is not found", request.Friend_user_id))
	case nil != err:
		return err
	}

	err = models.UserRelationDelete(storage.GetDb(), userId, request.Friend_user_id)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error(fmt.Sprintf("User doesn't have relation with user id %d", request.Friend_user_id))
	case nil != err:
		return err
	}

	return h.success(map[string]interface{}{
		"user": user.Public(),
	})
}
