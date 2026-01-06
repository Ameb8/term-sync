package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Ameb8/term-sync/editor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: term-sync <file>")
		os.Exit(1)
	}

	path := os.Args[1] // Get filepath

	doc := DocumentFromBytes

	ed := editor.InitEditor(doc)

	model := editor.Model{
		Document: doc,
		Editor:   ed,
		CursorX:  0,
		CursorY:  0,
	}

	p := tea.NewProgram(model, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
