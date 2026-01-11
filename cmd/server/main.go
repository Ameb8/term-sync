package main

import (
	"log"
	"net/http"

	"github.com/Ameb8/term-sync/internal/server"
)

func main() {
	// Create server instance
	server := server.NewServer()

	// Handle websocket connections at '/ws'
	http.HandleFunc("/ws", server.ServeWS)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start HTTP server
}
