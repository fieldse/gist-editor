// Custom multline text entry widget
package widgets

import (
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
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
