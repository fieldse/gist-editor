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
	Title                string
	Content              string
	IsDirty              bool
	editor               *widget.Entry         // the text editor field
	editWindow           fyne.Window           // the editor window
	previewEditContainer *PreviewEditContainer // a wrapper, containing the preview and edit widgets
	IsVisible            bool
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

	content, editor, previewEditContainer := editUI(cfg, f.Gist, w)
	w.SetContent(content)
	w.CenterOnScreen()

	return &Editor{
		editor:               editor,
		editWindow:           w,
		previewEditContainer: previewEditContainer,
	}
}

// Generates the UI for the edit window
// Returns the container, and a pointer to the content editor, and a wrapper for the single-pane and split-pane containers,
func editUI(cfg *AppConfig, g *github.Gist, w fyne.Window) (*fyne.Container, *widget.Entry, *PreviewEditContainer) {

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

	// Preview and edit pane wrapper
	previewEditContainer := PreviewEditContainer{}.New(previewPane, editPane)

	// Buttons
	spacer := layout.NewSpacer()
	saveButton := widget.NewButton("Save", func() {
		cfg.SaveFile()
	})
	closeButton := widget.NewButton("Close", func() {
		cfg.CloseFile()
		cfg.Editor.Hide()
	})
	buttons := ButtonContainer(4, spacer, previewEditContainer.ToggleButton, saveButton, closeButton)

	// Wrapper container
	content := container.NewBorder(titleBox, buttons, nil, nil, previewEditContainer.Content)
	return content, editor, previewEditContainer
}

// PreviewEditContainer is the container wrapper for the Preview and Edit
// panes. It alternately displays either the editor or editor with preview
// in split view
type PreviewEditContainer struct {
	Content        *fyne.Container // contains both the preview and edit view
	ToggleButton   *widget.Button  // the toggle Preview button
	previewPane    *fyne.Container
	editPane       *fyne.Container
	PreviewVisible bool
}

// New returns a new PreviewEditContainer with toggle functionality
func (p PreviewEditContainer) New(
	previewPane *fyne.Container,
	editPane *fyne.Container,
) *PreviewEditContainer {
	pp := &PreviewEditContainer{
		previewPane: previewPane,
		editPane:    editPane,
		Content:     container.NewStack(previewPane, editPane),
	}
	pp.ToggleButton = widget.NewButton("Show preview", func() {
		pp.TogglePreview()
	})
	// We don't show the preview pane by default
	previewPane.Hide()

	return pp
}

// TogglePreview toggles the visiblility of the markdown preview
func (p *PreviewEditContainer) TogglePreview() {
	if p.PreviewVisible {
		// do something
		// hide the preview
		p.previewPane.Hide()
		p.editPane.Show()
		p.ToggleButton.SetText("Show preview")
	} else {
		// do something else
		// show the preview
		p.previewPane.Show()
		p.editPane.Hide()
		p.ToggleButton.SetText("Hide preview")
	}
	p.PreviewVisible = !p.PreviewVisible
}
