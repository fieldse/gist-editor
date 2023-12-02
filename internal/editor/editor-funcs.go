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

// toLines breaks the current text selection to lines
func toLines(text string) []string {
	return strings.Split(text, "\n")
}

// replaceChunk replaces current selection in a piece of text with a given string
func replaceChunk(orig string, sel TextSelection, replaceWith string) (string, error) {
	// Extract row to edit
	asLines := toLines(orig)
	rowNum := sel.Row - 1 // character positions start at 1,1, not 0,0
	row := asLines[rowNum]

	// Get start and end character positions
	toReplace := sel.Content
	start := sel.Col - 1
	end := start + len(toReplace)

	// Sanity checks
	// -- cursor column position shouldn't exceed length of the row
	if end > len(row) {
		return "", fmt.Errorf("replace string failed: original cursor position exceeds content")
	}

	// -- cursor row position shouldn't exceed number of original rows
	if sel.Row > len(asLines)+1 {
		return "", fmt.Errorf("replace string failed: original cursor row exceeds content")
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
	asLines[rowNum] = newRow
	return strings.Join(asLines, "\n"), nil

}
