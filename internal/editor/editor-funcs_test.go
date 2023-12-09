// Functions for altering the content of the text editor text
package editorfunctions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleText = "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"

// Text selection from the above example text -- the word "line" from line 3
var selectionLineThreeWordTwo = TextSelection{Row: 3, Col: 9, Content: "line"}

// Text selection from the above example text -- from beginning of line 1 to end of line 2
var multiLineSelectionLines1and2 = TextSelection{Row: 2, Col: 15, Content: "example line 1\nexample line 2"}

// Multiline selection from the above example text -- from "line" in line 2 to "example" in line 3
var multiLineSelectionLines2and3 = TextSelection{Row: 3, Col: 8, Content: "line 2\nexample"}

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

func Test_getSelectedRows(t *testing.T) {
	// Example 1 - single line, line 3, partial text
	sel := selectionLineThreeWordTwo
	rows, err := getSelectedRows(exampleText, sel)
	assert.Nilf(t, err, "getSelectedRows failed: %v", err)
	require.Equalf(t, 1, len(rows), "selection 1 should be a single row")
	assert.Containsf(t, rows[0], sel.Content, "line should contain content: got %s", rows[0])
	assert.Equalf(t, "example line 3", rows[0], "line text should match expected: got %s", rows)

	// Example 2 - Lines 1-2, full text
	sel = multiLineSelectionLines1and2
	rows, err = getSelectedRows(exampleText, sel)
	assert.Nilf(t, err, "getSelectedRows failed: %v", err)
	require.Equalf(t, 2, len(rows), "selection should be two rows")

	// Rows should contain selected text
	assert.Containsf(t, "example line 1", rows[0], "row should contain expected text: got %s", rows[0])
	assert.Containsf(t, "example line 2", rows[1], "row should contain expected text: got %s", rows[1])

	// Example 3 - Lines 2-3, partial text
	sel = multiLineSelectionLines2and3
	rows, err = getSelectedRows(exampleText, sel)
	assert.Nilf(t, err, "getSelectedRows failed: %v", err)
	require.Equalf(t, 2, len(rows), "selection should be two rows")
	// Rows should contain selected text

	assert.Containsf(t, "example line 1", rows[0], "row should contain expected text: got %s", rows[0])
	assert.Containsf(t, "example line 2", rows[1], "row should contain expected text: got %s", rows[1])

	// Over line count should return an error
	_, err = getSelectedRows(exampleText, selectionLineThreeWordTwo)
	assert.NotNilf(t, err, "getSelectedRows should fail on outside of line range")
}

func Test_replaceChunk(t *testing.T) {
	expect := "example line 1\nexample line 2\nexample crazy chars 3\nexample line 4\nexample line 5"
	res, err := replaceChunk(exampleText, selectionLineThreeWordTwo, "crazy chars")
	assert.Nilf(t, err, "replacechunk failed: %v", err)
	assert.Equalf(t, expect, res, "result doesn't match expected: got '%s'", res)
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
		res, err := rowToH1(x.s, TextSelection{Col: 3, Row: 3, Content: "foo"})
		assert.Nil(t, err)
		assert.Equalf(t, x.expect, res, "expected %s, got %s", x.expect, res)
	}
}

func Test_getSelectionRange(t *testing.T) {
	// TODO: make test cases here
	var cases = []struct {
		sel         TextSelection
		reverse     bool
		expect      string
		shouldErr   bool // should the function return an error : ie, string out of bounds
		shouldMatch bool // should the match expected result
	}{
		{sel: selectionLineThreeWordTwo, reverse: false, expect: "line", shouldErr: false, shouldMatch: true},
		{sel: selectionLineThreeWordTwo, reverse: false, expect: "foo", shouldErr: false, shouldMatch: false}, // different word, should not match
		{sel: selectionLineThreeWordTwo, reverse: true, expect: "line", shouldErr: false, shouldMatch: false}, // in reverse, should return the wrong string
	}
	fmt.Printf("==== ORIGINAL TEXT: \n'%s'\n", exampleText)
	for i, x := range cases {
		fmt.Printf("==== TEST CASE [%d/%d] ==== : %+v\n", i+1, len(cases), x)
		res, err := getSelectionRange(exampleText, x.sel, x.reverse)
		fmt.Printf("=== [debug] got result: '%+v'\n", res)

		if x.shouldErr {
			assert.NotNil(t, err, "should return error: got %v", err)
		} else {
			assert.Nilf(t, err, "should not return error: got %v", err)
		}
		if x.shouldMatch {
			assert.Equalf(t, x.expect, res, "expected '%s', got '%s'", x.expect, res)
		} else {
			assert.NotEqual(t, x.expect, res, "should not get '%s', but got '%s'", x.expect, res)
		}
	}
}
