package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Returns a list view widget of all user gists
func ListWidget(hide func()) *fyne.Container {
	title := widget.NewLabel("Your gists")
	title.TextStyle.Bold = true

	// Test buttons
	spacer := layout.NewSpacer()
	okButton := widget.NewButton("ok", hide)

	// Example content widget
	buttons := container.NewHBox(spacer, okButton)
	vBox := container.NewVBox(title, spacer, buttons)

	content := container.NewHBox(vBox)
	return content
}

// Return a List view window
func ListWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Your Gists")
	w.Resize(fyne.NewSize(600, 400))

	content := ListWidget(w.Hide)
	w.SetContent(content)

	return w
}
