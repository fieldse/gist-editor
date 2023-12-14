// Markdown styling and editing actions for the editor toolbar
package editor

import (
	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/fieldse/gist-editor/internal/shared"
)

type TextSelection = shared.TextSelection
type Position = shared.Position

// toolbarActions is the set of Markdown syntax operations that can be performed
// on the editor text content
type toolbarActions struct {
	editor *MultiLineWidget
}

func newToolbarActions(m *MultiLineWidget) *toolbarActions {
	return &toolbarActions{
		editor: m,
	}
}

// textOperation is any text manipulation operation against the editor text & selection
type textOperation func(string, TextSelection) (string, error)

// doTextOperation performs a text operation on the current text of the editor,
// replacing its content with the result.
func (e *toolbarActions) doTextOperation(f textOperation) error {
	origText := e.editor.Content()
	selection := e.editor.GetSelection()
	newText, err := f(origText, selection)
	if err != nil {
		logger.Error("text operation failed", err)
		return err
	}
	e.editor.SetContent(newText)
	return nil
}

// H1 styles the current selection as H1
func (e *toolbarActions) H1() {
	e.doTextOperation(rowToH1)
}

// H2 styles the current selection as H2
func (e *toolbarActions) H2() {
	e.doTextOperation(rowToH2)
}

// H3 styles the current selection as H3
func (e *toolbarActions) H3() {
	e.doTextOperation(rowToH3)
}

// H4 styles the current selection as H4
func (e *toolbarActions) H4() {
	e.doTextOperation(rowToH4)
}

// Bold styles the current selection as Bold
func (e *toolbarActions) Bold() {
	e.doTextOperation(selectionToBold)
}

// Italic styles the current selection as Italic
func (e *toolbarActions) Italic() {
	e.doTextOperation(selectionToItalic)
}

// Stikethrough styles the current selection as Stikethrough
func (e *toolbarActions) Stikethrough() {
	e.doTextOperation(selectionToStrikethrough)
}

// Link styles the current selection as a link
func (e *toolbarActions) Link() {
	e.doTextOperation(selectionToStrikethrough)
}

// UL styles the current row as unordered list item
func (e *toolbarActions) UL() {
	e.doTextOperation(rowToUL)
}

// OL styles the current row as ordered list item
func (e *toolbarActions) OL() {
	e.doTextOperation(rowToOL)
}

// Checklist styles the current row as a checklist item
func (e *toolbarActions) Checklist() {
	e.doTextOperation(rowToChecklistItem)
}

// Image uploads and inserts an image at the current location
func (e *toolbarActions) Image() {
	logger.Debug("placeholder for Image action")
	// TODO
}

// QuoteBlock styles the current selection as a quote block
func (e *toolbarActions) QuoteBlock() {
	logger.Debug("placeholder for QuoteBlock action")
	// TODO
}

// CodeBlock styles the current selection as a code block
func (e *toolbarActions) CodeBlock() {
	e.doTextOperation(rowsToCodeBlock)
}

// PageBreak inserts a page break at the current position
func (e *toolbarActions) PageBreak() {
	e.doTextOperation(insertPageBreak)
}

// Undo the most recent changes to the text content
func (e *toolbarActions) Undo() {
	e.editor.Undo()
}

// Redo the most recent changes to the text content
func (e *toolbarActions) Redo() {
	e.editor.Redo()
}
