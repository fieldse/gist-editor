// Contains share types for the editor
package shared

// TextSelection represents the position and content of the editor's current text selection
type TextSelection struct {
	Position Position // cursor position, using rows and columns
	Content  string
}

// NewTextSelection returns a new TextSelection instance
func NewTextSelection(content string, row int, col int) TextSelection {
	return TextSelection{
		Position: Position{
			Row: row,
			Col: col,
		},
		Content: content,
	}
}

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