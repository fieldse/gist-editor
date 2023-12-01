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
func selectionToBold(entireText string, selection TextSelection) (string, error) {
	asLines := toLines(entireText)
	currentLine := asLines[selection.Row-1] // confirm this exists
	replaceWith := fmt.Sprintf("**%s**", selection.Content)
	return replaceChunk(currentLine, selection, replaceWith)
}

// toLines breaks the current text selection to lines
func toLines(text string) []string {
	return strings.Split(text, "\n")
}

// replaceChunk replaces a chunk of a line of text, from the starting position
func replaceChunk(s string, sel TextSelection, replaceWith string) (string, error) {

	// Sanity checks
	// -- cursor position shouldn't be greater than the length or the original string
	if sel.Col > len(s) {
		return "", fmt.Errorf("replace string failed: original string length shorter than expected")
	}

	// -- original should contain the given string
	if !strings.Contains(s, sel.Content) {
		return "", fmt.Errorf("replace string failed: original does not contain substring %s", sel.Content)
	}

	// Split into chunks
	start := sel.Col
	end := start + len(sel.Content)
	pref := s[0:start]  // start point is current cursor position
	mid := s[start:end] // this should equal our current selection
	suffix := s[end:]   // end point is start + N chars (length of substring)

	// Sanity checks: original should contain the given string
	if mid != sel.Content {
		return "", fmt.Errorf("current selection does not match expected substring: selection  '%s', but got '%s'", sel.Content, mid)
	}
	return pref + replaceWith + suffix, nil

}
