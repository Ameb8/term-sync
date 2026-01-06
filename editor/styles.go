package editor

import "github.com/charmbracelet/lipgloss"

var (
	textStyle = lipgloss.NewStyle()

	cursorStyle = lipgloss.NewStyle().
			Reverse(true)

	lineStyle = lipgloss.NewStyle().
			Width(0) // width set dynamically in View()
)
