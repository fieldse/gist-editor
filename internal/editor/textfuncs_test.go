// Functions for altering the content of the text editor text
package editor

import (
	"fmt"
	"testing"

	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/stretchr/testify/assert"
)

var exampleText = "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"

// Empty text selection -- cursor at 1,1
var emptyTextSelection = TextSelection{
	CursorPosition: Position{Row: 1, Col: 1},
	SelectionStart: Position{Row: 1, Col: 1},
	Content:        "",
}

// Text selection from the above example text -- the word "line" from line 3
var selectionLineThreeWordTwo = TextSelection{
	CursorPosition: Position{Row: 3, Col: 13},
	SelectionStart: Position{Row: 3, Col: 9},
	Content:        "line",
}

// Text selection from the above example text -- from beginning of line 1 to end of line 2
var multiLineSelectionLines1and2 = TextSelection{
	CursorPosition: Position{Row: 2, Col: 15},
	SelectionStart: Position{Row: 1, Col: 1},
	Content:        "example line 1\nexample line 2",
}

// Multiline selection from the above example text -- from "line" in line 2 to "example" in line 3
var multiLineSelectionLines2and3 = TextSelection{
	SelectionStart: Position{Row: 2, Col: 9},
	CursorPosition: Position{Row: 3, Col: 8},
	Content:        "line 2\nexample",
}

func Test_startAndEndRows(t *testing.T) {
	x, y := startAndEndRows(selectionLineThreeWordTwo)
	assert.Equal(t, x, 3, "expect start row to be 3: got %d", x)
	assert.Equal(t, y, 3, "expect start row to be 3: got %d", y)

	x, y = startAndEndRows(multiLineSelectionLines1and2)
	assert.Equal(t, x, 1, "expect start row to be 1: got %d", x)
	assert.Equal(t, y, 2, "expect start row to be 2: got %d", y)

	x, y = startAndEndRows(multiLineSelectionLines2and3)
	assert.Equal(t, x, 2, "expect start row to be 2: got %d", x)
	assert.Equal(t, y, 3, "expect start row to be 3: got %d", y)
}

func Test_rowToHeading(t *testing.T) {
	cases := []string{
		"example line 1\n# example line 2\n# example line 3\nexample line 4\nexample line 5",
		"example line 1\n - example line 2\n - example line 3\nexample line 4\nexample line 5",
		"example line 1\n - [ ] example line 2\n - [ ] example line 3\nexample line 4\nexample line 5",
		"example line 1\n#### example line 2\n#### example line 3\nexample line 4\nexample line 5",
	}
	for _, c := range cases {

		// Rows 2 and 3

		// ...to H1
		r, err := rowToH1(c, multiLineSelectionLines2and3)
		expect := "example line 1\n# example line 2\n# example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)

		// ...to H2
		r, err = rowToH2(c, multiLineSelectionLines2and3)
		expect = "example line 1\n## example line 2\n## example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)

		// ...to H3
		r, err = rowToH3(c, multiLineSelectionLines2and3)
		expect = "example line 1\n### example line 2\n### example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)

		// ...to H4
		r, err = rowToH4(c, multiLineSelectionLines2and3)
		expect = "example line 1\n#### example line 2\n#### example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
	}
}

func Test_rowToListItem(t *testing.T) {
	cases := []string{
		"example line 1\n# example line 2\n# example line 3\nexample line 4\nexample line 5",
		"example line 1\n - example line 2\n - example line 3\nexample line 4\nexample line 5",
		"example line 1\n - [ ] example line 2\n - [ ] example line 3\nexample line 4\nexample line 5",
		"example line 1\n#### example line 2\n#### example line 3\nexample line 4\nexample line 5",
	}
	for _, c := range cases {

		// Rows 2 and 3

		// ...to UL
		r, err := rowToUL(c, multiLineSelectionLines2and3)
		expect := "example line 1\n - example line 2\n - example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)

		// ...to checklist
		r, err = rowToChecklistItem(c, multiLineSelectionLines2and3)
		expect = "example line 1\n - [ ] example line 2\n - [ ] example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)

		// ...to OL item
		r, err = rowToOL(c, multiLineSelectionLines2and3)
		expect = "example line 1\n1. example line 2\n1. example line 3\nexample line 4\nexample line 5"
		assert.Nil(t, err)
		assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
	}
}

