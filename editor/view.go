package editor

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Regex pattern for identifying ANS escape codes
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// Compute visible-character width of column
func visualLen(s string) int {
	return len([]rune(ansiRegex.ReplaceAllString(s, "")))
}

// Adds cursor rune to line at given index
func (m *Model) renderCursorLine(line []rune) string {
	x := m.CursorX
	if x > len(line) {
		x = len(line)
	}

	var b strings.Builder

	// Before cursor
	if x > 0 {
		b.WriteString(string(line[:x]))
	}

	// Cursor cell
	if x < len(line) {
		b.WriteString(cursorStyle.Render(string(line[x])))
	} else {
		b.WriteString(cursorStyle.Render(" "))
	}

	// After cursor
	if x+1 < len(line) {
		b.WriteString(string(line[x+1:]))
	}

	return b.String()
}

// Function to create string for bubble tea display
func (m *Model) View() string {
	if m.Height == 0 || m.Width == 0 {
		return ""
	}

	var lines []string
	maxLines := min(len(m.Editor.lines), m.Height)

	for y := 0; y < maxLines; y++ {
		line := m.Editor.lines[y]

		var rendered string

		if y == m.CursorY {
			rendered = m.renderCursorLine(line)
		} else {
			rendered = string(line)
		}

		// Force line width (pads or truncates safely)
		rendered = lipgloss.NewStyle().
			Width(m.Width).
			Render(rendered)

		lines = append(lines, rendered)
	}

	// Clear remaining screen lines
	blank := lipgloss.NewStyle().
		Width(m.Width).
		Render("")

	for len(lines) < m.Height {
		lines = append(lines, blank)
	}

	// Add status bar to bottom of viewport
	status := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("Ctrl+O: Save | Ctrl+X: Quit")
	lines = append(lines, status)

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
