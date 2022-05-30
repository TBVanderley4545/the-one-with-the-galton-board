package server

import (
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

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		log.Printf("Message: %s", message)

		err = conn.WriteMessage(messageType, message)

		if err != nil {
			log.Println("Error writing message: ", err)
			break
		}
	}
}
