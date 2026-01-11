package server

import (
	"encoding/json"
	"log"
)

// Continously reads and dispatches messages from WS connection
func (c *Client) readPump(server *Server) {
	defer c.Conn.Close() // Ensure connection closes on exit

	for {
		_, data, err := c.Conn.ReadMessage() // Read next message

		if err != nil { // Error reading message
			log.Println("read:", err)
			return
		}

		var msg Message // Instantiate message

		// Read JSON message
		if err := json.Unmarshal(data, &msg); err != nil {
			// Error reading message
			log.Println("bad message:", err)
			continue
		}

		// Dispatch message based on type
		switch msg.Type {
		case "join":
			c.handleJoin(server, msg)
		}
	}
}

// Register client with the requested document
func (c *Client) handleJoin(server *Server, msg Message) {
	doc := server.GetOrCreateDoc(msg.Doc) // Get document

	// Lock document to modify client list
	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	// Register client
	doc.Clients[c] = true
	c.Doc = doc

	// Send full CRDT state
	resp := Message{
		Type: "full",
		Doc:  doc.ID,
		Full: doc.Entries,
	}

	data, _ := json.Marshal(resp) // Convert message to JSON
	c.Send <- data                // Send to client
}
