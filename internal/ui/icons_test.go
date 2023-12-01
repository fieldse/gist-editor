// tests for icons loading functions
package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadIcon(t *testing.T) {
	icons, err := ToolbarIcons{}.Load()
	assert.Nilf(t, err, "load icon data")
	assert.NotEmptyf(t, icons.BoldIcon.Content(), "icon data should not be empty")
}
