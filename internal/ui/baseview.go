// Base window for the app
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Intitialize and return the base window
func BaseWindow(cfg *AppConfig) fyne.Window {
	a := *cfg.App

	// Generate new master window, set size, center it on screen
	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))
	w.SetMaster() // master window, when closed closes all other windows
	w.CenterOnScreen()

	// Generate window content UI
	content := BaseView(cfg)
	w.SetContent(content)

	return w
}

// Base view upon opening the app
func BaseView(cfg *AppConfig) *fyne.Container {

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
