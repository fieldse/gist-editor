// The main file menu
package ui

import (
	"fyne.io/fyne/v2"
)

// // Returns a main File menu
func FileMenu(cfg *AppConfig) *fyne.MainMenu {
	// File menu
	openMenu := fyne.NewMenuItem("Open...", cfg.OpenFile)
	saveMenu := fyne.NewMenuItem("Save", cfg.SaveFile)
	saveAsMenu := fyne.NewMenuItem("Save as...", cfg.SaveFileAs)
	closeMenu := fyne.NewMenuItem("Close", cfg.CloseFile)

	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu, closeMenu)

	// Github settings & authentication settings
	githubTokenMenu := fyne.NewMenuItem("Github API Token", cfg.ShowGithubTokenModal)
	githubMenu := fyne.NewMenu("Github", githubTokenMenu)

	// Main app menu
	mainMenu := fyne.NewMainMenu(fileMenu, githubMenu)
	return mainMenu
}
