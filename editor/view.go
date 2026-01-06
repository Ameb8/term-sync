package editor

import "strings"

// Adds cursor rune to line at given index
func renderLineWithCursor(line []rune, x int) string {
	// Prevent cursor from extending past line end
	if x > len(line) {
		x = len(line)
	}

	return string(line[:x]) + "|" + string(line[x:])
}

// Function to create string for bubble tea display
func (m *Model) View() string {
	var b strings.Builder                          // Initialize display buffer
	maxLines := min(len(m.Editor.lines), m.Height) // Set max lines

	// Render all lines
	for y := 0; y < maxLines; y++ {
		line := m.Editor.lines[y] // Get current line

		var rendered string

		if y == m.CursorY { // Render cursor line
			rendered = renderLineWithCursor(line, m.CursorX)
		} else { // Render normally
			rendered = string(line)
		}

		// Clear rest of line
		if len(rendered) < m.Width {
			rendered += strings.Repeat(" ", m.Width-len(rendered))
		}

		// Write line to display buffer
		b.WriteString(rendered)
		b.WriteRune('\n')
	}

	return b.String()
}
