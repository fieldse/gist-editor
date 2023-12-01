// tests for icons loading functions
package ui

import (
	"path"
	"testing"

	"github.com/fieldse/gist-editor/internal/config"
	"github.com/stretchr/testify/assert"
)

var exampleIconNames = []string{"bold.png", "italic.png", "quote-block.png", "undo.png", "h1.png"}

var testIconDir = config.ICON_DIR

func Test_iconPaths(t *testing.T) {
	// Confirm the icon directory exists
	assert.DirExistsf(t, testIconDir, "icon directory does not exist")

	// Confirm a few icons exist
	for _, iconName := range exampleIconNames {
		fp := path.Join(testIconDir, iconName)
		assert.FileExistsf(t, fp, "icon file does not exist: %s", iconName)
	}
}

func Test_loadIcon(t *testing.T) {
	for _, iconName := range exampleIconNames {
		data, err := loadIcon(iconName)
		assert.Nilf(t, err, "load icon file should succeed: %s", iconName)
		assert.NotEmptyf(t, data, "icon data should not be empty for %s", iconName)
	}
}
