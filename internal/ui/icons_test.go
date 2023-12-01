// tests for icons loading functions
package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleIconNames = []string{"bold.png", "italic.png", "quote-block.png", "undo.png", "h1.png"}

func Test_loadIcon(t *testing.T) {
	for _, iconName := range exampleIconNames {
		data, err := loadIcon(iconName)
		assert.Nilf(t, err, "load icon file should succeed: %s", iconName)
		assert.NotEmptyf(t, data, "icon data should not be empty for %s", iconName)
	}
}
