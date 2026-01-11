package server

import (
	"sync"

	"github.com/Ameb8/term-sync/internal/document"
)

// Holds authoritative server-side state of document
type ServerDocument struct {
	ID      string           // Document ID
	Entries []document.Entry // Document's CRDT entries
	Clients map[*Client]bool // Map of connected clients
	Mutex   sync.Mutex       // Mutex for access serialization
}

// Creates an empty server-side document
func NewServerDocument(id string) *ServerDocument {
	return &ServerDocument{
		ID:      id,
		Entries: []document.Entry{},
		Clients: make(map[*Client]bool),
	}
}
