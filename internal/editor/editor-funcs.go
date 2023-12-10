// Functions for altering the content of the text editor text
package editorfunctions

import (
	"fmt"
	"strings"

	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/fieldse/gist-editor/internal/shared"
	"github.com/fieldse/gist-editor/internal/widgets"
)

type TextSelection = shared.TextSelection
type Position = shared.Position

// EditorFunctions is the set of Markdown syntax operations that can be performed
// on the editor text content
type EditorFunctions struct {
	editor         *widgets.MultiLineWidget
	getContent     func() string
	setContent     func(string)
	getSelection   func() shared.TextSelection
	selectionStart func() shared.Position
}

func NewEditorFunctions(m *widgets.MultiLineWidget) *EditorFunctions {
	return &EditorFunctions{
		editor:         m,
		getContent:     m.Content,
		setContent:     m.SetText,
		getSelection:   m.GetSelection,
		selectionStart: m.SelectionStart,
	}
}

// textOperation is any text manipulation operation against the editor text & selection
type textOperation func(string, TextSelection) (string, error)

// doTextOperation performs a text operation on the current text of the editor,
// replacing its content with the result.
func (e *EditorFunctions) doTextOperation(f textOperation) error {
	origText := e.getContent()
	selection := e.getSelection()
	newText, err := f(origText, selection)
	if err != nil {
		logger.Error("text operation failed", err)
		return err
	}
	e.setContent(newText)
	return nil
}

// H1 styles the current selection as H1
func (e *EditorFunctions) H1() {
	e.doTextOperation(rowToH1)
}

// H2 styles the current selection as H2
func (e *EditorFunctions) H2() {
	e.doTextOperation(rowToH2)
}

// H3 styles the current selection as H3
func (e *EditorFunctions) H3() {
	e.doTextOperation(rowToH3)
}

// H4 styles the current selection as H4
func (e *EditorFunctions) H4() {
	e.doTextOperation(rowToH4)
}

// Bold styles the current selection as Bold
func (e *EditorFunctions) Bold() {
	e.doTextOperation(selectionToBold)
}

// Italic styles the current selection as Italic
func (e *EditorFunctions) Italic() {
	e.doTextOperation(selectionToItalic)
}

// Stikethrough styles the current selection as Stikethrough
func (e *EditorFunctions) Stikethrough() {
	e.doTextOperation(selectionToStrikethrough)
}

// Link styles the current selection as a link
func (e *EditorFunctions) Link() {
	e.doTextOperation(selectionToStrikethrough)
}

// UL styles the current row as unordered list item
func (e *EditorFunctions) UL() {
	e.doTextOperation(rowToUL)
}

// OL styles the current row as ordered list item
func (e *EditorFunctions) OL() {
	logger.Debug("placeholder for OL action")
	// TODO
}

// Checklist styles the current row as a checklist item
func (e *EditorFunctions) Checklist() {
	logger.Debug("placeholder for Checklist action")
	e.doTextOperation(rowToChecklistItem)
}

// Image uploads and inserts an image at the current location
func (e *EditorFunctions) Image() {
	logger.Debug("placeholder for Image action")
	// TODO
}

// QuoteBlock styles the current selection as a quote block
func (e *EditorFunctions) QuoteBlock() {
	logger.Debug("placeholder for QuoteBlock action")
	// TODO
}

// CodeBlock styles the current selection as a code block
func (e *EditorFunctions) CodeBlock() {
	logger.Debug("placeholder for CodeBlock action")
	// TODO
}

// PageBreak inserts a page break at the current position
func (e *EditorFunctions) PageBreak() {
	// TODO
}

// Undo the most recent changes to the text content
func (e *EditorFunctions) Undo() {
	e.editor.Undo()
}

// Redo the most recent changes to the text content
func (e *EditorFunctions) Redo() {
	e.editor.Redo()
}

// selectionToBold adds Markdown bold styling to the current text selection:
// (ie: "foo" becomes "**foo**"")
func selectionToBold(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("**%s**", selection.Content)
	return replaceChunk(orig, selection, replaceWith)
}

// selectionToItalic adds Markdown italic styling to the current text selection:
// (ie: "foo" becomes "_foo_"")
func selectionToItalic(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("_%s_", selection.Content)
	return replaceChunk(orig, selection, replaceWith)
}

// selectionToStrikethrough adds Markdown strikethrough styling to the current text
// selection: 	(ie: "foo" becomes "~~foo~~"")
func selectionToStrikethrough(orig string, selection TextSelection) (string, error) {
	replaceWith := fmt.Sprintf("~~%s~~", selection.Content)
	return replaceChunk(orig, selection, replaceWith)
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

// replaceRowPrefix adds a styling prefix to the current row, replacing any existing
// heading or list styling.
//
//	rowNumber  	the cursor row position, as returned from the editor. Counts from 1)
//	orig 		the original editor contents
//
// Examples:
//
//	(prefix: '#') 	strings "foo', "# foo", and " - foo" all become "# foo"
//	(prefix: ' - ') strings "foo", "# foo", and " - foo" all become " - foo"
//	(prefix: 'baz') strings "foo", "# foo", and " - foo" all become "bazfoo"
func replaceRowPrefix(sel TextSelection, orig string, newPrefix string) (string, error) {
	var newRows []string
	rows, err := getSelectedRows(orig, sel)
	if err != nil {
		return "", err
	}
	for _, row := range rows {
		rowNum := 0                              // FIXME -- get row number
		newRow := newPrefix + stripPrefixes(row) // strip existing tags and append the new one
		replaceNthLine(rowNum, orig, newRow)
	}
	return strings.Join(newRows, "\n"), nil
}

// rowToH1 adds an H1 styling prefix to the current row, replacing any existing style
func rowToH1(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, "# ")
}

