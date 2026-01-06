package document

type SliceStore struct {
	entries []Entry
}

func NewSliceStore() *SliceStore {
	return &SliceStore{entries: []Entry{}}
}

func (s *SliceStore) Insert(entry Entry) {
	i := 0
	for i < len(s.entries) && CompareEntryID(s.entries[i].ID, entry.ID) < 0 {
		i++
	}
	s.entries = append(s.entries, Entry{})
	copy(s.entries[i+1:], s.entries[i:])
	s.entries[i] = entry
}

func (s *SliceStore) DeleteByCursor(cursor int) {
	visible := 0
	for i := 0; i < len(s.entries); i++ {
		e := &s.entries[i]
		if !e.Visible {
			continue
		}
		if visible == cursor {
			e.Visible = false
			return
		}
		visible++
	}
}

func (s *SliceStore) GetNeighbors(cursor int) (EntryID, EntryID) {
	visible := 0
	var left EntryID = BeginID
	for _, e := range s.entries {
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

func (s *SliceStore) IterVisible(f func(e Entry)) {
	for _, e := range s.entries {
		if e.Visible {
			f(e)
		}
	}
}

func (s *SliceStore) Len() int {
	count := 0
	for _, e := range s.entries {
		if e.Visible {
			count++
		}
	}
	return count
}
