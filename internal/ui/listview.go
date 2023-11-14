package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// Returns a list view widget of all user gists
func ListWidget(hide func()) *fyne.Container {
	spacer := layout.NewSpacer()
	okButton := widget.NewButton("Ok", hide)

	// Example content widget
	var data []string
	for _, x := range github.MockGistData {
		data = append(data, x.Filename)
	}
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

	titleContainer := TitleBox("Your Gists")
	listContainer := container.NewStack(list)

	buttons := ButtonContainer(3, spacer, spacer, okButton)

	// Total content includes title, list section, and buttons
	content := container.NewBorder(titleContainer, buttons, nil, nil, listContainer)
	return content
}

// ListWindow returns a List view window
func ListWindow(cfg *AppConfig) fyne.Window {
	a := *cfg.App
	w := a.NewWindow("Your Gists")
	w.Resize(fyne.NewSize(800, 600))

	content := ListWidget(w.Hide)
	w.SetContent(content)
	w.CenterOnScreen()

	return w
}
