package server

type ClientMessage struct {
	MessageText string `json:"msg"`
	TimeStamp   int    `json:"timestamp"`
}

type ServerMessage struct {
	MessageName MessageName
	MessageData interface{}
}

// Messages that can be passed to and received from the client
type MessageName string

const (
	BoardState  MessageName = "board state"
	NewObserver MessageName = "new observer"
)
