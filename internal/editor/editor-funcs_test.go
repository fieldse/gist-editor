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
	// TODO
}

func Test_replaceChunk(t *testing.T) {
	// TODO
}
