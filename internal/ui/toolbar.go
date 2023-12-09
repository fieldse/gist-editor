// Contains markdown editor toolbar component
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	editorfunctions "github.com/fieldse/gist-editor/internal/editor"
	"github.com/fieldse/gist-editor/internal/logger"
	"github.com/fieldse/gist-editor/internal/shared"
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
func MarkdownToolbarUI(a *AppConfig) *widget.Toolbar {
	icons, err := ToolbarIcons{}.Load()
	if err != nil {
		logger.Fatal("load resources failed", err)
	}

	// Edit text functions we have to pass to the toolbar
	getText := func() string {
		return a.Editor.Content()
	}
	getSelection := func() shared.TextSelection {
		return a.Editor.GetSelection()
	}
	replaceText := func(s string) {
		a.Editor.SetContent(s)
	}
	undoFunc := func() {
		a.Editor.Undo()
	}
	redoFunc := func() {
		a.Editor.Redo()
	}

	actions := editorfunctions.EditorFunctions{
		GetText:      getText,
		GetSelection: getSelection,
		ReplaceText:  replaceText,
	}

	// Menu items
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.BrokenImageIcon(), actions.DebugTextSelection),
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
