package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrade HTTP connections wo WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accept all connections for now
	},
}

// Handle new connectgions
func (s *Server) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Attempt to upgrade to websocket
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil { // Error upgrading connection
		log.Println("upgrade:", err)
		return
	}

	// Initialize client
	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 16),
	}

	// Go routines to read/write messages
	go client.writePump()
	client.readPump(s)
}
