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

func broadcastBoardState(pool *ConnectionPool, messageType int) error {
	var err error

	for _, conn := range pool.Connections {
		currErr := sendBoardState(conn, messageType)

		if currErr != nil {
			err = currErr
		}
	}

	return err
}

func handleNewConnection(conn *websocket.Conn, pool *ConnectionPool, messageType int) error {
	log.Println("We have a new observer!")

	pool.Connections = append(pool.Connections, conn)

	return sendBoardState(conn, messageType)
}

func handleNewBoardRequest(pool *ConnectionPool, messageType int, gridDepth int) error {
	if len(galton.CurrentBoard.Columns) == 0 {
		log.Println("Creating new board")
		galton.CurrentBoard = galton.CreateBoard(gridDepth)
	} else {
		log.Println("Board already exists.")
	}

	return broadcastBoardState(pool, messageType)
}

func handleAddBallsRequest(pool *ConnectionPool, messageType int, addBallCount int) error {
	log.Printf("Adding %d ball(s) to the board", addBallCount)

	galton.CurrentBoard.AddBalls(addBallCount)

	return broadcastBoardState(pool, messageType)
}

func handleResetBoard(pool *ConnectionPool, messageType int) error {
	log.Println("Clearing the board.")

	galton.CurrentBoard = galton.GaltonBoard{}

	return broadcastBoardState(pool, messageType)
}
