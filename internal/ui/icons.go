// Contains icons for the Markdown editor toolbar buttons
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var ICON_ASSETS = make(map[string][]byte)
var exampleIcon = theme.ContentAddIcon()

// ToolbarIcons is the set of icons for the Markdown editor toolbar
type ToolbarIcons struct {
	H1Icon         fyne.Resource
	H2Icon         fyne.Resource
	H3Icon         fyne.Resource
	BoldIcon       fyne.Resource
	ItalicIcon     fyne.Resource
	LinkIcon       fyne.Resource
	ImageIcon      fyne.Resource
	QuoteBlockIcon fyne.Resource
	CodeBlockIcon  fyne.Resource
	PageBreakIcon  fyne.Resource
	UndoIcon       fyne.Resource
	RedoIcon       fyne.Resource
}

// Load icon resources and returns a populated ToolbarIcons instance
func (t ToolbarIcons) Load() (ToolbarIcons, error) {
	// FIXME: load icon data instead of placeholders
	return ToolbarIcons{
		H1Icon:         exampleIcon,
		H2Icon:         exampleIcon,
		H3Icon:         exampleIcon,
		BoldIcon:       exampleIcon,
		ItalicIcon:     exampleIcon,
		LinkIcon:       exampleIcon,
		ImageIcon:      exampleIcon,
		QuoteBlockIcon: exampleIcon,
		CodeBlockIcon:  exampleIcon,
		PageBreakIcon:  exampleIcon,
		UndoIcon:       theme.ContentUndoIcon(),
		RedoIcon:       theme.ContentRedoIcon(),
	}, nil
}
