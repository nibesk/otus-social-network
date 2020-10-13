package webSocket

import "service-chat/app/services"

type StartUpPayload struct {
	Token string `mapstructure:"token" validate:"required"`
}

type MessageFromClientPayload struct {
	Message string `mapstructure:"message" validate:"required"`
	UserId  int    `mapstructure:"user_id" validate:"required"`
}

type MessageToClientPayload struct {
	Message   string              `json:"message"`
	UserId    int                 `json:"user_id"`
	User      *services.UserModel `json:"user"`
	Timestamp int64               `json:"timestamp"`
}
