// Contains share types for the editor
package shared

// TextSelection represents the position and content of the editor's current text selection
type TextSelection struct {
	CursorPosition Position // cursor position, using rows and columns
	Content        string
	SelectionStart Position // beginning position of the selection
}

// HasSelection returns true if there is a current text selection
func (s TextSelection) HasSelection() bool { return len(s.Content) > 0 }

// TODO -- this might be useful for clarifying the exact position of the character selection
type AbsoluteCharacterRange struct {
	Start int
	End   int
}

// Position represents a cursor position in a multi-row selection.
// Both row and column values start from 1
type Position struct {
	Col int
	Row int
}
