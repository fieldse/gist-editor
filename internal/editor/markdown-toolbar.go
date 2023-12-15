// Contains markdown editor toolbar component
package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/fieldse/gist-editor/internal/icons"
)

var Icons = icons.ToolbarIcons

// MarkdownToolbar is a toolbar for the markdown text editor
type MarkdownToolbar struct {
	Toolbar fyne.Widget
}

// New returns a new toolbar component for the markdown text editor,
// with attached functions and icons
func New(e *MultiLineWidget) *widget.Toolbar {
	// Text editing functions to pass to the toolbar
	actions := newToolbarActions(e)

	// Menu items
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.BrokenImageIcon(), func() { e.DebugTextSelection() }),
		widget.NewToolbarAction(Icons.H1Icon, actions.H1),
		widget.NewToolbarAction(Icons.H2Icon, actions.H2),
		widget.NewToolbarAction(Icons.H3Icon, actions.H3),
		widget.NewToolbarAction(Icons.BoldIcon, actions.Bold),
		widget.NewToolbarAction(Icons.ItalicIcon, actions.Italic),
		widget.NewToolbarAction(Icons.UnderlineIcon, actions.Underline),
		widget.NewToolbarAction(Icons.LinkIcon, actions.Link),
		widget.NewToolbarAction(Icons.ImageIcon, actions.Image),
		widget.NewToolbarAction(Icons.QuoteBlockIcon, actions.QuoteBlock),
		widget.NewToolbarAction(Icons.CodeBlockIcon, actions.CodeBlock),
		widget.NewToolbarAction(Icons.PageBreakIcon, actions.PageBreak),
		widget.NewToolbarAction(Icons.UndoIcon, actions.Undo),
		widget.NewToolbarAction(Icons.RedoIcon, actions.Redo),
		widget.NewToolbarAction(Icons.EraserIcon, actions.ClearFormatting),
	)
}
