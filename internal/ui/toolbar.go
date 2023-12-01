// Contains markdown editor toolbar component
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

// MarkdownToolbar represents a toolbar for the markdown text editor
type MarkdownToolbar struct {
	Toolbar fyne.Widget
}

func (m MarkdownToolbar) New() *MarkdownToolbar {
	return &MarkdownToolbar{}
}

// toolbarActions is the action functions for the toolbar items
type toolbarActions struct {
	H1Action         func()
	H2Action         func()
	H3Action         func()
	BoldAction       func()
	ItalicAction     func()
	LinkAction       func()
	ImageAction      func()
	QuoteBlockAction func()
	CodeBlockAction  func()
	PageBreakAction  func()
	UndoAction       func()
	RedoAction       func()
}

// FIXME -- add functions
func placeholderAction(actionName string) func() {
	return func() {
		logger.Debug("placeholder action called for %s", actionName)
	}
}

// newToolbarActions returns a new toolbar actions struct
func newToolbarActions() toolbarActions {
	return toolbarActions{
		H1Action:         placeholderAction("H1Action"),
		H2Action:         placeholderAction("H2Action"),
		H3Action:         placeholderAction("H3Action"),
		BoldAction:       placeholderAction("BoldAction"),
		ItalicAction:     placeholderAction("ItalicAction"),
		LinkAction:       placeholderAction("LinkAction"),
		ImageAction:      placeholderAction("ImageAction"),
		QuoteBlockAction: placeholderAction("QuoteBlockAction"),
		CodeBlockAction:  placeholderAction("CodeBlockAction"),
		PageBreakAction:  placeholderAction("PageBreakAction"),
		UndoAction:       placeholderAction("UndoAction"),
		RedoAction:       placeholderAction("RedoAction"),
	}
}

// MarkdownToolbar returns toolbar component for the markdown text editor
// Will throw error and exit program on failure
func MarkdownToolbarUI(a *AppConfig) *widget.Toolbar {
	icons, err := ToolbarIcons{}.Load()
	if err != nil {
		logger.Fatal("load resources failed", err)
	}
	actions := newToolbarActions()

	// Menu items
	return widget.NewToolbar(
		widget.NewToolbarAction(icons.H1Icon, actions.H1Action),
		widget.NewToolbarAction(icons.H2Icon, actions.H2Action),
		widget.NewToolbarAction(icons.H3Icon, actions.H3Action),
		widget.NewToolbarAction(icons.BoldIcon, actions.BoldAction),
		widget.NewToolbarAction(icons.ItalicIcon, actions.ItalicAction),
		widget.NewToolbarAction(icons.LinkIcon, actions.LinkAction),
		widget.NewToolbarAction(icons.ImageIcon, actions.ImageAction),
		widget.NewToolbarAction(icons.QuoteBlockIcon, actions.QuoteBlockAction),
		widget.NewToolbarAction(icons.CodeBlockIcon, actions.CodeBlockAction),
		widget.NewToolbarAction(icons.PageBreakIcon, actions.PageBreakAction),
		widget.NewToolbarAction(icons.UndoIcon, actions.UndoAction),
		widget.NewToolbarAction(icons.RedoIcon, actions.RedoAction),
	)
}
