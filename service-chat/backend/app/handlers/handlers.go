package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net/http"
	"service-chat/app/config"
	"service-chat/app/customErrors"
	"service-chat/app/models"
	"service-chat/app/services"
	"service-chat/app/webSocket"
	"strconv"
)

func (h *Handler) ViewIndexHandler() error {
	h.writer.Write([]byte("<h1>Hello from service chat!</h1>"))

	return nil
}

func (h *Handler) ConnectToSocket() error {
	h.disableResponse()

	upgrader := websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: true,
	}
	if config.Env.Server.WsCheckOrigin {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	username := "User:" + string(([]rune(uuid.New().String()))[0:8])

	// Upgrading the HTTP connection socket connection
	conn, err := upgrader.Upgrade(h.writer, h.request, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	webSocket.GetHub().CreateNewSocketUser(conn, username)

	return nil
}

func (h *Handler) GetMessages() error {
	token, err := h.getAuthorizationHeader()
	if nil != err {
		return err
	}

	friendUserIdInt, err := strconv.Atoi(mux.Vars(h.request)["userId"])
	if nil != err {
		return customErrors.TypedError{"friend_user_id must be numeric"}
	}

	serviceUsers := services.ServiceUsers{Token: token}
	user, err := serviceUsers.GetUser()
	if nil != err {
		return err
	}

	messages, err := models.MessageFindConversation(user.User_id, friendUserIdInt)
	if nil != err {
		return err
	}

	messagesWithUser := make([]*webSocket.MessageToClientPayload, len(messages))
	for i, m := range messages {
		messagesWithUser[i] = &webSocket.MessageToClientPayload{
			Message:   m.Text,
			UserId:    m.From_user_id,
			Timestamp: m.CreatedAt.Unix(),
		}
	}

	return h.success(map[string]interface{}{
		"messages": messagesWithUser,
	})
}
