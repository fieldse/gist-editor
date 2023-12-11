// Functions for altering the content of the text editor text
package editor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleText = "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"

// Text selection from the above example text -- the word "line" from line 3
var selectionLineThreeWordTwo = TextSelection{
	CursorPosition: Position{Row: 3, Col: 13},
	SelectionStart: Position{Row: 3, Col: 9},
	Content:        "line"}

// Text selection from the above example text -- from beginning of line 1 to end of line 2
var multiLineSelectionLines1and2 = TextSelection{
	CursorPosition: Position{Row: 2, Col: 15},
	SelectionStart: Position{Row: 1, Col: 1},
	Content:        "example line 1\nexample line 2"}

// Multiline selection from the above example text -- from "line" in line 2 to "example" in line 3
var multiLineSelectionLines2and3 = TextSelection{
	SelectionStart: Position{Row: 2, Col: 9},
	CursorPosition: Position{Row: 3, Col: 8},
	Content:        "line 2\nexample"}

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

func Test_rowToH1(t *testing.T) {
	var cases = []struct {
		s      string
		expect string
		sel    TextSelection
	}{
		{s: "line 1\nline 2\nfoo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n# foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n## foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n## foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n### foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n#### foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n - foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n - [ ] foo", expect: "line 1\nline 2\n# foo"},
		{s: "line 1\nline 2\n----bar", expect: "line 1\nline 2\n# ----bar"},
		{s: "line 1\nline 2\n#bar", expect: "line 1\nline 2\n# #bar"},
	}
	for _, x := range cases {
		res, err := rowToH1(x.s, TextSelection{CursorPosition: Position{Col: 3, Row: 3}, Content: "foo"})
		assert.Nil(t, err)
		assert.Equalf(t, x.expect, res, "expected %s, got %s", x.expect, res)
	}
}
