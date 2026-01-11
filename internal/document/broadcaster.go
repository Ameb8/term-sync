package document

type Broadcaster interface {
	BroadcastChanges(changes []Change)
}

type ChangeType string

const (
	Insert ChangeType = "insert"
	Delete ChangeType = "delete"
)

type Change struct {
	Type   ChangeType `json:"type"`
	Entry  *Entry     `json:"entry,omitempty"`  // for insert
	Cursor int        `json:"cursor,omitempty"` // for delete
}
