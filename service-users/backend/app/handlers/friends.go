package handlers

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"service-users/app/handlers/requests"
	"service-users/app/models"
)

func (h *Handler) ApiGetAvailableFriendsHandler() error {
	userId, err := h.getSessionUserId()
	if nil != err {
		return errors.WithStack(err)
	}

	request := requests.AvailableFriendsRequest{}
	request.Parse(h.request)
	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	users, err := models.UserFindAllExceptUserId(h.db, userId, request)
	if nil != err {
		return err
	}

	userRelations, err := models.UserRelationFindByUserId(h.db, userId)
	if nil != err {
		return err
	}

	userRelationsMap := make(map[int]bool)
	for _, relation := range userRelations {
		userRelationsMap[relation.Friend_user_id] = true
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
	userId, err := h.getSessionUserId()
	if nil != err {
		return err
	}

	users, err := models.UserFindFriendsForUser(h.db, userId)
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

func (h *Handler) ApiAddFriendHandler() error {
	var request *requests.UserRelationRequest
	if err := h.decodeRequest(&request); nil != err {
		return errors.WithStack(err)
	}

	if violations := h.checkValidations(request); nil != violations {
		return h.error(violations)
	}

	userId, err := h.getSessionUserId()
	if nil != err {
		return err
	}

	if userId == request.Friend_user_id {
		return h.error("You can't add yourself as a friend")
	}

	user, err := models.UserFindById(h.db, request.Friend_user_id)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error(fmt.Sprintf("User with id %d is not found", request.Friend_user_id))
	case nil != err:
		return err
	}

	_, err = models.UserRelationCreate(h.db, userId, request.Friend_user_id)
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

	userId, err := h.getSessionUserId()
	if nil != err {
		return errors.WithStack(err)
	}

	user, err := models.UserFindById(h.db, request.Friend_user_id)
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return h.error(fmt.Sprintf("User with id %d is not found", request.Friend_user_id))
	case nil != err:
		return err
	}

	err = models.UserRelationDelete(h.db, userId, request.Friend_user_id)
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
