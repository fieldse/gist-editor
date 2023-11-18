// Base window for the app
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// MainWindow is the main app window with menu & methods for create & show
type MainWindow struct {
	Window     fyne.Window
	menu       *fyne.MainMenu
	SetCanSave func(bool) // toggle whether Save / SaveAs is allowed in the main menu
}

// Show shows the main window and starts the application
func (m MainWindow) ShowAndRun() {
	m.Window.ShowAndRun()
}

// Close Closes the main window and starts the application
func (m MainWindow) Close() {
	m.Window.Close()
}

// New intitializes the main window UI and returns a MainWindow instance
func (m MainWindow) New(cfg *AppConfig) MainWindow {
	a := *cfg.App

	// Generate new master window, set size, center it on screen
	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))
	w.SetMaster() // master window, when closed closes all other windows
	w.CenterOnScreen()

	// Generate window content UI
	content := mainWindowUI(cfg)
	w.SetContent(content)

	// Create the main menu
	menu, setCanSave := FileMenu(cfg)
	w.SetMainMenu(menu)

	return MainWindow{
		Window:     w,
		menu:       menu,
		SetCanSave: setCanSave,
	}
}

// mainWindowUI creates the main window UI content
func mainWindowUI(cfg *AppConfig) *fyne.Container {

	// Title
	title := TitleText("Welcome to the Gist editor!")
	subLabel := widget.NewLabel("Click View Gists to see all your gists, or create a new one with New Gist.")

	// Title and welcome text container
	titleContainer := container.NewVBox(title, subLabel)

	// Buttons for "View Gists" and "New Gist"
	viewGistsButton := widget.NewButton("View Gists", cfg.ShowListWindow)
	newGistButton := widget.NewButton("New Gist", cfg.NewFile)
	closeBtn := widget.NewButton("Exit", cfg.Exit)

	// Centered buttons grid
	buttons := container.NewGridWithColumns(3, newGistButton, viewGistsButton, closeBtn)

	spacer := layout.NewSpacer()

	// Vertical grid layout
	content := container.New(layout.NewGridLayoutWithRows(5), titleContainer, spacer, spacer, spacer, buttons)
	return content
}
