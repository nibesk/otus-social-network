package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net/http"
	"service-chat/app/config"
	"service-chat/app/webSocket"
)

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