func Test_insertPageBreak(t *testing.T) {

	var cases = []struct {
		sel    TextSelection
		expect string
	}{
		// Insert before single line selection row 3
		{
			sel:    selectionLineThreeWordTwo,
			expect: "example line 1\nexample line 2\n-----\nexample line 3\nexample line 4\nexample line 5",
		},
		// Insert before multiline selection rows 2-3
		{
			sel:    multiLineSelectionLines2and3,
			expect: "example line 1\n-----\nexample line 2\nexample line 3\nexample line 4\nexample line 5",
		},
		// Insert before empty text selection
		{
			sel:    emptyTextSelection,
			expect: "-----\nexample line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5",
		},
	}

	for _, c := range cases {
		r, err := insertPageBreak(exampleText, c.sel)
		assert.Nil(t, err)
		assert.Equalf(t, c.expect, r, "insert page break: expected '%s', got '%s'", c.expect, r)
	}
}

func Test_insertRowBeforeSelection(t *testing.T) {

	var cases = []struct {
		sel    TextSelection
		expect string
	}{
		// Insert before single line selection row 3
		{
			sel:    selectionLineThreeWordTwo,
			expect: "example line 1\nexample line 2\nFOOBAR\nexample line 3\nexample line 4\nexample line 5",
		},
		// Insert before multiline selection rows 2-3
		{
			sel:    multiLineSelectionLines2and3,
			expect: "example line 1\nFOOBAR\nexample line 2\nexample line 3\nexample line 4\nexample line 5",
		},
		// Insert before empty text selection
		{
			sel:    emptyTextSelection,
			expect: "FOOBAR\nexample line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5",
		},
	}

	for _, c := range cases {
		r, err := insertRowBeforeSelection(exampleText, c.sel, "FOOBAR")
		assert.Nil(t, err)
		assert.Equalf(t, c.expect, r, "insert row before selection: expected '%s', got '%s'", c.expect, r)
	}
}

func Test_selectionToBold(t *testing.T) {
	r, err := selectionToBold(exampleText, selectionLineThreeWordTwo)
	expect := "example line 1\nexample line 2\nexample **line** 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_selectionToItalic(t *testing.T) {
	r, err := selectionToItalic(exampleText, selectionLineThreeWordTwo)
	expect := "example line 1\nexample line 2\nexample _line_ 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_selectionToStrikethrough(t *testing.T) {
	r, err := selectionToStrikethrough(exampleText, selectionLineThreeWordTwo)
	expect := "example line 1\nexample line 2\nexample ~~line~~ 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_toLines(t *testing.T) {
	res := toLines(exampleText)
	assert.Lenf(t, res, 5, "should be 5 lines")
	assert.Equalf(t, res[0], "example line 1", "line text should match expected: got %s", res[0])
	assert.Equalf(t, res[4], "example line 5", "line text should match expected: got %s", res[4])
}

func Test_replaceSelection(t *testing.T) {
	expect := "example line 1\nexample line 2\nexample crazy chars 3\nexample line 4\nexample line 5"
	res, err := replaceSelection(exampleText, selectionLineThreeWordTwo, "crazy chars")
	assert.Nilf(t, err, "replaceSelection failed: %v", err)
	assert.Equalf(t, expect, res, "result doesn't match expected: got '%s'", res)
}

