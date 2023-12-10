// Contains markdown editor toolbar component
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	editorfunctions "github.com/fieldse/gist-editor/internal/editor"
	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/fieldse/gist-editor/internal/shared"
	"github.com/fieldse/gist-editor/internal/widgets"
)

// MarkdownToolbar represents a toolbar for the markdown text editor
type MarkdownToolbar struct {
	Toolbar fyne.Widget
}

func (m MarkdownToolbar) New() *MarkdownToolbar {
	return &MarkdownToolbar{}
}

// MarkdownToolbar returns toolbar component for the markdown text editor
// Will throw error and exit program on failure
func MarkdownToolbarUI(e *widgets.MultiLineWidget) *widget.Toolbar {
	icons, err := ToolbarIcons{}.Load()
	if err != nil {
		logger.Fatal("load resources failed", err)
	}
	// Text editing functions to pass to the toolbar
	getText := func() string {
		return e.Text
	}
	setText := func(s string) {
		e.SetText(s)
	}
	getSelection := func() shared.TextSelection {
		return e.GetSelection()
	}
	selectionStart := func() shared.Position {
		return e.SelectionStart()
	}
	undoFunc := func() {
		e.Undo()
	}
	redoFunc := func() {
		e.Redo()
	}
	debugFunc := func() {
		e.DebugTextSelection()
	}
	actions := editorfunctions.NewEditorFunctions(getText, setText, getSelection, selectionStart)

	// Menu items
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.BrokenImageIcon(), debugFunc),
		widget.NewToolbarAction(icons.H1Icon, actions.H1),
		widget.NewToolbarAction(icons.H2Icon, actions.H2),
		widget.NewToolbarAction(icons.H3Icon, actions.H3),
		widget.NewToolbarAction(icons.BoldIcon, actions.Bold),
		widget.NewToolbarAction(icons.ItalicIcon, actions.Italic),
		widget.NewToolbarAction(icons.LinkIcon, actions.Link),
		widget.NewToolbarAction(icons.ImageIcon, actions.Image),
		widget.NewToolbarAction(icons.QuoteBlockIcon, actions.QuoteBlock),
		widget.NewToolbarAction(icons.CodeBlockIcon, actions.CodeBlock),
		widget.NewToolbarAction(icons.PageBreakIcon, actions.PageBreak),
		widget.NewToolbarAction(icons.UndoIcon, undoFunc),
		widget.NewToolbarAction(icons.RedoIcon, redoFunc),
	)
}
