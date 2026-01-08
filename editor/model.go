package editor

import (
	"fmt"
	"os"

	"github.com/Ameb8/term-sync/document"
	tea "github.com/charmbracelet/bubbletea"
)

// Bubble Tea model
type Model struct {
	// Document/Editor state
	Doc *document.Document

	// Cursor coordinates
	CursorX int
	CursorY int

	// Vertical viewport offset
	ViewportY int

	// Cursor size
	Width  int
	Height int

	// Path to file on system
	Path string
}

func (m *Model) DocumentCursorIndex() int {
	index := 0
	lines := m.Doc.Project()

	// Count full lines before CursorY
	for y := 0; y < m.CursorY && y < len(lines); y++ {
		index += len(lines[y]) // characters
		index++                // newline
	}

	// Add column offset on current line
	if m.CursorY < len(lines) {
		index += min(m.CursorX, len(lines[m.CursorY]))
	}

	return index
}

func (m *Model) Init() tea.Cmd {
	return nil
}

// Save file to host
func (m *Model) Save() error {
	if m.Path == "" {
		return fmt.Errorf("no path set")
	}

	data := []byte(m.Doc.String())
	return os.WriteFile(m.Path, data, 0o644)
}
