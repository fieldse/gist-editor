// Functions for altering the content of the text editor text
package editorfunctions

import (
	"fmt"
	"strings"

	"github.com/fieldse/gist-editor/internal/logger"
)

// TextSelection represents the position and content of the editor's current text selection
type TextSelection struct {
	Col     int
	Row     int
	Content string
}

// EditorFunctions is the set of Markdown syntax operations that can be performed
// on the editor text content
type EditorFunctions struct {
	GetText      func() string        // get the editor's current text and cursor position
	GetSelection func() TextSelection // get the editor's selected text and cursor position
	ReplaceText  func(string)         // replace the content of the editor
}

func (e EditorFunctions) New(getText func() string, getSelection func() TextSelection, replaceText func(string)) *EditorFunctions {
	return &EditorFunctions{
		GetText:      getText,
		GetSelection: getSelection,
		ReplaceText:  replaceText,
	}
}

// textOperation is any text manipulation operation against the editor text & selection
type textOperation func(string, TextSelection) (string, error)

// doTextOperation performs a text operation on the current text of the editor,
// replacing its content with the result.
func (e *EditorFunctions) doTextOperation(f textOperation) error {
	origText := e.GetText()
	selection := e.GetSelection()
	newText, err := f(origText, selection)
	if err != nil {
		logger.Error("text operation failed", err)
		return err
	}
	e.ReplaceText(newText)
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
	// TODO
}

// Checklist styles the current row as a checklist item
func (e *EditorFunctions) Checklist() {
	e.doTextOperation(rowToChecklistItem)
}

// Image uploads and inserts an image at the current location
func (e *EditorFunctions) Image() {
	// TODO
}

// QuoteBlock styles the current selection as a quote block
func (e *EditorFunctions) QuoteBlock() {
	// TODO
}

// CodeBlock styles the current selection as a code block
func (e *EditorFunctions) CodeBlock() {
	// TODO
}

// PageBreak inserts a page break at the current position
func (e *EditorFunctions) PageBreak() {
	// TODO
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
var rowPrefixes = []string{
	"# ", "## ", "### ", "#### ", "##### ", // headings
	" - [ ] ", " - [x] ", "- [ ] ", "- [x] ", // checklists
	" - ", // ul lists
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
// heading or list styling
// Examples:
//
//	(prefix: '#') 	strings "foo', "# foo", and " - foo" all become "# foo"
//	(prefix: ' - ') strings "foo", "# foo", and " - foo" all become " - foo"
//	(prefix: 'baz') strings "foo", "# foo", and " - foo" all become "bazfoo"
func replaceRowPrefix(rowNumber int, orig string, newPrefix string) (string, error) {
	row, err := getNthLine(rowNumber, orig)
	if err != nil {
		return "", err
	}
	newRow := newPrefix + stripPrefixes(row) // strip existing tags and append the new one
	return replaceNthLine(rowNumber, orig, newRow)
}

// rowToH1 adds an H1 styling prefix to the current row, replacing any existing style
func rowToH1(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, "# ")
}

// rowToH2 adds an H2 styling prefix to the current row, replacing any existing style
func rowToH2(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, "## ")
}

// rowToH3 adds an H3 styling prefix to the current row, replacing any existing style
func rowToH3(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, "### ")
}

// rowToH4 adds an H4 styling prefix to the current row, replacing any existing style
func rowToH4(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, "### ")
}

// rowToUL adds an undordered list style to the current row, replacing any existing style
func rowToUL(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, " - ")
}

// rowToChecklistItem adds an checklist style prefix to the current row, replacing any existing style
func rowToChecklistItem(orig string, selection TextSelection) (string, error) {
	return replaceRowPrefix(selection.Row, orig, " - [ ] ")
}

// getNthLine returns the Nth line of a piece of text, separated by newlines.
// (Note that lines start at 1, not zero)
// Returns error if N exceeds number of lines.
func getNthLine(n int, text string) (string, error) {
	asLines := toLines(text)
	if n > len(asLines) {
		return "", fmt.Errorf("line number %d exceeds lines in text", n)
	}
	return asLines[n-1], nil
}

// replaceNthLine replaces the Nth line of a piece of text with a new string.
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
	rowNum := sel.Row
	row, err := getNthLine(rowNum, orig)
	if err != nil {
		return "", err
	}

	// Get start and end character positions
	toReplace := sel.Content
	start := sel.Col - 1
	end := start + len(toReplace)

	// Sanity checks
	// -- cursor column position shouldn't exceed length of the row
	if end > len(row) {
		return "", fmt.Errorf("replace string failed: original cursor position exceeds content")
	}
	// -- row should contain the given string
	if !strings.Contains(row, toReplace) {
		return "", fmt.Errorf("replace string failed: original does not contain substring %s", toReplace)
	}

	pref := row[0:start]  // start point is current cursor position
	mid := row[start:end] // this should equal our current selection
	suffix := row[end:]   // end point is start + N chars (length of substring)

	// Sanity checks: middle chunk should be the current selection
	if mid != toReplace {
		return "", fmt.Errorf("current selection does not match given substring: selection is  '%s', but got '%s'", toReplace, mid)
	}

	// Replace the row with the substituted version
	newRow := pref + replaceWith + suffix
	return replaceNthLine(rowNum, orig, newRow)

}
