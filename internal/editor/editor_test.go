// Custom multline text entry widget
package editor

import (
	"testing"

	"fyne.io/fyne/v2/app"
	"github.com/stretchr/testify/assert"
)

// Return a test MultiLineWidget fixture
func newTestMultiLine() *MultiLineWidget {
	a := app.New()
	w := a.NewWindow("test window")
	e := NewMultilineWidget("foo\nbar\nbaz\nbuz")
	w.SetContent(e) // required to run tests against a fyne widget
	return e
}

func TestNewMultilineWidget(t *testing.T) {
	e := newTestMultiLine()
	assert.Equal(t, "foo\nbar\nbaz\nbuz", e.Text, "test should match expected")
}

func Test_cursorPosition(t *testing.T) {
	e := newTestMultiLine()

	// Get cursor position
	curRow, curCol := e.CursorColumn, e.CursorRow
	assert.Equal(t, 0, curRow, "cursor row should be at starting position") // note that fyne entry widgets count position from zero
	assert.Equal(t, 0, curCol, "cursor col should be at starting position")
}

func Test_SelectionStart(t *testing.T) {
	e := newTestMultiLine()

	// Get selection start position
	pos := e.SelectionStart()

	// With no selection active, fyne returns 0,0
	assert.Equal(t, 0, pos.Row, "selection row should be negative")
	assert.Equal(t, 0, pos.Col, "selection col should be negative")
}
