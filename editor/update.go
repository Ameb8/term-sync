package editor

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c": // Exit editor
			return m, tea.Quit

		case "left": // Move cursor left
			m.CursorX = max(0, m.CursorX-1)

		case "right": // Move cursor right
			m.CursorX++

		case "up": // Move cursor up
			if m.CursorY > 0 {
				m.CursorY--
				lineLen := len(m.Editor.lines[m.CursorY])
				if m.CursorX > lineLen {
					m.CursorX = lineLen
				}
			}

		case "down": // Move cursor down
			if m.CursorY < len(m.Editor.lines)-1 {
				m.CursorY++
				lineLen := len(m.Editor.lines[m.CursorY])
				if m.CursorX > lineLen {
					m.CursorX = lineLen
				}
			}

		case "enter": // Create newline
			cursor := m.DocumentCursorIndex() // Get cursor index
			m.Doc.InsertAt(cursor, '\n')      // Insert newline character
			m.Editor.Rebuild(m.Doc)           // Update editor state

			// Update cursor position
			m.CursorX = 0
			m.CursorY++

		default: // Write letter to file
			if len(msg.Runes) > 0 { // Write characters
				cursor := m.DocumentCursorIndex()    // Get cursor index
				m.Doc.InsertAt(cursor, msg.Runes[0]) // Insert character
				m.Editor.Rebuild(m.Doc)              // Update editor state
				m.CursorX++                          // Increment cursor position
			}
		}

		/*
			case ServerMsg:
				m.Document.ApplyRemote(msg.Event)
				m.Editor.Rebuild(m.Document)
				return m, listenServer(m.serverCh)
		*/

	}

	return m, nil
}
