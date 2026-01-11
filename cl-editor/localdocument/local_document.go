package localdocument


import "math/rand"

type entryStore interface {
	insert(entry Entry)
	deleteByCursor(cursor int)
	getNeighbors(cursor int) (EntryID, EntryID)
	iterVisible(func(e Entry))
	len() int
}

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

// Component of character's identifier
type PathElem struct {
	Digit int // Numerical identifier
	Site  int // Client's ID
}

// The unique identifier of a character
type EntryID struct {
	Elements []PathElem // Ordered sequence of PathElems as unique identifier
}

// Single character in document
type Entry struct {
	ID      EntryID // Character's immutable unique identifier
	Value   rune    // The actual character to display
	Visible bool    // Sets whether character is visible in document
}

// Document representation
type Document struct {
	entries    entryStore // Ordered slice of all characters
	Site       int        // this client’s unique ID
	projection projection
}



type LocalDocument struct {
	document Document // Underlying CRDT document
	projection projection // Document projection
}



func LocalDocumentFromBytes(data []byte, site int) *LocalDocument {
	return &LocalDocument{
		document: DocumentFromBytes(byte, site),
		projection: newLineProjection(),
	}
}


func DocumentFromBytes(data []byte, site int) *Document {
	doc := &LocalDocument{
		Site:       site,
		entries:    newSliceStore(),
		projection: newLineProjection(),
	}

	for _, r := range string(data) {
		doc.InsertAt(doc.entries.len(), r)
	}

	return doc
}

func (doc *LocalDocument) rebuildProjection() {
	doc.projection.reset() // Reset projection state
	cursor := 0 // Reset cursor

	// Add visible characters
	doc.document.iterVisible(func(r rune {
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
	doc.doc.DeleteAt(cursor)
	doc.projection.delete(cursor)
}

func (doc *LocalDocument) String() string {
	return doc.projection.string()
}

func (doc *LocalDocument) Project() [][]rune {
	return doc.projection.getLines()
}
