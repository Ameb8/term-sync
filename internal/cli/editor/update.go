package editor

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) clampViewport() {
	usableHeight := m.Height - 1
	if usableHeight <= 0 {
		return
	}

	// Cursor moved above viewport
	if m.CursorY < m.ViewportY {
		m.ViewportY = m.CursorY
	}

	// Cursor moved below viewport
	if m.CursorY >= m.ViewportY+usableHeight {
		m.ViewportY = m.CursorY - usableHeight + 1
	}

	if m.ViewportY < 0 {
		m.ViewportY = 0
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Handle window resize
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	// Handle Key press
	case tea.KeyMsg:
		// Get document projection
		lines := m.Doc.Project()

		if len(lines) == 0 {
			return m, nil
		}

		switch msg.String() {

		case "ctrl+c": // Exit editor
			return m, tea.Quit

		case "left": // Move cursor left
			if m.CursorX > 0 {
				m.CursorX--
			} else if m.CursorY > 0 {
				m.CursorY--
				m.CursorX = len(lines[m.CursorY])
			}

		case "right": // Move cursor right
			if m.CursorY < len(lines) {
				lineLen := len(lines[m.CursorY])
				if m.CursorX < lineLen {
					m.CursorX++
				}
			}

		case "up": // Move cursor up
			if m.CursorY > 0 {
				m.CursorY--
				lineLen := len(lines[m.CursorY])
				if m.CursorX > lineLen {
					m.CursorX = lineLen
				}
			}

		case "down": // Move cursor down
			if m.CursorY < len(lines)-1 {
				m.CursorY++
				lineLen := len(lines[m.CursorY])
				if m.CursorX > lineLen {
					m.CursorX = lineLen
				}
			}

		case "enter": // Create newline
			cursor := m.DocumentCursorIndex() // Get cursor index
			m.Doc.InsertAt(cursor, '\n')      // Insert newline character

			// Update cursor position
			m.CursorX = 0
			m.CursorY++

		case "backspace", "delete", "ctrl+h":
			if m.CursorX == 0 && m.CursorY == 0 {
				break
			}

			cursor := m.DocumentCursorIndex()
			m.Doc.DeleteAt(cursor)

			// Move cursor left
			if m.CursorX > 0 {
				m.CursorX--
			} else if m.CursorY > 0 {
				m.CursorY--
				m.CursorX = len(lines[m.CursorY])
			}

		case "ctrl+x": // Exit editor
			if err := m.Save(); err != nil {
				log.Println("Error saving file:", err)
			}
			return m, tea.Quit

		case "ctrl+o": // Save file
			if err := m.Save(); err != nil {
				log.Println("Error saving file:", err)
			} else {
				log.Println("File saved!")
			}

		default: // Write letter to file
			if len(msg.Runes) > 0 { // Write characters
				cursor := m.DocumentCursorIndex()    // Get cursor index
				m.Doc.InsertAt(cursor, msg.Runes[0]) // Insert character
				m.CursorX++                          // Increment cursor position
			}
		}

		m.clampViewport()
		/*
			case ServerMsg:
				m.Document.ApplyRemote(msg.Event)
				m.Editor.Rebuild(m.Document)
				return m, listenServer(m.serverCh)
		*/

	}

	return m, nil
}
