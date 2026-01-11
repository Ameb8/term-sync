package main

import "github.com/gorilla/websocket"

// Struct to represent connected client
type Client struct {
	Conn *websocket.Conn // Websocket connection
	Send chan []byte     // Message buffer
	Doc  *ServerDocument // Document being edited by client
}