func Test_isMultiline(t *testing.T) {
	s1 := TextSelection{
		SelectionStart: Position{Row: 1, Col: 1},
		CursorPosition: Position{Row: 1, Col: 12},
		Content:        "hello world",
	}
	s2 := TextSelection{
		SelectionStart: Position{Row: 1, Col: 1},
		CursorPosition: Position{Row: 2, Col: 3},
		Content:        "hello world\nhi",
	}
	assert.Falsef(t, isMultiline(s1), "should be single line: %v", s1)
	assert.Truef(t, isMultiline(s2), "should be multiline: %v", s2)
}

func Test_stripPrefixes(t *testing.T) {
	var cases = []struct {
		s      string
		expect string
	}{
		{s: "foo", expect: "foo"},
		{s: "# foo", expect: "foo"},
		{s: "## foo", expect: "foo"},
		{s: "## foo", expect: "foo"},
		{s: "### foo", expect: "foo"},
		{s: "#### foo", expect: "foo"},
		{s: "- foo", expect: "foo"},
		{s: " - foo", expect: "foo"},
		{s: " - [ ] foo", expect: "foo"},
		{s: "----bar", expect: "----bar"},
		{s: "#bar", expect: "#bar"},
	}
	for _, x := range cases {
		res := stripPrefixes(x.s)
		assert.Equalf(t, x.expect, res, "expected %s, got %s", x.expect, res)
	}
}

func Test_prefixSelectedRows(t *testing.T) {

	// CASE 1: Change first three rows, with no existing prefixes
	// (ie: plain format, no headings, etc.)

	// Original text will be something like this:
	text1 := "line 1\nline 2\nline 3\nline 4"

	// Selection will be the first three lines, ending at first character of line 3
	sel1 := TextSelection{
		SelectionStart: Position{Row: 1, Col: 1},
		CursorPosition: Position{Row: 3, Col: 1},
		Content:        "line 1\nline 2",
	}

	// We're going to prefix each row of the selection with this:
	prefix := "# "

	// Expecting the result to change the first three rows, but not the last.
	expectCase := fmt.Sprintf("%sline 1\n%sline 2\n%sline 3\nline 4", prefix, prefix, prefix)

	res, err := prefixSelectedRows(text1, sel1, prefix)
	assert.Nil(t, err)
	assert.Equalf(t, expectCase, res, "expected %s, got %s", expectCase, res)

	// CASE 2: Change first three rows, with existing style prefixes
	// (ie: h1, list item, checklist item, ordered list item.)

	// Original text will be something like this:
	case2Items := []struct {
		text    string
		content string
	}{
		// h1 styled
		{
			text:    "# line 1\n# line 2\n# line 3\nline 4",
			content: "# line 1\n# line 2",
		},
		// list item styled
		{
			text:    " - line 1\n - line 2\n - line 3\nline 4",
			content: " - line 1\n - line 2",
		},
		// checklist item styled
		{
			text:    " - [ ] line 1\n - [ ] line 2\n - [ ] line 3\nline 4",
			content: " - [ ] line 1\n - [ ] line 2",
		},
		// checklist item styled with no preceding spaces
		{
			text:    "- [ ] line 1\n- [ ] line 2\n- [ ] line 3\nline 4",
			content: "- [ ] line 1\n- [ ] line 2",
		},
		// ordered list styled
		{
			text:    "1. line 1\n2. line 2\n3. line 3\nline 4",
			content: "1. line 1\n2. line 2",
		},
	}

	for i, c := range case2Items {
		logger.Debug("=== TEST CASE [%d/%d] -- '%s'", i+1, len(case2Items), c.content)
		// Expecting the result to change the first three rows, but not the last.

		// Selection will be the first three lines, ending at first character of line 3
		sel := TextSelection{
			SelectionStart: Position{Row: 1, Col: 1},
			CursorPosition: Position{Row: 3, Col: 1},
			Content:        c.content,
		}
		res, err := prefixSelectedRows(c.text, sel, prefix)
		assert.Nil(t, err)
		assert.Equalf(t, expectCase, res, "expected '%s', got '%s'", expectCase, res)
	}

}
