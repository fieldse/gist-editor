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
	window fyne.Window
	list   *widget.List
	gists  []github.Gist
}

// Show shows the list view window
func (l *ListView) Show() {
	l.window.Show()
}

// Hide Hides the list view window
func (l *ListView) Hide() {
	l.window.Hide()
}

// SetGists populates the list view data
func (l *ListView) SetGists(data []github.Gist) {
	l.gists = data
	l.list = newList(data)
}

// Clear clears the list view data
func (l *ListView) Clear() {
	l.gists = []github.Gist{}
	l.list = newList(l.gists)
}

// newList returns a new Fyne list widget with the given Gist data
func newList(gists []github.Gist) *widget.List {
	var data []string
	for _, x := range gists {
		data = append(data, x.Filename)
	}
	l := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	return l
}

// Returns a list view widget of all user gists
func listWidget(hide func(), gists []github.Gist) *fyne.Container {
	spacer := layout.NewSpacer()
	okButton := widget.NewButton("Ok", hide)

	// List data view
	l := newList(gists)
	titleContainer := TitleBox("Your Gists")
	listContainer := container.NewStack(l)

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

	content := listWidget(w.Hide, []github.Gist{})
	w.SetContent(content)
	w.CenterOnScreen()

	return &ListView{
		window: w,
	}
}
