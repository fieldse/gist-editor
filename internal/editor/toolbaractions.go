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

// doTextOperation performs a text operation on the current text of an editor,
// replacing its content with the result.
func doTextOperation(f textOperation, e *MultiLineWidget) error {
	origText := e.Content()
	selection := e.GetSelection()
	newText, err := f(origText, selection)
	if err != nil {
		logger.Error("text operation failed", err)
		return err
	}
	e.SetContent(newText)
	return nil
}

// H1 styles the current selection as H1
func (e *toolbarActions) H1() {
	doTextOperation(rowToH1, e.editor)
}

// H2 styles the current selection as H2
func (e *toolbarActions) H2() {
	doTextOperation(rowToH2, e.editor)
}

// H3 styles the current selection as H3
func (e *toolbarActions) H3() {
	doTextOperation(rowToH3, e.editor)
}

// H4 styles the current selection as H4
func (e *toolbarActions) H4() {
	doTextOperation(rowToH4, e.editor)
}

// Bold styles the current selection as Bold
func (e *toolbarActions) Bold() {
	doTextOperation(selectionToBold, e.editor)
}

// Italic styles the current selection as Italic
func (e *toolbarActions) Italic() {
	doTextOperation(selectionToItalic, e.editor)
}

// Stikethrough styles the current selection as Stikethrough
func (e *toolbarActions) Stikethrough() {
	doTextOperation(selectionToStrikethrough, e.editor)
}

// Link styles the current selection as a link
func (e *toolbarActions) Link() {
	doTextOperation(selectionToStrikethrough, e.editor)
}

// UL styles the current row as unordered list item
func (e *toolbarActions) UL() {
	doTextOperation(rowToUL, e.editor)
}

// OL styles the current row as ordered list item
func (e *toolbarActions) OL() {
	doTextOperation(rowToOL, e.editor)
}

// Checklist styles the current row as a checklist item
func (e *toolbarActions) Checklist() {
	doTextOperation(rowToChecklistItem, e.editor)
}

// Image uploads and inserts an image at the current location
func (e *toolbarActions) Image() {
	logger.Debug("placeholder for Image action")
	// TODO
}

// QuoteBlock styles the current selection as a quote block
func (e *toolbarActions) QuoteBlock() {
	doTextOperation(rowsToQuoteBlock, e.editor)
}

// CodeBlock styles the current selection as a code block
func (e *toolbarActions) CodeBlock() {
	doTextOperation(rowsToCodeBlock, e.editor)
}

// PageBreak inserts a page break at the current position
func (e *toolbarActions) PageBreak() {
	doTextOperation(insertPageBreak, e.editor)
}

// ClearFormatting clears any row styling (eg: h1, h2, checklist, list item, quote, code block)
// from the selected text content
func (e *toolbarActions) ClearFormatting() {
	doTextOperation(clearFormatting, e.editor)
}

// Undo the most recent changes to the text content
func (e *toolbarActions) Undo() {
	e.editor.Undo()
}

// Redo the most recent changes to the text content
func (e *toolbarActions) Redo() {
	e.editor.Redo()
}
