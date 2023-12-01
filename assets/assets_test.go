package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringAssets(t *testing.T) {
	x := loadStringAsset()
	expect := "Hello world!"
	assert.Equalf(t, expect, x, "example text expected to equal %s", expect)
}

func Test_nestedAssets(t *testing.T) {
	x := loadNestedStringAsset()
	expect := "Hello world!"
	assert.Equalf(t, expect, x, "example text expected to equal %s", expect)
}

func Test_loadFileAsset(t *testing.T) {
	x := loadFileAsset()
	content, err := x.ReadFile("somedir/test.txt")
	assert.Nilf(t, err, "loadFileAsset failed")
	assert.NotEmptyf(t, content, "content should not be empty")
}
