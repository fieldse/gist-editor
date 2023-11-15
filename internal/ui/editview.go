// Edit gist view
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// Editor represents the Gist editor window, and provides methods
// to update the title & content of the editor widget
type Editor struct {
	Title      string
	Content    string
	IsDirty    bool
	editor     *widget.Entry // the text editor field
	editWindow fyne.Window   // the editor window
}

// Show displays the editor window
func (e *Editor) Show() {
	e.editWindow.Show()
}

// Hide hides the editor window
func (e *Editor) Hide() {
	e.editWindow.Hide()
}

// SetContent sets the contents of the text editor field
func (e *Editor) SetContent(text string) {
	e.editor.SetText(text)
}

// Clear resets the title and contents of the text editor
func (e *Editor) Clear() {
	e.Title = "Edit"
	e.editor.SetText("")
}

// Create creates a new Editor window and text editor widget
func (e Editor) Create(cfg *AppConfig) *Editor {
	a := *cfg.App
	w := a.NewWindow("Edit Gist")
	f := cfg.CurrentFile
	w.Resize(fyne.NewSize(800, 600))

	content, editor := EditUI(cfg, f.Gist, w)
	w.SetContent(content)
	w.CenterOnScreen()

	return &Editor{
		editor:     editor,
		editWindow: w,
	}
}

// Generates the UI for the edit window
// Returns the container, and a pointer to the content editor
func EditUI(cfg *AppConfig, g *github.Gist, w fyne.Window) (*fyne.Container, *widget.Entry) {
	spacer := layout.NewSpacer()
	saveButton := widget.NewButton("Save", func() {
		cfg.SaveFile()
	})
	closeButton := widget.NewButton("Close", func() {
		cfg.CloseFile()
		cfg.Editor.Hide()
	})

	// Title
	titleBox := TitleBox(g.Filename)

	// Editor input
	editor := widget.NewMultiLineEntry()
	editor.SetText(g.Content)
	editPane := container.NewBorder(widget.NewLabel("Edit"), nil, nil, nil, editor)

	// Preview pane
	preview := widget.NewRichTextFromMarkdown(editor.Text)
	previewPane := container.NewBorder(widget.NewLabel("Preview"), nil, nil, nil, preview)
	editor.OnChanged = preview.ParseMarkdown // parse markdown to rich text on changed

	// Split pane view
	splitView := container.NewHSplit(editPane, previewPane)

	// Buttons
	buttons := ButtonContainer(3, spacer, saveButton, closeButton)

	// Wrapper container
	content := container.NewBorder(titleBox, buttons, nil, nil, splitView)
	return content, editor
}
