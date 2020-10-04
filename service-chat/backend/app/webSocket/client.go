package webSocket

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// Event struct of socket events
type Event struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub      *Hub
	wsConn   *websocket.Conn
	send     chan Event
	username string
	userID   string
}

func (client *Client) handleIncomingEvents(incomingEvent Event) {
	switch incomingEvent.Event {
	case eventJoin:
		log.Printf("client [%s]. Join Event triggered", client.username)

		responseEvent := Event{
			Event: incomingEvent.Event,
			Payload: JoinDisconnectPayload{
				UserID: client.userID,
				Users:  client.hub.getAllConnectedUsers(),
			},
		}

		// do not send messages to current user on join incomingEvent. Raise condition in channel "client.send". Use mutex on client if required
		client.hub.BroadcastToAllClientsExceptOne(responseEvent, *client)

	case eventDisconnect:
		log.Printf("client [%s]. Disconnect Event triggered", client.username)

		responseEvent := Event{
			Event: incomingEvent.Event,
			Payload: JoinDisconnectPayload{
				UserID: client.userID,
				Users:  client.hub.getAllConnectedUsers(),
			},
		}

		client.hub.BroadcastToAllClientsExceptOne(responseEvent, *client)

	case eventMessage:
		log.Printf("client [%s]. Message Event triggered", client.username)

		selectedUserID := incomingEvent.Payload.(map[string]interface{})["userId"].(string)
		responseEvent := Event{
			Event: "message",
			Payload: map[string]interface{}{
				"username": client.hub.getUsernameByUserID(selectedUserID),
				"message":  incomingEvent.Payload.(map[string]interface{})["message"],
				"userId":   selectedUserID,
			},
		}
		client.hub.EmitToClient(responseEvent, selectedUserID)
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.wsConn.Close()
		log.Printf("[CLOSE CONNECTION] client [%c]", c.username)
	}()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error { c.wsConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	var event Event

	for {
		_, payload, err := c.wsConn.ReadMessage()

		if err != nil {
			var unexpectedCloseError string
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				unexpectedCloseError = "IsUnexpectedCloseError"
			}
			log.Printf("[ERROR] readPump cycle break %s: %v", unexpectedCloseError, err)
			break
		}

		decoder := json.NewDecoder(bytes.NewReader(payload))
		decoderErr := decoder.Decode(&event)

		if decoderErr != nil {
			log.Printf("[ERROR] readPump decoder: %+v", decoderErr)
			break
		}

		c.handleIncomingEvents(event)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.wsConn.Close()
	}()

	for {
		select {
		case payload, ok := <-c.send:
			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(payload)
			finalPayload := reqBodyBytes.Bytes()

			c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.wsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.wsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(finalPayload)

			n := len(c.send)
			for i := 0; i < n; i++ {
				json.NewEncoder(reqBodyBytes).Encode(<-c.send)
				w.Write(reqBodyBytes.Bytes())
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.wsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
