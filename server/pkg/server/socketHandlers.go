package server

import (
	"encoding/json"
	"log"

	"github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/galton"
	"github.com/gorilla/websocket"
)

func sendBoardState(conn *websocket.Conn, messageType int) error {
	boardState := ServerMessage{
		MessageName: BoardState,
		MessageData: galton.CurrentBoard,
	}

	boardStateMessage, err := json.Marshal(boardState)

	if err != nil {
		log.Print("Error marshaling board state: ", err)
		return err
	}

	err = conn.WriteMessage(messageType, []byte(boardStateMessage))

	if err != nil {
		log.Println("Error writing board state message: ", err)
		return err
	}

	return nil
}

func handleNewConnection(conn *websocket.Conn, messageType int) error {
	log.Println("We have a new observer!")

	return sendBoardState(conn, messageType)
}

func handleNewBoardRequest(conn *websocket.Conn, messageType int, gridDepth int) error {
	if len(galton.CurrentBoard.Columns) == 0 {
		log.Println("Creating new board")
		galton.CurrentBoard = galton.CreateBoard(gridDepth)
	} else {
		log.Println("Board already exists.")
	}

	return sendBoardState(conn, messageType)
}
