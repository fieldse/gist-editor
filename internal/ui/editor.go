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
	Title               string
	Content             string
	IsDirty             bool
	editor              *widget.Entry   // the text editor field
	editWindow          fyne.Window     // the editor window
	previewPane         *fyne.Container // the editor as split-pane view with preview
	singlePane          *fyne.Container // the editor as single-pane view
	togglePreviewButton *widget.Button
	IsVisible           bool
	PreviewVisible      bool // is the preview pane visible, or are we in single pane view
}

// Show displays the editor window
func (e *Editor) Show() {
	e.IsVisible = true
	e.editWindow.Show()
}

// Hide hides the editor window
func (e *Editor) Hide() {
	e.IsVisible = false
	e.editWindow.Hide()
}

// TogglePreview toggles visibility of the markdown preview pane
func (e *Editor) TogglePreview() {
	if !e.PreviewVisible { // Hide single-pane view, show split-pane
		e.singlePane.Hide()
		e.previewPane.Show()
		e.togglePreviewButton.SetText("Hide Preview")
		e.PreviewVisible = true
	} else {
		e.singlePane.Show() // Hide split-pane view, show single-pane
		e.previewPane.Hide()
		e.togglePreviewButton.SetText("Show Preview")
		e.PreviewVisible = false
	}
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

// New creates a new Editor window and text editor widget
func (e Editor) New(cfg *AppConfig) *Editor {
	a := *cfg.App
	w := a.NewWindow("Edit Gist")
	f := cfg.CurrentFile
	w.Resize(fyne.NewSize(800, 600))

	content, editor, previewPane, singlePane, togglePreviewButton := editUI(cfg, f.Gist, w)
	w.SetContent(content)
	w.CenterOnScreen()

	return &Editor{
		editor:              editor,
		editWindow:          w,
		previewPane:         previewPane,
		singlePane:          singlePane,
		togglePreviewButton: togglePreviewButton,
	}
}

// Generates the UI for the edit window
// Returns the container, and a pointer to the content editor, preview pane, and single-pane containers,
// and the toggle preview button
func editUI(cfg *AppConfig, g *github.Gist, w fyne.Window) (*fyne.Container, *widget.Entry, *fyne.Container, *fyne.Container, *widget.Button) {

	// Title
	titleBox := TitleBox(g.Filename)

	// Text editor toolbar
	textEditorToolbar := MarkdownToolbarUI(cfg)

	// Editor input
	editor := widget.NewMultiLineEntry()
	editor.SetText(g.Content)

	// Top section -- edit toolbar & title
	topBox := container.NewVBox(widget.NewLabel("Edit"), textEditorToolbar)
	editPane := container.NewBorder(topBox, nil, nil, nil, editor)

	// Preview pane
	preview := widget.NewRichTextFromMarkdown(editor.Text)
	previewPane := container.NewBorder(widget.NewLabel("Preview"), nil, nil, nil, preview)
	editor.OnChanged = preview.ParseMarkdown // parse markdown to rich text on changed

	// Split pane view
	splitView := container.NewHSplit(editPane, previewPane)

	// Single pane (no preview)
	singlePane := container.NewCenter(editPane)

	// Buttons
	spacer := layout.NewSpacer()
	togglePreviewButton := widget.NewButton("Show preview", func() {
		cfg.Editor.TogglePreview()
	})
	saveButton := widget.NewButton("Save", func() {
		cfg.SaveFile()
	})
	closeButton := widget.NewButton("Close", func() {
		cfg.CloseFile()
		cfg.Editor.Hide()
	})
	buttons := ButtonContainer(4, spacer, togglePreviewButton, saveButton, closeButton)

	// Wrapper container
	content := container.NewBorder(titleBox, buttons, nil, nil, splitView)
	return content, editor, previewPane, singlePane, togglePreviewButton
}
