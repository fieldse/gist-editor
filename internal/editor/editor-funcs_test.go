// Functions for altering the content of the text editor text
package editorfunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWithFoo(t *testing.T) {
	new := replaceWithFoo("some text")
	assert.Equalf(t, "foo", new, "should be replaced with 'foo'")
}

func Test_selectionToBold(t *testing.T) {
	// TODO
}

func Test_toLines(t *testing.T) {
	sampleText := "line 1\nline 2\nline 3\nline 4\nline 5"
	res := toLines(sampleText)
	assert.Lenf(t, res, 5, "should be 5 lines")
	assert.Equalf(t, res[0], "line 1", "line text should match expected: got %s", res[0])
	assert.Equalf(t, res[4], "line 5", "line text should match expected: got %s", res[4])
}

func Test_replaceChunk(t *testing.T) {
	orig := "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"
	expect := "example line 1\nexample line 2\nexample crazy chars 3\nexample line 4\nexample line 5"
	sel := TextSelection{Row: 3, Col: 9, Content: "line"}
	res, err := replaceChunk(orig, sel, "crazy chars")
	assert.Nilf(t, err, "replacechunk failed: %v", err)
	assert.Equalf(t, expect, res, "result doesn't match expected: got '%s'", res)
}
