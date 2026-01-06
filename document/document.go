package document

import "math/rand"

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
	Entries []Entry // Ordered slice of all characters
	Site    int     // this client’s unique ID
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

func (doc *Document) neighborsAt(cursor int) (EntryID, EntryID) {
	visible := 0
	var left EntryID = BeginID

	for _, e := range doc.Entries {
		if !e.Visible {
			continue
		}

		if visible == cursor {
			return left, e.ID
		}

		left = e.ID
		visible++
	}

	return left, EndID
}

func (doc *Document) InsertAt(cursor int, r rune) {
	leftID, rightID := doc.neighborsAt(cursor)
	newID := EntryIDBetween(leftID, rightID, doc.Site)

	entry := Entry{ID: newID, Value: r, Visible: true}
	doc.insertSorted(entry)

	//broadcastInsert(entry)
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
	doc := &Document{Site: site}

	for _, r := range string(data) {
		doc.InsertAt(len(doc.Entries), r)
	}

	return doc
}

func (doc *Document) insertSorted(entry Entry) {
	i := 0
	for i < len(doc.Entries) && CompareEntryID(doc.Entries[i].ID, entry.ID) < 0 {
		i++
	}

	doc.Entries = append(doc.Entries, Entry{})
	copy(doc.Entries[i+1:], doc.Entries[i:])
	doc.Entries[i] = entry
}

func (doc *Document) DeleteAt(cursor int) {
	if cursor <= 0 { // Validate cursor position
		return
	}

	visible := 0 // Track visible characters

	// Find
	for i := 0; i < len(doc.Entries); i++ {
		e := &doc.Entries[i]

		// Skip non-visible characters
		if !e.Visible {
			continue
		}

		// Preceeding character found
		if visible == cursor-1 {
			e.Visible = false // Set as non-visible
			// broadcastDelete(e.ID)
			return
		}

		visible++
	}
}
