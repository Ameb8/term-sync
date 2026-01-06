package editor

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "left":
			m.CursorX = max(0, m.CursorX-1)

		case "right":
			m.CursorX++

		case "enter":
			cursor := m.DocumentCursorIndex()
			m.Doc.InsertAt(cursor, '\n')
			m.Editor.Rebuild(m.Doc)
			m.CursorX = 0
			m.CursorY++

		default:
			if len(msg.Runes) > 0 {
				cursor := m.DocumentCursorIndex()
				m.Doc.InsertAt(cursor, msg.Runes[0])
				m.Editor.Rebuild(m.Doc)
				m.CursorX++
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
