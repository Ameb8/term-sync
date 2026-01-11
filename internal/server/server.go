package server

import "sync"

// Provides concurrency-safe access to all documents
type Server struct {
	Docs  map[string]*ServerDocument // Maps IDs to documents
	Mutex sync.Mutex                 // Document map mutex
}

// Creates a new empty server
func NewServer() *Server {
	return &Server{
		Docs: make(map[string]*ServerDocument),
	}
}

// Fetches a document by ID or creates it
func (s *Server) GetOrCreateDoc(id string) *ServerDocument {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	doc, ok := s.Docs[id]
	if !ok {
		doc = NewServerDocument(id)
		s.Docs[id] = doc
	}
	return doc
}
