package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

func newGist() {
	// TODO
}

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	BaseWindow *fyne.Window
	ListWindow *fyne.Window
}

// Base view upon opening the app
func BaseView(showList func()) *fyne.Container {

	// Title
	title := widget.NewLabel("Welcome to the Gist editor!")
	title.TextStyle.Bold = true
	subLabel := widget.NewLabel("Click View Gists to see all your gists, or create a new one with New Gist.")

	// Title and welcome text container
	titleContainer := container.NewVBox(title, subLabel)

	// Buttons for "View Gists" and "New Gist"
	b1 := widget.NewButton("View Gists", showList)
	b2 := widget.NewButton("New Gist", newGist)

	// Centered buttons grid
	buttons := container.NewGridWithColumns(2, b1, b2)

	spacer := layout.NewSpacer()

	// Vertical grid layout
	content := container.New(
		layout.NewGridLayoutWithRows(5), titleContainer, spacer, spacer, spacer, buttons,
	)
	return content
}

func StartUI() {
	a := app.New()
	// var cfg AppConfig

	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))

	// Gists list window
	l := a.NewWindow("Your Gists")
	logger.Debug("list window: %+v\n", l)
	logger.Debug("app: %+v\n", l)
	listContent := ListWidget()

	logger.Debug("listContent: %+v\n", listContent)
	l.SetContent(listContent)

	// Base view window
	content := BaseView(func() { l.Show() })

	w.SetContent(content)
	w.ShowAndRun()

}
