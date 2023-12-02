// Functions for altering the content of the text editor text
package editorfunctions

import (
	"fmt"
	"strings"
)

// TextSelection represents the position and content of the editor's current text selection
type TextSelection struct {
	Col     int
	Row     int
	Content string
}

// replaceWithFoo replaces the entire text with "foo"
func replaceWithFoo(text string) string {
	return "foo"
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
