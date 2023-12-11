// Functions for altering the content of the text editor text
package editor

import (
	"fmt"
	"strings"

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
	logger.Debug("placeholder for OL action")
	// TODO
}

// Checklist styles the current row as a checklist item
func (e *toolbarActions) Checklist() {
	logger.Debug("placeholder for Checklist action")
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
	logger.Debug("placeholder for CodeBlock action")
	// TODO
}

// PageBreak inserts a page break at the current position
func (e *toolbarActions) PageBreak() {
	// TODO
}

// Undo the most recent changes to the text content
func (e *toolbarActions) Undo() {
	e.editor.Undo()
}

// Redo the most recent changes to the text content
func (e *toolbarActions) Redo() {
	e.editor.Redo()
}

// startAndEndPositions returns start and end positions of a selection
func startAndEndPositions(t TextSelection) (Position, Position) {
	curPos, selPos := t.CursorPosition, t.SelectionStart
	if curPos.Row > selPos.Row {
		return selPos, curPos
	}
	if curPos.Row < selPos.Row {
		return curPos, selPos
	}
	// equal row, so compare columns
	if curPos.Col > selPos.Col {
		return selPos, curPos
	}
	return curPos, selPos
}

// isMultiline checks if a text selection spans multiple rows
func isMultiline(t TextSelection) bool {
	return t.SelectionStart.Row == t.CursorPosition.Row
}

// replaceSelection replaces the selected segment of a text with the given string.
// This will fail if the selection is multiple line
func replaceSelection(text string, selection TextSelection, replaceWith string) (string, error) {
	if isMultiline(selection) {
		return "", fmt.Errorf("multiple line selection not supported")
	}
	return replaceChunk(text, selection, replaceWith)
}

// selectionToBold adds Markdown bold styling to the current text selection:
// (ie: "foo" becomes "**foo**"")
func selectionToBold(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("**%s**", selection.Content)
	return replaceSelection(orig, selection, replaceWith)
}

// selectionToItalic adds Markdown italic styling to the current text selection:
// (ie: "foo" becomes "_foo_"")
func selectionToItalic(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("_%s_", selection.Content)
	return replaceSelection(orig, selection, replaceWith)
}

// selectionToStrikethrough adds Markdown strikethrough styling to the current text
// selection: 	(ie: "foo" becomes "~~foo~~"")
func selectionToStrikethrough(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("~~%s~~", selection.Content)
	return replaceSelection(orig, selection, replaceWith)
}

// These patterns will be assume as styling at the beginning of a row, and will be
// replaced during replaceRowPrefix operations.
// FIXME: replace this simplistic pattern check with regex
var rowPrefixes = []string{
	"# ", "## ", "### ", "#### ", "##### ", // headings
	" - [ ] ", " - [x] ", "- [ ] ", "- [x] ", // checklists
	"- ", " - ", // ul lists
	"1. ", "2. ", "3. ", "4. ", "5. ", "6. ", "7. ", "8. ", "9.", "10. ", // ordered lists
}

// stripPrefixes strips common Markdown styling characters, such as h1, h2, bullets, checklist
func stripPrefixes(s string) string {
	for _, x := range rowPrefixes {
		if strings.HasPrefix(s, x) {
			return strings.TrimPrefix(s, x)
		}
	}
	return s
}

// replaceRowPrefixes adds a prefix to a each row in a range from a text
// This replaces any existing Markdown heading or list styling on the row.
//
//	text 		the original editor contents
//	startRow  	the first row of the range (starts from 1)
//	endRow  	the last row of the range  (starts from 1)
//	newPrefix	the new style prefix to add to each row
//
// Examples:
//
//	(Given rows "foo', "# foo", and " - foo")
//		prefix: '#' 		all rows become "# foo"
//		prefix: ' - '  		all rows become " - foo"
//		prefix: 'baz'  		all rows become "baz foo"
func replaceRowPrefixes(text string, startRow int, endRow int, newPrefix string) (string, error) {
	asRows := toLines(text)
	// TODO - add checks against the row lengths
	for i := startRow; i <= endRow; i++ {
		row := asRows[i+1]
		// TODO -- improve this by not mutating the original set
		asRows[i+1] = replacePrefix(row, newPrefix)
	}
	return strings.Join(asRows, "\n"), nil
}

// prefixSelectedRows replaces the prefix on all selected rows of a text
func prefixSelectedRows(text string, sel TextSelection, newPrefix string) (string, error) {
	curPos := sel.CursorPosition.Row
	selPos := sel.SelectionStart.Row
	var start, end int
	if curPos > selPos {
		start = selPos
		end = curPos
	} else {
		start = curPos
		end = selPos
	}
	return replaceRowPrefixes(text, start, end, newPrefix)
}

// replacePrefix adds a styling prefix to a text string, replacing any existing
// Markdown heading or list styling.
func replacePrefix(text string, newPrefix string) string {
	return newPrefix + stripPrefixes(text) // strip existing tags and append the new one
}

// rowToH1 adds an H1 styling prefix to the current row, replacing any existing style
func rowToH1(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, "# ")
}

// rowToH2 adds an H2 styling prefix to the current row, replacing any existing style
func rowToH2(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, "## ")
}

// rowToH3 adds an H3 styling prefix to the current row, replacing any existing style
func rowToH3(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, "### ")
}

// rowToH4 adds an H4 styling prefix to the current row, replacing any existing style
func rowToH4(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, "### ")
}

// rowToUL adds an undordered list style to the current row, replacing any existing style
func rowToUL(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, " - ")
}

// rowToChecklistItem adds an checklist style prefix to the current row, replacing any existing style
func rowToChecklistItem(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, " - [ ] ")
}

// toLines breaks the current text selection to lines
func toLines(text string) []string {
	return strings.Split(text, "\n")
}

// replaceChunk replaces current selection in a piece of text with a given string
// TODO: unify this with replaceSelection
func replaceChunk(text string, sel TextSelection, replaceWith string) (string, error) {
	asLines := toLines(text)
	selected := sel.Content
	start, end := startAndEndPositions(sel)
	rowNum := sel.CursorPosition.Row - 1
	row := asLines[rowNum]
	// -- row should contain the given string
	if !strings.Contains(row, selected) {
		return "", fmt.Errorf("replace string failed: original does not contain substring %s", selected)
	}

	pref := row[0:start.Col]      // the chunk before the selection
	mid := row[start.Col:end.Col] // this should equal our current selection
	suffix := row[end.Col:]       // the chunk after the selection

	// Sanity checks: middle chunk should be the current selection
	if mid != selected {
		return "", fmt.Errorf("current selection does not match given substring: selection is  '%s', but got '%s'", selected, mid)
	}

	// Replace the row with the substituted version
	asLines[rowNum] = pref + replaceWith + suffix
	return strings.Join(asLines, "\n)"), nil
}
