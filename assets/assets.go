// test loading static assets
package assets

import (
	"embed"
)

//go:embed test.txt
var exampleText string

//go:embed somedir/test.txt
var nestedText string

//go:embed somedir
var subdirAssets embed.FS

// loadStringAsset test loading a string asset directly from file content
func loadStringAsset() string {
	return exampleText
}

// loadNestedStringAsset test loading a string asset directly from file content in a subdirectory
func loadNestedStringAsset() string {
	return nestedText
}

// loadFileAsset test loading a file asset using embed.FS
func loadFileAsset() embed.FS {
	return subdirAssets
}
