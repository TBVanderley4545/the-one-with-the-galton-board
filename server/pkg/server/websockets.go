package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	return origin == "http://localhost:3000"
}}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error during upgrade: ", err)
	}

	defer conn.Close()

socketLoop:
	for {
		messageType, message, err := conn.ReadMessage()
		var decodedMessage ClientMessage

		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		if err := json.Unmarshal(message, &decodedMessage); err != nil {
			upgrader.Error(w, r, http.StatusBadRequest, err)

			log.Print("Error parsing json: ", err)
		}

		switch decodedMessage.MessageText {
		case string(NewObserver):
			if err := handleNewConnection(conn, messageType); err != nil {
				break socketLoop
			}
		case string(CreateBoard):
			if err := handleNewBoardRequest(conn, messageType, decodedMessage.Quantity); err != nil {
				break socketLoop
			}
		case string(AddBalls):
			if err := handleAddBallsRequest(conn, messageType, decodedMessage.Quantity); err != nil {
				break socketLoop
			}
		case string(ResetBoard):
			if err := handleResetBoard(conn, messageType); err != nil {
				break socketLoop
			}
		}
	}
}
