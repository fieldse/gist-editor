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

func EditWindow(cfg *AppConfig) *fyne.Window {
	a := *cfg.App
	w := a.NewWindow("Edit Gist")
	w.Resize(fyne.NewSize(800, 600))

	content := EditUI(cfg, data, w)
	w.SetContent(content)
	w.CenterOnScreen()

	return &w
}

// Generates the UI for the edit window
func EditUI(cfg *AppConfig, gist github.Gist, w fyne.Window) *fyne.Container {

	spacer := layout.NewSpacer()
	saveButton := widget.NewButton("Save", func() {
		cfg.SaveFile()
	})
	closeButton := widget.NewButton("Close", func() {
		w.Hide()
	})

	// Title
	titleBox := TitleBox(gist.Filename)

	// Editor input
	edit := widget.NewMultiLineEntry()
	edit.SetText(gist.Content)
	editPane := container.NewBorder(widget.NewLabel("Edit"), nil, nil, nil, edit)

	// Preview pane
	preview := widget.NewRichTextFromMarkdown(edit.Text)
	previewPane := container.NewBorder(widget.NewLabel("Preview"), nil, nil, nil, preview)
	edit.OnChanged = preview.ParseMarkdown // parse markdown to rich text on changed

	// Split pane view
	splitView := container.NewHSplit(editPane, previewPane)

	// Buttons
	buttons := ButtonContainer(3, spacer, saveButton, closeButton)

	// Wrapper container
	content := container.NewBorder(titleBox, buttons, nil, nil, splitView)
	return content
}
