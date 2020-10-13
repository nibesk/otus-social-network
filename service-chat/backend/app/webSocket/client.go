package webSocket

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
	"service-chat/app/customErrors"
	"service-chat/app/services"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub    *Hub
	wsConn *websocket.Conn
	send   chan Event
	id     string
	token  string
	user   *services.UserModel
}

// Incoming events handler
func (client *Client) handleIncomingEvents(e Event) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC] %w; Stack trace %s", r, string(debug.Stack()))
		}
	}()

	log.Printf(">>%s from client [%s]", e.Event, client.id)

	h := Handler{client: client}

	// check if users authorized in service
	if "" == client.token && eventStartUp != e.Event {
		h.sendErrorsToSelf("client has empty token. send startup event")

		return
	}

	if eventStartUp == e.Event && nil != client.user {
		h.sendErrorsToSelf("client already authorized")

		return
	}

	var err error
	switch e.Event {
	case eventMessage:
		err = h.MessageHandler(e)
	case eventStartUp:
		err = h.StartUpHandler(e)
	}

	if nil != err {
		switch causedErr := errors.Cause(err).(type) {
		case *customErrors.TypedError:
			h.sendErrorsToSelf(causedErr.Msg)
		default:
			h.sendErrorsToSelf("Something went wrong...")
		}

		log.Printf("[ERROR] %+v\n", err)
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
		log.Printf("[CLOSE CONNECTION] client [%s]", c.id)
	}()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error { c.wsConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	var event Event

	for {
		_, payload, err := c.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				log.Printf("[UnexpectedCloseError] readPump cycle break: %v", err)
			}

			break
		}

		decoder := json.NewDecoder(bytes.NewReader(payload))
		if err := decoder.Decode(&event); err != nil {
			log.Printf("[ERROR] readPump decoder: %+v", err)

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

			log.Printf("<<%s to client [%s]", payload.Event, c.id)

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
