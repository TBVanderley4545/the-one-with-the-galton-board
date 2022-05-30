package main

import (
	"github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.RootRoute)
	http.HandleFunc("/galton-board/", server.SocketHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
