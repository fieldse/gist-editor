// The main file menu
package ui

import (
	"fyne.io/fyne/v2"
)

// // Returns a main File menu
func FileMenu(cfg *AppConfig) *fyne.MainMenu {
	openMenu := fyne.NewMenuItem("Open...", openFile)
	saveMenu := fyne.NewMenuItem("Save", saveFile)
	saveAsMenu := fyne.NewMenuItem("Save as...", saveFileAs)
	closeMenu := fyne.NewMenuItem("Close", closeFile)
	exitMenu := fyne.NewMenuItem("Quit", cfg.Exit)

	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu, closeMenu, exitMenu)
	mainMenu := fyne.NewMainMenu(fileMenu)
	return mainMenu
}
