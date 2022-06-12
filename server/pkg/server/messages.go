package server

type ClientMessage struct {
	MessageText string `json:"msg"`
	Quantity    int    `json:"quantity"`
}

type ServerMessage struct {
	MessageName MessageName
	MessageData interface{}
}

// Messages that can be passed to and received from the client
type MessageName string

const (
	NewObserver MessageName = "new observer"
	CreateBoard MessageName = "create board"
	BoardState  MessageName = "board state"
	AddBalls    MessageName = "add balls"
	ResetBoard  MessageName = "reset board"
)
