// Contains icons for the Markdown editor toolbar buttons
package ui

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
)

// ToolbarIcons is the set of icons for the Markdown editor toolbar
type ToolbarIcons struct {
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

// loadIcon loads the data for a single icon file
// Raises exception and exits on failure
func loadIcon(filename string) ([]byte, error) {
	b := make([]byte, 400) // FIXME: load icon data from /static
	err := errors.New("not implemented -- loadIcon from file")
	if err != nil {
		return b, fmt.Errorf("failed to load icon data for %s - %w", filename, err)
	}
	return b, nil
}

// loadAllIcons loads and returns icon resources for all the Markdown editor buttons
func loadAllIcons() (ToolbarIcons, error) {

	// TODO: there has to be a better way to do this without the repeated error checks
	h1Icon, err := loadIcon("h1-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h2Icon, err := loadIcon("h2-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h3Icon, err := loadIcon("h3-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	boldIcon, err := loadIcon("bold-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	italicIcon, err := loadIcon("italic-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	linkIcon, err := loadIcon("link-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	imageIcon, err := loadIcon("image-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	quoteBlockIcon, err := loadIcon("quote-block-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	codeBlockIcon, err := loadIcon("code-block-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	inlineCodeBlockIcon, err := loadIcon("inline-code-block-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	pageBreakIcon, err := loadIcon("page-break-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	undoIcon, err := loadIcon("undo-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	redoIcon, err := loadIcon("redo-icon.png")
	if err != nil {
		return ToolbarIcons{}, err
	}

	return ToolbarIcons{
		H1Icon:              fyne.NewStaticResource("H1Icon", h1Icon),
		H2Icon:              fyne.NewStaticResource("H2Icon", h2Icon),
		H3Icon:              fyne.NewStaticResource("H3Icon", h3Icon),
		BoldIcon:            fyne.NewStaticResource("BoldIcon", boldIcon),
		ItalicIcon:          fyne.NewStaticResource("ItalicIcon", italicIcon),
		LinkIcon:            fyne.NewStaticResource("LinkIcon", linkIcon),
		ImageIcon:           fyne.NewStaticResource("ImageIcon", imageIcon),
		QuoteBlockIcon:      fyne.NewStaticResource("QuoteBlockIcon", quoteBlockIcon),
		CodeBlockIcon:       fyne.NewStaticResource("CodeBlockIcon", codeBlockIcon),
		InlineCodeBlockIcon: fyne.NewStaticResource("InlineCodeBlockIcon", inlineCodeBlockIcon),
		PageBreakIcon:       fyne.NewStaticResource("PageBreakIcon", pageBreakIcon),
		UndoIcon:            fyne.NewStaticResource("UndoIcon", undoIcon),
		RedoIcon:            fyne.NewStaticResource("RedoIcon", redoIcon),
	}, nil
}
