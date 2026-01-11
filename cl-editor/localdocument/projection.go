package localdocument

import "strings"

type LineProjection struct {
	lines        [][]rune
	cachedString string
	dirty        bool
}

func newLineProjection() *LineProjection {
	return &LineProjection{
		lines: [][]rune{[]rune{}},
		dirty: true,
	}
}

func (p *LineProjection) cursorToLineCol(cursor int) (line, col int) {
	pos := 0
	for i, l := range p.lines {
		if cursor <= pos+len(l) {
			return i, cursor - pos
		}
		pos += len(l) + 1 // +1 for '\n'
	}
	// fallback
	last := len(p.lines) - 1
	return last, len(p.lines[last])
}

func (p *LineProjection) reset() {
	p.lines = [][]rune{[]rune{}}
	p.cachedString = ""
	p.dirty = true
}

func (p *LineProjection) insert(cursor int, r rune) {
	line, col := p.cursorToLineCol(cursor)

	if r == '\n' {
		cur := p.lines[line]
		newLine := append([]rune{}, cur[col:]...)
		p.lines[line] = cur[:col]
		p.lines = append(
			p.lines[:line+1],
			append([][]rune{newLine}, p.lines[line+1:]...)...,
		)
	} else {
		l := p.lines[line]
		l = append(l[:col], append([]rune{r}, l[col:]...)...)
		p.lines[line] = l
	}
	p.dirty = true
}

func (p *LineProjection) delete(cursor int) {
	// Backspace at start does nothing
	if cursor <= 0 || cursor > p.len() {
		return
	}

	// Get index of character before cursor
	pos := cursor - 1
	line, col := p.cursorToLineCol(pos)

	// Delete character within line
	if col < len(p.lines[line]) {
		l := p.lines[line]
		p.lines[line] = append(l[:col], l[col+1:]...)
		p.dirty = true
		return
	}

	// Delete newline and merge lines
	if line < len(p.lines)-1 {
		p.lines[line] = append(p.lines[line], p.lines[line+1]...)
		p.lines = append(p.lines[:line+1], p.lines[line+2:]...)
		p.dirty = true
	}
}

func (p *LineProjection) string() string {
	if !p.dirty {
		return p.cachedString
	}

	var b strings.Builder
	for i, line := range p.lines {
		b.WriteString(string(line))
		if i < len(p.lines)-1 {
			b.WriteByte('\n')
		}
	}

	p.cachedString = b.String()
	p.dirty = false
	return p.cachedString
}

func (p *LineProjection) len() int {
	count := 0
	for _, line := range p.lines {
		count += len(line)
	}
	// newlines count as characters except after last line
	if len(p.lines) > 1 {
		count += len(p.lines) - 1
	}
	return count
}

func (p *LineProjection) lineCount() int {
	return len(p.lines)
}

func (p *LineProjection) getLines() [][]rune {
	return p.lines
}
