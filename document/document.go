package document

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

func EntryIDBetween(left, right EntryID, site int) EntryID {
	depth := 0

	for {
		var lDigit, rDigit int

		if depth < len(left.Elements) {
			lDigit = left.Elements[depth].Digit
		} else {
			lDigit = 0
		}

		if depth < len(right.Elements) {
			rDigit = right.Elements[depth].Digit
		} else {
			rDigit = 1 << 30
		}

		if rDigit-lDigit > 1 {
			d := lDigit + 1 + rand.Intn(rDigit-lDigit-1)
			newElems := append([]PathElem{}, left.Elements[:depth]...)
			newElems = append(newElems, PathElem{Digit: d, Site: site})
			return EntryID{Elements: newElems}
		}

		depth++
	}
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

func (doc *Document) rebuildProjection() {
	doc.projection.reset()

	cursor := 0
	doc.entries.iterVisible(func(e Entry) {
		doc.projection.insert(cursor, e.Value)
		cursor++
	})
}

// Insert char at given cursor location
func (doc *Document) InsertAt(cursor int, r rune) {
	// Determine id for new entry
	leftID, rightID := doc.entries.getNeighbors(cursor)
	newID := EntryIDBetween(leftID, rightID, doc.Site)

	// Create and insert entry
	entry := Entry{ID: newID, Value: r, Visible: true}
	doc.entries.insert(entry)

	// Update projection
	doc.projection.insert(cursor, r)

	//broadcastInsert(entry)
}

func (doc *Document) DeleteAt(cursor int) {
	doc.entries.deleteByCursor(cursor)
	doc.projection.delete(cursor)
}

func (doc *Document) String() string {
	return doc.projection.string()
}

func (doc *Document) Project() [][]rune {
	return doc.projection.getLines()
}

func (e Entry) getRune() {
	return e.Value
}

func (doc *Document) IterVisible(f func(r rune)) {
	doc.entries.iterVisible(func(e Entry) {
		f(e.Value)
	})
}
