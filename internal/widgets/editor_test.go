// Custom multline text entry widget
package widgets

import (
	"testing"

	"fyne.io/fyne/v2/app"
	"github.com/stretchr/testify/assert"
)

func TestNewMultilineWidget(t *testing.T) {
	a := app.New()
	w := a.NewWindow("test window")

	e := NewMultilineWidget("foo\nbar\nbaz\nbuz")
	w.SetContent(e) // required to run tests against a fyne widget
	assert.Equal(t, "foo\nbar\nbaz\nbuz", e.Text, "test should match expected")

	// Get cursor position
	curRow, curCol := e.CursorColumn, e.CursorRow
	assert.Equal(t, 0, curRow, "cursor row should be at starting position") // note that fyne entry widgets count position from zero
	assert.Equal(t, 0, curCol, "cursor col should be at starting position")
}
