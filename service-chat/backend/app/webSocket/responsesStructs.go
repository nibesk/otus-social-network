package webSocket

// JoinDisconnectPayload will have struct for payload of join disconnect
type JoinDisconnectPayload struct {
	Users  []UserStruct `json:"users"`
	UserID string       `json:"userID"`
}
