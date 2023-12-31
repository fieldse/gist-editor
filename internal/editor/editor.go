// Custom multline text entry widget
package editor

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

// CursorPosition returns the cursor position, in row/column format
func (m *MultiLineWidget) CursorPosition() shared.Position {
	return shared.Position{Row: m.CursorRow + 1, Col: m.CursorColumn + 1}
}

// GetSelection returns the current text selection and position.
// The Fyne entry widget counts position from 0,0.
// This function returns position from 1,1, to match standard editor conventions.
func (m *MultiLineWidget) GetSelection() shared.TextSelection {
	return shared.TextSelection{
		CursorPosition: m.CursorPosition(),
		Content:        m.SelectedText(),
		SelectionStart: m.SelectionStart(),
	}
}

// SelectedRowRange returns the row numbers of the current text selection.
// This returns start -> end, regardless if the selection is forwards or backwards.
// For example, if the current selection spans from rows 2 to 3, this would return 2,3
// IF the selection is reversed, it will still return 2,3
func (m *MultiLineWidget) SelectedRowRange() (int, int) {
	sel := m.GetSelection()
	selStart, curPos := sel.SelectionStart.Row, sel.CursorPosition.Row
	if curPos < selStart {
		return curPos, selStart
	}
	return selStart, curPos
}

// SelectionStart returns the selection cursor start position.
func (m *MultiLineWidget) SelectionStart() shared.Position {
	var row = reflect.ValueOf(m).Elem().FieldByName("selectRow").Int()
	var col = reflect.ValueOf(m).Elem().FieldByName("selectColumn").Int()
	return shared.Position{
		Row: int(row + 1),
		Col: int(col + 1),
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
	logger.Debug("cursor position: %+v", m.CursorPosition())
	logger.Debug("selection: %+v", m.GetSelection())
	start, end := m.SelectedRowRange()
	logger.Debug("selected rows: %d, %d", start, end)
}
