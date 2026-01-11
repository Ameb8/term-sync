package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Ameb8/term-sync/internal/cli/editor"
	"github.com/Ameb8/term-sync/internal/cli/localdocument"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: term-sync <file>")
		os.Exit(1)
	}

	path := os.Args[1]             // Get filepath
	data, err := os.ReadFile(path) // Read file into memory

	if err != nil { // Error reading file
		log.Fatal(err)
	}

	// Construct Document and Editor objects
	doc := localdocument.LocalDocumentFromBytes(data, 0)

	model := &editor.Model{
		Doc:     doc,
		CursorX: 0,
		CursorY: 0,
		Path:    path,
	}

	p := tea.NewProgram(model, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
