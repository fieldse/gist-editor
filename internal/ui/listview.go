package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

// Returns a list view widget of all user gists
func ListWidget() *fyne.Container {
	title := widget.NewLabel("Your gists")
	title.TextStyle.Bold = true

	// Test buttons
	testButton := widget.NewButton("test", func() {
		logger.Info("button pressed")
	})
	okButton := widget.NewButton("ok", func() {
		logger.Info("OK button pressed")
	})

	// Example content widget
	buttons := container.NewHBox(testButton, okButton)
	vBox := container.NewVBox(title, buttons)

	content := container.NewHBox(title, vBox)
	return content
}

// Return a List view window
func ListWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Your Gists")
	w.Resize(fyne.NewSize(600, 400))

	content := ListWidget()
	w.SetContent(content)

	return w
}
