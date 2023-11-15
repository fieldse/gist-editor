// The main file menu
package ui

import (
	"fyne.io/fyne/v2"
	"github.com/fieldse/gist-editor/internal/logger"
)

// Returns a main File menu, and a function to toggle Save allowed
func FileMenu(cfg *AppConfig) (*fyne.MainMenu, func(bool)) {
	// File menu
	openMenu := fyne.NewMenuItem("Open...", cfg.OpenFile)
	saveMenu := fyne.NewMenuItem("Save", cfg.SaveFile)
	saveAsMenu := fyne.NewMenuItem("Save as...", cfg.SaveFileAs)
	saveMenu.Disabled = true   // save menus disabled until we have an open file
	saveAsMenu.Disabled = true // save menus disabled until we have an open file
	closeMenu := fyne.NewMenuItem("Close", cfg.CloseFile)
	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu, closeMenu)

	// TODO - Keyboard Shortcuts

	// Function to toggle "Save" allowed on the File menu
	setCanSave := func(b bool) {
		logger.Debug("toggle canSave: %v", b)
		saveMenu.Disabled = !b
		saveAsMenu.Disabled = !b
	}

	// Github settings & authentication settings
	githubTokenMenu := fyne.NewMenuItem("Github API Token", cfg.ShowGithubTokenModal)
	githubMenu := fyne.NewMenu("Github", githubTokenMenu)

	// Main app menu
	mainMenu := fyne.NewMainMenu(fileMenu, githubMenu)
	return mainMenu, setCanSave
}
