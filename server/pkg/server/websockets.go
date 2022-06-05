package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	return origin == "http://localhost:3000"
}}

type SocketMessage struct {
	MessageText string `json:"msg"`
	TimeStamp   int    `json:"timestamp"`
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error during upgrade: ", err)
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		var decodedMessage SocketMessage

		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		log.Printf("Message: %s", message)

		if err := json.Unmarshal(message, &decodedMessage); err != nil {
			upgrader.Error(w, r, http.StatusBadRequest, err)

			log.Print("Error parsing json: ", err)
		}

		fmt.Println(decodedMessage.TimeStamp)

		err = conn.WriteMessage(messageType, []byte("Message has been received"))

		if err != nil {
			log.Println("Error writing message: ", err)
			break
		}
	}
}
