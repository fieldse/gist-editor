package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runUI(t *testing.T) {

	a := AppConfig{}.New()
	a.MakeUI()

	assert.NotNil(t, a.App, "app should be initialized")
	assert.NotNil(t, a.MainWindow, "main window should exist")
	assert.NotNil(t, a.Editor, "editor window should exist")
	assert.NotNil(t, a.GithubSettingsWindow, "GithubSettingsWindow window should exist")
	assert.NotNil(t, a.ListWindow, "ListWindow window should exist")
}

func Test_ReadConfig(t *testing.T) {
	// TODO  -- add test for loading app config
	a := AppConfig{}.New()
	a.MakeUI()
	res := a.ReadConfig()
	assert.Nil(t, res, "read config should succeed")
}
