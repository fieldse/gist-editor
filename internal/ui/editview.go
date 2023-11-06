// Edit gist view
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// Placeholder data for Gist content
var data = github.ExampleGist

func EditWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Edit Gist")
	w.Resize(fyne.NewSize(500, 600))

	content := EditUI(data, w.Hide)
	w.SetContent(content)

	return w
}

// TODO -- placeholder for Save func
func saveGist() {}

// Generates the UI for the edit window
func EditUI(gist github.Gist, hide func()) *fyne.Container {

	spacer := layout.NewSpacer()
	saveButton := widget.NewButton("Save", saveGist)
	closeButton := widget.NewButton("Close", hide)

	// Title
	titleBox := TitleBox(gist.Filename)

	// Content
	contentBox := widget.NewMultiLineEntry()
	contentBox.SetText(gist.Content)

	// Buttons
	buttons := ButtonContainer(3, spacer, saveButton, closeButton)

	// Wrapper container
	content := container.NewBorder(titleBox, buttons, nil, nil, contentBox)
	return content
}
