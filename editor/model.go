package editor

import (
	"github.com/Ameb8/term-sync/document"
	tea "github.com/charmbracelet/bubbletea"
)

// Bubble Tea model
type Model struct {
	// Document/Editor state
	Doc    *document.Document
	Editor *Editor

	// Cursor coordinates
	CursorX int
	CursorY int

	// Cursor size
	Width  int
	Height int

	// Server channel to listen to remote updates
	// serverCh <-chan ServerEvent
}

// Visible text in editor
type Editor struct {
	lines [][]rune
}

// Initializes editor from given Document
func InitEditor(doc *document.Document) *Editor {
	e := &Editor{}
	e.Rebuild(doc)
	return e
}

func (e *Editor) Rebuild(doc *document.Document) {
	e.lines = nil
	var line []rune

	for _, entry := range doc.Entries {
		if !entry.Visible {
			continue
		} else if entry.Value == '\n' {
			e.lines = append(e.lines, line)
			line = nil
		} else { //
			line = append(line, entry.Value)
		}
	}

	e.lines = append(e.lines, line)
}

func (m *Model) DocumentCursorIndex() int {
	index := 0

	// Count full lines before CursorY
	for y := 0; y < m.CursorY && y < len(m.Editor.lines); y++ {
		index += len(m.Editor.lines[y]) // characters
		index++                         // newline
	}

	// Add column offset on current line
	if m.CursorY < len(m.Editor.lines) {
		index += min(m.CursorX, len(m.Editor.lines[m.CursorY]))
	}

	return index
}

func (m Model) Init() tea.Cmd {
	return nil
}
