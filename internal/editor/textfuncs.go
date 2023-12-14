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

// startAndEndRows returns the start and end row numbers of a selection
func startAndEndRows(t TextSelection) (int, int) {
	start, end := startAndEndPositions(t)
	return start.Row, end.Row
}

// isMultiline checks if a text selection spans multiple rows
func isMultiline(t TextSelection) bool {
	return t.SelectionStart.Row != t.CursorPosition.Row
}

// insertPageBreak inserts a Markdown page break (-----) before the current selection row
func insertPageBreak(text string, sel TextSelection) (string, error) {
	return insertRowBeforeSelection(text, sel, "-----")
}

// insertRowBeforeSelection insert a row of text in before the beginning of the
// current selection row
func insertRowBeforeSelection(text string, sel TextSelection, toInsert string) (string, error) {
	startRow, _ := startAndEndRows(sel)
	rows := toLines(text)
	if startRow > len(rows) {
		return "", fmt.Errorf("text selection exceeds row count")
	}
	i := startRow - 1 // Row count starts at 1
	if i < 0 {
		return "", fmt.Errorf("row index below zero")
	}
	var newRows []string = append(rows[:i], append([]string{toInsert}, rows[i:]...)...)
	return strings.Join(newRows, "\n"), nil
}

// wrapRows inserts text string as rows both before and after the
// current selection
func wrapRows(text string, sel TextSelection, toInsert string) (string, error) {
	var newRows []string
	startRow, endRow := startAndEndRows(sel)
	rows := toLines(text)

	// Start and end indexes
	var pre, post int
	pre = startRow - 1
	post = endRow + 1
	// if startRow == endRow { // if we're on a single line selection, ensure we append the last item
	// 	post = pre
	// }

	// Validate row counts
	if pre < 0 || post > len(rows)+1 {
		return "", fmt.Errorf("row index out of range")
	}
	// Insert preceding row
	newRows, err := insertToSlice(rows, toInsert, pre)
	if err != nil {
		return "", err
	}
	// Insert trailing row
	newRows, err = insertToSlice(newRows, toInsert, post)
	if err != nil {
		return "", err
	}

	return strings.Join(newRows, "\n"), nil
}

// insertToSlice inserts an element into a slice of strings at the given index
func insertToSlice(arr []string, s string, index int) ([]string, error) {
	if index > len(arr)+1 { // out of range
		return []string{}, fmt.Errorf("insert to slice: index out of range")
	}
	return append(arr[0:index], append([]string{s}, arr[index:]...)...), nil
}

// replaceSelection replaces the selected segment of a text with the given string.
// This will fail if the selection is multiple line
func replaceSelection(text string, sel TextSelection, replaceWith string) (string, error) {
	if isMultiline(sel) {
		return "", fmt.Errorf("multiple line selection not supported")
	}
	asLines := toLines(text)
	selected := sel.Content
	start, end := startAndEndPositions(sel)
	rowNum := sel.CursorPosition.Row - 1
	startChar := start.Col - 1
	endChar := end.Col - 1
	row := asLines[rowNum]
	// -- row should contain the given string
	if !strings.Contains(row, selected) {
		return "", fmt.Errorf("replace string failed: original does not contain substring %s", selected)
	}

	pref := row[0:startChar]      // the chunk before the selection
	mid := row[startChar:endChar] // this should equal our current selection
	suffix := row[endChar:]       // the chunk after the selection

	// Sanity checks: middle chunk should be the current selection
	if mid != selected {
		return "", fmt.Errorf("current selection does not match given substring: selection is  '%s', but got '%s'", selected, mid)
	}

	// Replace the row with the substituted version
	asLines[rowNum] = pref + replaceWith + suffix
	return strings.Join(asLines, "\n"), nil
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
	"- [ ] ", "- [x] ", // checklists
	"- ",                                                                 // ul lists
	"1. ", "2. ", "3. ", "4. ", "5. ", "6. ", "7. ", "8. ", "9.", "10. ", // ordered lists
}

// stripPrefixes strips common Markdown styling characters, such as h1, h2, bullets, checklist
func stripPrefixes(s string) string {
	stripped := strings.Trim(s, " ")
	for _, x := range rowPrefixes {
		if strings.HasPrefix(stripped, x) {
			return strings.TrimPrefix(stripped, x)
		}
	}
	return s
}

// replaceRowPrefixes adds a prefix to a each row in a range from a text
// This replaces any existing Markdown heading or list styling on the row.
//
//	Examples:
//
// Given text, selection spanning three rows:
//
//	# foo
//	- bar
//	3. baz
//
// Results:
//
//	prefix: '#' 		result: "# foo\n# bar\n# baz"
//	prefix: ' - '  		result: " - foo\n - bar\n - baz"
//	prefix: '1. '  		result: "1. foo\n1. bar\n1. baz"
func prefixSelectedRows(text string, sel TextSelection, newPrefix string) (string, error) {
	asRows := toLines(text)
	startRow, endRow := startAndEndRows(sel)
	// Iterate by row, replacing the prefix
	for i := startRow; i <= endRow; i++ {
		row := asRows[i-1]
		prefixed := replacePrefix(row, newPrefix)
		asRows[i-1] = prefixed
	}
	return strings.Join(asRows, "\n"), nil
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
	return prefixSelectedRows(orig, selection, "#### ")
}

// rowToUL adds an undordered list style to the current row, replacing any existing style
func rowToUL(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, " - ")
}

// rowToOL adds an ordered list style to the current row, replacing any existing style
func rowToOL(orig string, selection TextSelection) (string, error) {
	// TODO: make this intelligently actually introduce numbered items.
	return prefixSelectedRows(orig, selection, "1. ")
}

// rowToChecklistItem adds an checklist style prefix to the current row, replacing any existing style
func rowToChecklistItem(orig string, selection TextSelection) (string, error) {
	return prefixSelectedRows(orig, selection, " - [ ] ")
}

// rowsToCodeBlock wraps the current selection in code blocks style
func rowsToCodeBlock(orig string, selection TextSelection) (string, error) {
	return wrapRows(orig, selection, "```")
}

// toLines breaks the current text selection to lines
func toLines(text string) []string {
	return strings.Split(text, "\n")
}
