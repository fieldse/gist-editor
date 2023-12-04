// Functions for altering the content of the text editor text
package editorfunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleText = "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"

// Text selection from the above example text -- the world "line" from line 3
var exampleSelection = TextSelection{Row: 3, Col: 9, Content: "line"}

func Test_selectionToBold(t *testing.T) {
	r, err := selectionToBold(exampleText, exampleSelection)
	expect := "example line 1\nexample line 2\nexample **line** 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_selectionToItalic(t *testing.T) {
	r, err := selectionToItalic(exampleText, exampleSelection)
	expect := "example line 1\nexample line 2\nexample _line_ 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_selectionToStrikethrough(t *testing.T) {
	r, err := selectionToStrikethrough(exampleText, exampleSelection)
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

func Test_getNthLine(t *testing.T) {
	res, err := getNthLine(1, exampleText)
	assert.Nilf(t, err, "getNthLine failed: %v", err)
	assert.Equalf(t, "example line 1", res, "line text should match expected: got %s", res)

	res, err = getNthLine(5, exampleText)
	assert.Nilf(t, err, "getNthLine failed: %v", err)
	assert.Equalf(t, "example line 5", res, "line text should match expected: got %s", res)

	// Over line count should return an error
	_, err = getNthLine(6, exampleText)
	assert.NotNilf(t, err, "getNthLine should fail on outside of line range")
}

func Test_replaceChunk(t *testing.T) {
	expect := "example line 1\nexample line 2\nexample crazy chars 3\nexample line 4\nexample line 5"
	res, err := replaceChunk(exampleText, exampleSelection, "crazy chars")
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
