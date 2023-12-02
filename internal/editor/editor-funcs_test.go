// Functions for altering the content of the text editor text
package editorfunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleText = "example line 1\nexample line 2\nexample line 3\nexample line 4\nexample line 5"

// Text selection from the above example text -- the world "line" from line 3
var exampleSelection = TextSelection{Row: 3, Col: 9, Content: "line"}

func TestReplaceWithFoo(t *testing.T) {
	new := replaceWithFoo("some text")
	assert.Equalf(t, "foo", new, "should be replaced with 'foo'")
}

func Test_selectionToBold(t *testing.T) {
	r, err := selectionToBold(exampleText, exampleSelection)
	expect := "example line 1\nexample line 2\nexample **line** 3\nexample line 4\nexample line 5"
	assert.Nil(t, err)
	assert.Equalf(t, expect, r, "replaced text should equal expected: got %v instead", r)
}

func Test_toLines(t *testing.T) {
	res := toLines(exampleText)
	assert.Lenf(t, res, 5, "should be 5 lines")
	assert.Equalf(t, res[0], "example line 1", "line text should match expected: got %s", res[0])
	assert.Equalf(t, res[4], "example line 5", "line text should match expected: got %s", res[4])
}

func Test_replaceChunk(t *testing.T) {
	expect := "example line 1\nexample line 2\nexample crazy chars 3\nexample line 4\nexample line 5"
	res, err := replaceChunk(exampleText, exampleSelection, "crazy chars")
	assert.Nilf(t, err, "replacechunk failed: %v", err)
	assert.Equalf(t, expect, res, "result doesn't match expected: got '%s'", res)
}
