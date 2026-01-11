package localdocument

import (
	"github.com/Ameb8/term-sync/internal/document"
)

type projection interface {
	// Update projection state
	insert(cursor int, r rune)
	delete(cursor int)

	reset()

	string() string // View as string
	getLines() [][]rune

	// Utility methods
	len() int
	lineCount() int
}

type LocalDocument struct {
	document   *document.Document // Underlying CRDT document
	projection projection         // Document projection
}

func LocalDocumentFromBytes(data []byte, site int) *LocalDocument {
	return &LocalDocument{
		document:   document.DocumentFromBytes(data, site),
		projection: newLineProjection(),
	}
}

func (doc *LocalDocument) rebuildProjection() {
	doc.projection.reset() // Reset projection state
	cursor := 0            // Reset cursor

	// Add visible characters
	doc.document.IterVisible(func(r rune) {
		doc.projection.insert(cursor, r)
		cursor++
	})
}

// Insert char at given cursor location
func (doc *LocalDocument) InsertAt(cursor int, r rune) {
	doc.document.InsertAt(cursor, r) // Update CRDT state
	doc.projection.insert(cursor, r) // Update projection

	//broadcastInsert(entry)
}

func (doc *LocalDocument) DeleteAt(cursor int) {
	doc.document.DeleteAt(cursor)
	doc.projection.delete(cursor)
}

func (doc *LocalDocument) String() string {
	return doc.projection.string()
}

func (doc *LocalDocument) Project() [][]rune {
	return doc.projection.getLines()
}
