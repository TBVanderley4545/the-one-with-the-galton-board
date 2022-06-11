package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	return origin == "http://localhost:3000"
}}

type ClientMessage struct {
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
		var decodedMessage ClientMessage

		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		if err := json.Unmarshal(message, &decodedMessage); err != nil {
			upgrader.Error(w, r, http.StatusBadRequest, err)

			log.Print("Error parsing json: ", err)
		}

		if decodedMessage.MessageText == "new observer" {
			log.Println("We have a new observer!!")

			if len(galton.CurrentBoard.Columns) == 0 {
				log.Println("Creating new board")
				galton.CurrentBoard = galton.CreateBoard(10)
			} else {
				log.Println("Board already exists.")
				log.Println(galton.CurrentBoard)
			}

			if boardState, err := json.Marshal(galton.CurrentBoard); err != nil {
				log.Print("Error parsing json: ", err)

				err = conn.WriteMessage(messageType, []byte(boardState))

				if err != nil {
					log.Println("Error writing message: ", err)
					break
				}
			}
		}

		log.Println(decodedMessage.TimeStamp)

		err = conn.WriteMessage(messageType, []byte("Message has been received"))

		if err != nil {
			log.Println("Error writing message: ", err)
			break
		}
	}
}
