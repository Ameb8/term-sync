package main

import (
	"log"
	"net/http"
)

func main() {
	// Create server instance
	server := NewServer()

	// Handle websocket connections at '/ws'
	http.HandleFunc("/ws", server.ServeWS)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start HTTP server
}
