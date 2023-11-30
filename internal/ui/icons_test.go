// tests for icons loading functions
package ui

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleIconNames = []string{"bold.png", "italic.png", "quote-block.png", "undo.png", "h1.png"}

func Test_iconPaths(t *testing.T) {
	// Confirm the icon directory exists
	assert.DirExistsf(t, ICON_DIR, "icon directory does not exist")

	// Confirm a few icons exist
	for _, iconName := range exampleIconNames {
		fp := path.Join(ICON_DIR, iconName)
		assert.FileExistsf(t, fp, "icon file does not exist: %s", iconName)
	}
}

func Test_loadIcon(t *testing.T) {

	// Confirm a few icons exist
	for _, iconName := range exampleIconNames {
		data, err := loadIcon(iconName)
		assert.Nilf(t, err, "load icon file should succeed: %s", iconName)
		assert.NotEmptyf(t, data, "icon data should not be empty for %s", iconName)
	}
}
