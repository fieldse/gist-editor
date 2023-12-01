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

// TextSelection represents the position and content of the editor's current text selection
type TextSelection struct {
	Col     int
	Row     int
	Content string
}

// GetSelection returns the current text selection and position
func (e *Editor) GetSelection() TextSelection {
	return TextSelection{
		Col:     e.editor.CursorColumn,
		Row:     e.editor.CursorRow,
		Content: e.editor.SelectedText(),
	}
}

// ReplaceSelection returns the current text selection and cursor position
func (e *Editor) ReplaceSelection(orig TextSelection, newText TextSelection) {
	// TODO: function to replace text selection by row & column
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

// PreviewEditContainer is the wrapper for the Preview and Edit panes.
// It shows the editor on the left, with preview on the right in split view.
// The toggle method hides the preview pane (moves to collapsed state).
type PreviewEditContainer struct {
	Content      *container.Split // split view of the preview and edit panes
	ToggleButton *widget.Button   // the toggle Preview button
	previewPane  *fyne.Container
}

// New returns a new PreviewEditContainer with toggle functionality
func (p PreviewEditContainer) New(
	previewPane *fyne.Container,
	editPane *fyne.Container,
) *PreviewEditContainer {

	wrapper := container.NewHSplit(editPane, previewPane)

	// The offset is the ratio between right and left panes, from 0-1.
	// 	0.0 means right pane occupies 100% space, left pane collapsed.
	// 	1.0 means left pane occupies 100% space, right pane collapsed.
	wrapper.SetOffset(1.0)
	pp := &PreviewEditContainer{
		Content:     wrapper,
		previewPane: previewPane,
	}
	previewPane.Hide()
	pp.ToggleButton = widget.NewButton("Show preview", func() {
		pp.TogglePreview()
	})

	return pp
}

// PreviewIsVisible returns whether the markdown preview pane is visible,
// determined by the offset attribute of the Split element.
// If the user has adjusted it, it may be visible.
func (p *PreviewEditContainer) PreviewIsVisible() bool {
	return p.previewPane.Visible() && p.Content.Offset < 0.9 // Assume a sane default of >10% visibility means it's visible.
}

// TogglePreview toggles the visiblility of the markdown preview pane.
func (p *PreviewEditContainer) TogglePreview() {
	// The visiblility of the markdown preview pane is determined by the offset
	// attribute of the Split element.
	// If the user has adjusted it, it may be visible.
	if p.PreviewIsVisible() {
		p.previewPane.Hide()
		p.Content.SetOffset(1.0) // hide the preview
		p.ToggleButton.SetText("Show preview")
	} else {
		p.previewPane.Show()
		p.Content.SetOffset(0.5) // show the preview and editor at 50/50 ratio
		p.ToggleButton.SetText("Hide preview")
	}
}
