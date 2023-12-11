// Contains markdown editor toolbar component
package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/icons"
	"github.com/fieldse/gist-editor/internal/logger"
)

// MarkdownToolbar is a toolbar for the markdown text editor
type MarkdownToolbar struct {
	Toolbar fyne.Widget
}

func (m MarkdownToolbar) New() *MarkdownToolbar {
	return &MarkdownToolbar{}
}

// MarkdownToolbar returns toolbar component for the markdown text editor
// Will throw error and exit program on failure
func MarkdownToolbarUI(e *MultiLineWidget) *widget.Toolbar {
	icons, err := icons.ToolbarIcons{}.Load()
	if err != nil {
		logger.Fatal("load resources failed", err)
	}
	// Text editing functions to pass to the toolbar
	actions := NewTextFunctions(e)

	// Menu items
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.BrokenImageIcon(), func() { e.DebugTextSelection() }),
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
		widget.NewToolbarAction(icons.UndoIcon, actions.Undo),
		widget.NewToolbarAction(icons.RedoIcon, actions.Redo),
	)
}
