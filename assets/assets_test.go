package assets

import (
	"testing"

	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_loadIconAsset(t *testing.T) {
	x := iconAssets
	content, err := x.ReadFile("icons/bold.png")
	assert.Nilf(t, err, "load icon asset failed")
	assert.NotEmptyf(t, content, "content should not be empty")
	logger.Debug("icon content loaded: \n%v", content)
}

func Test_preloadIcons(t *testing.T) {
	err := preloadIcons()
	assert.Nil(t, err, "load icon asset failed")

	for _, name := range []string{"h1.png", "h2.png", "h3.png", "bold.png"} {
		content, ok := IconMap[name]
		assert.Truef(t, ok, "%s should exist in IconMap", name)
		assert.NotEmptyf(t, content, "content should not be empty")
	}
}
