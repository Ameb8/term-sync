package editor

import "strings"

// Adds cursor rune to line at given index
func renderLineWithCursor(line []rune, x int) string {
	if x > len(line) {
		x = len(line)
	}
	return string(line[:x]) + "|" + string(line[x:])
}

// Function to create string for bubble tea display
func (m *Model) View() string {
	var b strings.Builder

	maxLines := min(len(m.Editor.lines), m.Height)
	for y := 0; y < maxLines; y++ {
		line := m.Editor.lines[y]

		if y == m.CursorY {
			b.WriteString(renderLineWithCursor(line, m.CursorX))
		} else {
			b.WriteString(string(line))
		}

		b.WriteRune('\n')
	}

	return b.String()
}
