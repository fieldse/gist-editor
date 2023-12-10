// Custom multline text entry widget
package widgets

import (
	"reflect"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/fieldse/gist-editor/internal/shared"
)

// NewMultilineWidget returns a new custom multiline widget
func NewMultilineWidget(content string) *MultiLineWidget {
	e := &MultiLineWidget{}
	e.ExtendBaseWidget(e)
	e.SetText(content)
	e.MultiLine = true
	e.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
	return e
}

// MultiLineWidget is a custom multiline entry widget, with improved cursor functions
type MultiLineWidget struct {
	widget.Entry
}

// Content returns the editor's text content
func (m *MultiLineWidget) Content() string {
	return m.Text
}

// SetContent replaces the editor's text content
func (m *MultiLineWidget) SetContent(t string) {
	m.SetText(t)
}

// GetSelection returns the current text selection and position.
// The Fyne entry widget counts position from 0,0.
// This function returns position from 1,1, to match standard editor conventions.
func (m *MultiLineWidget) GetSelection() shared.TextSelection {
	row := m.CursorRow + 1
	col := m.CursorColumn + 1
	content := m.SelectedText()
	return shared.NewTextSelection(content, row, col)
}

// SelectionStart returns the selection cursor start position.
// Returns -1,-1 if there is no selection
func (m *MultiLineWidget) SelectionStart() shared.Position {
	var row = reflect.ValueOf(m).Elem().FieldByName("selectRow").Int()
	var col = reflect.ValueOf(m).Elem().FieldByName("selectColumn").Int()
	return shared.Position{
		Row: int(row),
		Col: int(col),
	}
}

// ContentRows returns the content of the widget, broken into rows by newline
func (m *MultiLineWidget) ContentRows() []string {
	return strings.Split(m.Text, "\n")
}

// RowCount returns the number of rows in the text content
func (m *MultiLineWidget) RowCount() int {
	return strings.Count(m.Text, "\n")
}

// Undo the most recent changes to the text content
func (m *MultiLineWidget) Undo() {
	logger.Debug("placeholder for Undo action")
	// TODO
}

// Redo the most recent changes to the text content
func (m *MultiLineWidget) Redo() {
	logger.Debug("placeholder for Redo action")
	// TODO
}

// Debug current text selection
func (m *MultiLineWidget) DebugTextSelection() {
	logger.Debug("content: %s", m.Text)
}