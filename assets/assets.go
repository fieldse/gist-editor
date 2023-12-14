// Loads embedded icon assets
package assets

import (
	"embed"
	"fmt"
	"path"

	"github.com/fieldse/gist-editor/internal/logger"
)

// IconMap is a map of the icon filename to their data as byte slices
var IconMap = make(map[string][]byte)

//go:embed icons
var iconAssets embed.FS
var iconNames = []string{
	"h1.png",
	"h2.png",
	"h3.png",
	"bold.png",
	"eraser.png",
	"italic.png",
	"link.png",
	"image.png",
	"quote-block.png",
	"code-block.png",
	"page-break.png",
	"undo.png",
	"redo.png",
}

// Preload all icon assets
func init() {
	err := preloadIcons()
	if err != nil {
		logger.Fatal("Failed to preload icons", err)
	}
}

// loadAsset loads an asset by name, storing it as a byte slice in the
// package-level IconMap variable
func loadAsset(iconName string) error {
	data, err := iconAssets.ReadFile(path.Join("icons", iconName))
	if err != nil {
		return fmt.Errorf("failed to load icon %s: %v", iconName, err)
	}
	IconMap[iconName] = data
	return nil
}

// preloadIcons preloads and embeds the icon assets, for use in the application toolbars
func preloadIcons() error {
	for _, iconName := range iconNames {
		err := loadAsset(iconName)
		if err != nil {
			return err
		}
	}
	return nil
}
