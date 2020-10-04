package webSocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// UserStruct is used for sending users with socket id
type UserStruct struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
}

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
}

var hubStore *Hub

func InitHub() *Hub {
	hubStore = &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}

	return hubStore
}

func GetHub() *Hub {
	return hubStore
}

// Run will execute Go Routines to check incoming Socket events
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.handleUserRegisterEvent(client)

		case client := <-hub.unregister:
			hub.handleUserDisconnectEvent(client)
		}
	}
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

const (
	eventJoin       = "join"
	eventDisconnect = "disconnect"
	eventMessage    = "message"
)

// CreateNewSocketUser creates a new socket user
func (hub *Hub) CreateNewSocketUser(connection *websocket.Conn, username string) {
	uniqueID := uuid.New()

	log.Printf("client with userID [%s] has been created", uniqueID)

	client := &Client{
		hub:      hub,
		wsConn:   connection,
		send:     make(chan Event),
		username: username,
		userID:   uniqueID.String(),
	}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	client.hub.register <- client
}

// handleUserRegisterEvent will handle the Join event for New socket users
func (hub *Hub) handleUserRegisterEvent(client *Client) {
	hub.clients[client.userID] = client
	client.handleIncomingEvents(Event{
		Event:   eventJoin,
		Payload: client.userID,
	})
}

// handleUserDisconnectEvent will handle the Disconnect event for socket users
func (hub *Hub) handleUserDisconnectEvent(client *Client) {
	_, ok := hub.clients[client.userID]
	if !ok {
		return
	}

	delete(hub.clients, client.userID)
	close(client.send)

	client.handleIncomingEvents(Event{
		Event:   eventDisconnect,
		Payload: client.userID,
	})
}

// EmitToClient will emit the socket event to specific socket user
func (hub *Hub) EmitToClient(payload Event, userID string) {
	client, ok := hub.clients[userID]
	if !ok {
		return
	}

	select {
	case client.send <- payload:
	default:
		close(client.send)
		delete(hub.clients, userID)
	}
}

// BroadcastToAllClients will emit the socket events to all socket users
func (hub *Hub) BroadcastToAllClients(payload Event) {
	for id, client := range hub.clients {
		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, id)
		}
	}
}

func (hub *Hub) BroadcastToAllClientsExceptOne(payload Event, exceptClient Client) {
	for id, client := range hub.clients {
		if exceptClient.userID == client.userID {
			continue
		}

		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, id)
		}
	}
}

func (hub *Hub) getUsernameByUserID(userID string) string {
	client, ok := hub.clients[userID]
	if !ok {
		return ""
	}

	return client.username
}

func (hub *Hub) getAllConnectedUsers() []UserStruct {
	var users = make([]UserStruct, 0, len(hub.clients))
	for _, client := range hub.clients {
		users = append(users, UserStruct{
			Username: client.username,
			UserID:   client.userID,
		})
	}
	return users
}

func (hub *Hub) getAllConnectedUsersExceptOne(exceptClient Client) []UserStruct {
	var users []UserStruct
	for _, client := range hub.clients {
		if exceptClient.userID == client.userID {
			continue
		}

		users = append(users, UserStruct{
			Username: client.username,
			UserID:   client.userID,
		})
	}
	return users
}