// rowToH2 adds an H2 styling prefix to the current row, replacing any existing style
func rowToH2(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, "## ")
}

// rowToH3 adds an H3 styling prefix to the current row, replacing any existing style
func rowToH3(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, "### ")
}

// rowToH4 adds an H4 styling prefix to the current row, replacing any existing style
func rowToH4(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, "### ")
}

// rowToUL adds an undordered list style to the current row, replacing any existing style
func rowToUL(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, " - ")
}

// rowToChecklistItem adds an checklist style prefix to the current row, replacing any existing style
func rowToChecklistItem(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection, orig, " - [ ] ")
}

// getSelectionRange returns the preceding or trailing N characters from a text,
// relative to the position of the cursor.
// Parameters:
//
//	text    	the original text
//	sel     	the selection and cursor position
//	reverse		move backward in the selection from the cursor
func getSelectionRange(text string, sel TextSelection, reverse bool) (string, error) {
	// Cursor position
	rowNum := sel.Position.Row
	colNum := sel.Position.Col

	asLines := toLines(text)
	// Get absolute position of the cursor relative to the original string
	// start at the column character number
	var charIndex int = colNum - 1
	if rowNum > len(asLines)+1 {
		return "", fmt.Errorf("selection row %d out of range", rowNum)
	}
	for r := 0; r < rowNum; r++ {
		// Add row lengths until we have reached the last row
		if rowNum != r {
			charIndex = charIndex + len(asLines[r]) + 1 // we add a character for the newline character
		}
	}
	// we should now have the absolute index of the selection cursor, relative to the original text
	var startChar, endChar int
	if reverse { // If it's in reverse order, the starting position precedes the cursor by some characters
		startChar = charIndex - len(sel.Content)
		endChar = charIndex
	} else { // Otherwise, the selection end follows the cursor
		startChar = charIndex
		endChar = charIndex + len(sel.Content)

	}
	if endChar > len(text) {
		return "", fmt.Errorf("selection exceeds text length: character %d, text length is %d", endChar, len(text))
	}
	return text[startChar:endChar], nil
}

// getSelectedRows returns the row(s) of a text selection, separated by newlines.
func getSelectedRows(text string, sel TextSelection) ([]string, error) {
	var rows []string

	asLines := toLines(text)

	// If the selection goes backwards more than the current line,
	// have a multi-line selection
	numlines := strings.Count(sel.Content, "\n") + 1

	// FIXME: this selection could be either forwards or backwards.
	var isForwards bool = true

	// Iterate (forward/backward) for N lines
	j := sel.Position.Row - 1 // index of the current row in the lines array
	for i := 0; i < numlines; i++ {
		currentRow := asLines[j]
		if isForwards {
			// If user has selected forwards, the cursor is at the end of the selection. Therefore move backwards by line.
			rows = append([]string{currentRow}, rows...)
			j = j - 1
		} else { // If cursor selection is going backwards, the cursor is at the start of the selection.
			// Therefore we move forwards by line
			rows = append(rows, currentRow)
			j = j + 1
		}
	}
	return rows, nil
}

// replaceNthLine replaces the Nth line of a piece of text with a new string.
// Line count starts at 1, not at zero.
// Returns error if N exceeds number of lines.
func replaceNthLine(n int, text string, replaceWith string) (string, error) {
	asLines := toLines(text)
	if n > len(asLines) {
		return "", fmt.Errorf("line number %d exceeds lines in text", n)
	}
	asLines[n-1] = replaceWith // row counts start at 1
	return strings.Join(asLines, "\n"), nil
}

// toLines breaks the current text selection to lines
func toLines(text string) []string {
	return strings.Split(text, "\n")
}

// replaceChunk replaces current selection in a piece of text with a given string
func replaceChunk(orig string, sel TextSelection, replaceWith string) (string, error) {
	// Extract row to edit
	rowNum := sel.Position.Row
	rows, err := getSelectedRows(orig, sel)
	if err != nil {
		return "", err
	}

	// FIXME -- handle multiple rows
	if len(rows) != 1 {
		return "", fmt.Errorf("fixme: replaceChunk needs to handle multiple rows")
	}
	row := rows[0]

	// Get start and end character positions
	selected := sel.Content
	// If there is a selection, the cursor position is at the end.
	// So, we count backwards to find the start position
	end := sel.Position.Col - 1  // Cursor position, zero based
	start := end - len(selected) // Counting backwards to find the start of the selection

	// Sanity checks
	// -- cursor column position shouldn't exceed length of the row
	if end > len(row) {
		return "", fmt.Errorf("replace string failed: original cursor position exceeds content")
	}
	// -- row should contain the given string
	if !strings.Contains(row, selected) {
		return "", fmt.Errorf("replace string failed: original does not contain substring %s", selected)
	}

	pref := row[0:start]  // the chunk before the selection
	mid := row[start:end] // this should equal our current selection
	suffix := row[end:]   // the chunk after the selection

	// Sanity checks: middle chunk should be the current selection
	if mid != selected {
		return "", fmt.Errorf("current selection does not match given substring: selection is  '%s', but got '%s'", selected, mid)
	}

	// Replace the row with the substituted version
	newRow := pref + replaceWith + suffix
	return replaceNthLine(rowNum, orig, newRow)
}
