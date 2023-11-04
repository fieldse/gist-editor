package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Returns a list view widget of all user gists
func ListWidget(hide func()) *fyne.Container {
	title := TitleText("Your gists")

	// Test buttons
	spacer := layout.NewSpacer()
	okButton := widget.NewButton("ok", hide)

	// Example content widget
	var data = []string{"a", "string", "list"}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	buttons := container.NewGridWithColumns(3, spacer, spacer, okButton)
	content := container.NewStack(title, list, buttons)

	return content
}

// Return a List view window
func ListWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Your Gists")
	w.Resize(fyne.NewSize(500, 300))

	content := ListWidget(w.Hide)
	w.SetContent(content)

	return w
}
