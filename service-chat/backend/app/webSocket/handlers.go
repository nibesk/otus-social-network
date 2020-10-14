package webSocket

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"log"
	"service-chat/app/models"
	"service-chat/app/services"
	"service-chat/app/utils"
	"time"
)

type Handler struct {
	client *Client
}

func (h *Handler) sendErrorsToSelf(data interface{}) error {
	switch v := data.(type) {
	case string:
		h.client.send <- createErrorEvent(v)
	case []string:
		h.client.send <- createErrorsEvent(v)
	default:
		log.Println("undefined error payload")
	}

	return nil
}

func (h *Handler) MessageHandler(event Event) error {
	var payload MessageFromClientPayload
	decodePayload(event, &payload)
	if violations := checkValidations(payload); nil != violations {
		return h.sendErrorsToSelf(violations)
	}

	responseEvent := Event{
		Event: eventMessage,
		Payload: MessageToClientPayload{
			Message:   payload.Message,
			UserId:    h.client.user.User_id,
			User:      h.client.user,
			Timestamp: time.Now().Unix(),
		},
	}

	t, err := models.ThreadInsure(h.client.user.User_id, payload.UserId)
	if nil != err {
		return err
	}

	m := &models.Message{
		Text:      payload.Message,
		User_id:   h.client.user.User_id,
		Thread_id: t.ID,
	}
	models.MessageCreate(m)

	h.client.hub.EmitToClientByUserid(responseEvent, payload.UserId)

	return nil
}

func (h *Handler) StartUpHandler(event Event) error {
	var payload StartUpPayload
	decodePayload(event, &payload)
	if violations := checkValidations(payload); nil != violations {
		return h.sendErrorsToSelf(violations)
	}

	serviceUsers := services.ServiceUsers{Token: payload.Token}
	user, err := serviceUsers.GetUser()
	if nil != err {
		return err
	}

	h.client.hub.AddUserToClientsMap(user.User_id, h.client.id)

	h.client.user = user
	h.client.token = payload.Token

	log.Printf("Auth userId [%d] to client [%s]", user.User_id, h.client.id)

	return nil
}

func checkValidations(s interface{}) []string {
	err := utils.GetValidator().Struct(s)

	if err != nil {
		violations := make([]string, 0, 5)
		for _, err := range err.(validator.ValidationErrors) {
			violations = append(violations, fmt.Sprintf("Field validation for '%s' failed. %s", err.Field(), err.Tag()))
		}

		return violations
	}

	return nil
}

func decodePayload(event Event, obj interface{}) error {
	err := mapstructure.Decode(event.Payload, &obj)
	if nil != err {
		return errors.WithStack(err)
	}

	return nil
}
