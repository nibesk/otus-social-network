package webSocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"service-chat/app/utils"
	"sync"
	"time"
)

// UserStruct is used for sending users with socket id
type UserStruct struct {
	Username string `json:"username"`
	UserID   string `json:"id"`
}

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	sync.Mutex
	clients         map[string]*Client
	userToClientMap map[int][]string

	register   chan *Client
	unregister chan *Client
}

var hubStore *Hub

func InitHub() *Hub {
	hubStore = &Hub{
		clients:         make(map[string]*Client),
		userToClientMap: make(map[int][]string),

		register:   make(chan *Client),
		unregister: make(chan *Client),
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
			hub.handleClientRegisterEvent(client)

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

// CreateNewSocketUser creates a new socket user
func (hub *Hub) CreateNewSocketUser(conn *websocket.Conn, username string) {
	uniqueID := uuid.New()

	log.Printf("client with id [%s] has been created", uniqueID)

	client := &Client{
		hub:    hub,
		wsConn: conn,
		send:   make(chan Event),
		id:     uniqueID.String(),
	}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

	client.hub.register <- client
}

// handleClientRegisterEvent will handle the Join event for New socket users
func (hub *Hub) handleClientRegisterEvent(client *Client) {
	hub.clients[client.id] = client
}

// handleUserDisconnectEvent will handle the Disconnect event for socket users
func (hub *Hub) handleUserDisconnectEvent(client *Client) {
	_, ok := hub.clients[client.id]
	if !ok {
		return
	}

	if nil != client.user {
		hub.DeleteUserFromClientsMap(client.user.User_id, client.id)
	}

	delete(hub.clients, client.id)
	close(client.send)
}

func (hub *Hub) GetTargetClientByClientId(clientId string) *Client {
	client, ok := hub.clients[clientId]
	if !ok {
		return nil
	}

	return client
}

func (h *Hub) AddUserToClientsMap(userId int, clientId string) {
	h.Lock()
	defer h.Unlock()

	_, ok := h.userToClientMap[userId]
	if ok {
		h.userToClientMap[userId] = append(h.userToClientMap[userId], clientId)
	} else {
		h.userToClientMap[userId] = []string{clientId}
	}
}

func (h *Hub) DeleteUserFromClientsMap(userId int, clientId string) {
	h.Lock()
	defer h.Unlock()

	if len(h.userToClientMap[userId]) > 1 {
		h.userToClientMap[userId] = utils.StringRemoveItemFromArray(h.userToClientMap[userId], clientId)

		return
	}

	delete(h.userToClientMap, userId)
}

func (h *Hub) EmitToClientByUserid(event Event, userId int) {
	clientIds, ok := h.userToClientMap[userId]
	if !ok {
		return
	}

	for _, clientId := range clientIds {
		h.EmitToClient(event, clientId)
	}

	return
}

// EmitToClient will emit the socket event to specific socket user
func (hub *Hub) EmitToClient(event Event, clientId string) {
	client, ok := hub.clients[clientId]
	if !ok {
		return
	}

	select {
	case client.send <- event:
	default:
		close(client.send)
		delete(hub.clients, client.id)
	}
}

// BroadcastToAllClients will emit the socket events to all socket users
func (hub *Hub) BroadcastToAllClients(event Event) {
	for id, client := range hub.clients {
		select {
		case client.send <- event:
		default:
			close(client.send)
			delete(hub.clients, id)
		}
	}
}

func (hub *Hub) BroadcastToAllClientsExceptOne(event Event, exceptClient Client) {
	for id, client := range hub.clients {
		if exceptClient.id == client.id {
			continue
		}

		select {
		case client.send <- event:
		default:
			close(client.send)
			delete(hub.clients, id)
		}
	}
}

func (hub *Hub) getAllConnectedUsers() []UserStruct {
	var users = make([]UserStruct, 0, len(hub.clients))
	for _, client := range hub.clients {
		users = append(users, UserStruct{
			UserID: client.id,
		})
	}
	return users
}

func (hub *Hub) getAllConnectedUsersExceptOne(exceptClient Client) []UserStruct {
	var users []UserStruct
	for _, client := range hub.clients {
		if exceptClient.id == client.id {
			continue
		}

		users = append(users, UserStruct{
			UserID: client.id,
		})
	}
	return users
}
