package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/server"
)

func main() {
	port := 8080
	formattedPort := fmt.Sprintf(":%d", port)

	log.Printf("Server has started on port: %d\n", port)

	http.HandleFunc("/", server.RootRoute)
	http.HandleFunc("/galton-board/", server.SocketHandler)

	log.Fatal(http.ListenAndServe(formattedPort, nil))
}
