// Contains markdown editor toolbar component
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MarkdownToolbar represents a toolbar for the markdown text editor
type MarkdownToolbar struct {
	Toolbar fyne.Widget
}

func (m MarkdownToolbar) New() *MarkdownToolbar {
	return &MarkdownToolbar{}
}

// toolbarIcons is the set of toolbar item icons
type toolbarIcons struct {
	H1Icon              fyne.Resource
	H2Icon              fyne.Resource
	H3Icon              fyne.Resource
	BoldIcon            fyne.Resource
	ItalicIcon          fyne.Resource
	LinkIcon            fyne.Resource
	ImageIcon           fyne.Resource
	QuoteBlockIcon      fyne.Resource
	CodeBlockIcon       fyne.Resource
	InlineCodeBlockIcon fyne.Resource
	PageBreakIcon       fyne.Resource
	UndoIcon            fyne.Resource
	RedoIcon            fyne.Resource
}

// toolbarActions is the action functions for the toolbar items
type toolbarActions struct {
	H1Action              func()
	H2Action              func()
	H3Action              func()
	BoldAction            func()
	ItalicAction          func()
	LinkAction            func()
	ImageAction           func()
	QuoteBlockAction      func()
	CodeBlockAction       func()
	InlineCodeBlockAction func()
	PageBreakAction       func()
	UndoAction            func()
	RedoAction            func()
}

// newToolbarActions returns a new toolbar actions struct
func newToolbarActions() toolbarActions {
	return toolbarActions{} // FIXME -- add functions
}

// newToolbarIcons returns a new set of toolbar icons
func newToolbarIcons() toolbarIcons {
	return toolbarIcons{} // FIXME -- add toolbar icons
}

// MarkdownToolbar returns toolbar component for the markdown text editor
func MarkdownToolbarUI(a *AppConfig) *widget.Toolbar {
	icons := newToolbarIcons()
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
		widget.NewToolbarAction(icons.InlineCodeBlockIcon, actions.InlineCodeBlockAction),
		widget.NewToolbarAction(icons.PageBreakIcon, actions.PageBreakAction),
		widget.NewToolbarAction(icons.UndoIcon, actions.UndoAction),
		widget.NewToolbarAction(icons.RedoIcon, actions.RedoAction),
	)
}
