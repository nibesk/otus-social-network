package webSocket

// Event struct of socket events
type Event struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

const (
	eventMessage          = "message"
	eventForwardedMessage = "messageForwarded"
	eventError            = "error"
	eventStartUp          = "startUp"
)

func createErrorEvent(msg string) Event {
	return Event{Event: eventError, Payload: []string{msg}}
}

func createErrorsEvent(msgs []string) Event {
	return Event{Event: eventError, Payload: msgs}
}
