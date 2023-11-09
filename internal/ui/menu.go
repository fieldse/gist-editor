// The main file menu
package ui

import "fyne.io/fyne/v2"

// // Returns a main File menu
func FileMenu(cfg *AppConfig) *fyne.MainMenu {
	openMenu := fyne.NewMenuItem("Open...", cfg.openFileFunc)
	closeMenu := fyne.NewMenuItem("Close", cfg.closeFileFunc)
	exitMenu := fyne.NewMenuItem("Exit", cfg.Exit)

	fileMenu := fyne.NewMenu("File", openMenu, closeMenu, exitMenu)
	mainMenu := fyne.NewMainMenu(fileMenu)
	return mainMenu
}
