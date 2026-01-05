func main() {
	crdt := NewCRDTDocument()
	serverCh := make(chan ServerEvent)

	m := Model{
		crdt:     crdt,
		cursor:   Cursor{0, 0},
		viewport: Viewport{0, 0, 80, 24},
		serverCh: serverCh,
	}

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
