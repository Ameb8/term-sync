package localdocument

import (
	"github.com/Ameb8/term-sync/document"
)

type LocalDocument struct {
	document   document.Document
	projection projection
	site       int
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
	entries entryStore // Ordered slice of all characters
}

var (
	BeginID = EntryID{Elements: nil}
	EndID   = EntryID{Elements: []PathElem{{Digit: 1 << 30, Site: 0}}}
)

func CompareEntryID(a, b EntryID) int {
	n := len(a.Elements)
	if len(b.Elements) < n {
		n = len(b.Elements)
	}

	for i := 0; i < n; i++ {
		if a.Elements[i].Digit != b.Elements[i].Digit {
			return a.Elements[i].Digit - b.Elements[i].Digit
		}
		if a.Elements[i].Site != b.Elements[i].Site {
			return a.Elements[i].Site - b.Elements[i].Site
		}
	}

	return len(a.Elements) - len(b.Elements)
}

func DocumentFromBytes(data []byte, site int) *Document {
	doc := &Document{
		Site:       site,
		entries:    newSliceStore(),
		projection: newLineProjection(),
	}

	for _, r := range string(data) {
		doc.InsertAt(doc.entries.len(), r)
	}

	return doc
}

func (localDoc *LocalDocument) rebuildProjection() {
	localDoc.projection.reset()

	cursor := 0
	doc.entries.iterVisible(func(e Entry) {
		doc.projection.insert(cursor, e.Value)
		cursor++
	})
}

// Insert char at given cursor location
func (localDoc *LocalDocument) InsertAt(cursor int, r rune) {
	// Determine id for new entry
	localDoc.document.InsertAt(cursor, r)

	// Update projection
	localDoc.projection.insert(cursor, r)

	//broadcastInsert(entry)
}

func (localDoc *LocalDocument) DeleteAt(cursor int) {
	localDoc.document.DeleteAt(cursor)
	localDoc.projection.delete(cursor)
}

func (doc *LocalDocument) String() string {
	return localDoc.projection.string()
}

func (localDoc *LocalDocument) Project() [][]rune {
	return localDoc.projection.getLines()
}
