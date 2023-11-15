package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// ListView is the user's Gist list view window
type ListView struct {
	listWindow fyne.Window
	gists      []GistFile
}

// Show shows the list view window
func (l *ListView) Show() {
	l.listWindow.Show()
}

// SetGists populates the list view data
func (l *ListView) SetGists(data []GistFile) {
	l.gists = data
}

// Returns a list view widget of all user gists
func listWidget(hide func()) *fyne.Container {
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

// New creates a ListView instance with list view window
func (l ListView) New(cfg *AppConfig) *ListView {
	a := *cfg.App
	w := a.NewWindow("Your Gists")
	w.Resize(fyne.NewSize(800, 600))

	content := listWidget(w.Hide)
	w.SetContent(content)
	w.CenterOnScreen()

	return &ListView{
		listWindow: w,
	}
}
