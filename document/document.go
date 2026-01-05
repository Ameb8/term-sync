package document

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

func (doc *Document) InsertAt(cursor int, r rune) {
	leftID, rightID := doc.neighborsAt(cursor)
	newID := EntryIDBetween(leftID, rightID, doc.Site)

	entry := Entry{ID: newID, Value: r, Visible: true}
	doc.insertSorted(entry)

	broadcastInsert(entry)
}
