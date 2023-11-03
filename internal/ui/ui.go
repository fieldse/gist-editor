package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func viewGists() {
	// TODO
}

func newGist() {
	// TODO
}

func Startui() {
	a := app.New()
	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))

	// Title
	title := widget.NewLabel("Welcome to the Gist editor!")
	title.TextStyle.Bold = true
	subLabel := widget.NewLabel("Click View Gists or New Gist below.")

	// Title and welcome text container
	titleContainer := container.NewVBox(title, subLabel)

	// Buttons for "View Gists" and "New Gist"
	b1 := widget.NewButton("View Gists", viewGists)
	b2 := widget.NewButton("New Gist", newGist)

	// Centered buttons grid
	buttons := container.NewGridWithColumns(2, b1, b2)

	spacer := layout.NewSpacer()

	// Vertical grid layout
	content := container.New(layout.NewGridLayoutWithRows(3), titleContainer, spacer, buttons)

	w.SetContent(content)
	w.ShowAndRun()

}
