// Miscellaneous app config stuff
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the various project directories exist
func Test_ProjectDirs(t *testing.T) {
	dirs := []string{
		PROJECT_ROOT,
		STATIC_DIR,
		STATIC_DIR,
	}
	for _, dir := range dirs {
		assert.DirExistsf(t, dir, "directory %s should exist", dir)
	}
}
