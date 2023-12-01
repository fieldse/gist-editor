// Miscellaneous app config stuff
package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/fieldse/gist-editor/internal/logger"
)

var PROJECT_ROOT, STATIC_DIR, ICON_DIR string

// Set project root and various directories
func init() {
	PROJECT_ROOT = projectRoot()
	STATIC_DIR = path.Join(PROJECT_ROOT, "internal", "static")
	ICON_DIR = path.Join(PROJECT_ROOT, "internal", "static", "icons")
}

// projectRoot returns the absolute path of the project root directory
func projectRoot() string {
	root, err := filepath.Abs(path.Join("..", ".."))
	if err != nil {
		logger.Error("failed to get project root", err)
		os.Exit(1)
	}
	return root
}
