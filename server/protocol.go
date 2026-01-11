package main

import "alexbuckley/term-sync/document"

// Message definitions for server/client communication
type Message struct {
	Type string `json:"type"`          // Message type
	Doc  string `json:"doc,omitempty"` // ID of referenced document

	// Full CRDT state on client join
	Full []document.Entry `json:"full,omitempty"`
}
