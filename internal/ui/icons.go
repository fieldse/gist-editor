// Contains icons for the Markdown editor toolbar buttons
package ui

import (
	"fmt"
	"os"
	"path"

	"fyne.io/fyne/v2"
)

var ICON_DIR = path.Join("..", "static", "icons")

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

// loadIcon loads the data for a single icon file
// Raises exception and exits on failure
func loadIcon(filename string) ([]byte, error) {
	var b []byte
	f, err := os.Open(path.Join(ICON_DIR, filename))
	if err != nil {
		return b, fmt.Errorf("failed to load icon file for %s - %w", filename, err)
	}
	_, err = f.Read(b)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read icon data for %s", filename)
	}
	return b, nil
}

// loadAllIcons loads and returns icon resources for all the Markdown editor buttons
func loadAllIcons() (ToolbarIcons, error) {

	// TODO: there has to be a better way to do this without the repeated error checks
	h1Icon, err := loadIcon("h1.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h2Icon, err := loadIcon("h2.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h3Icon, err := loadIcon("h3.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	boldIcon, err := loadIcon("bold.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	italicIcon, err := loadIcon("italic.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	linkIcon, err := loadIcon("link.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	imageIcon, err := loadIcon("image.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	quoteBlockIcon, err := loadIcon("quote-block.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	codeBlockIcon, err := loadIcon("code-block.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	pageBreakIcon, err := loadIcon("page-break.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	undoIcon, err := loadIcon("undo.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	redoIcon, err := loadIcon("redo.png")
	if err != nil {
		return ToolbarIcons{}, err
	}

	return ToolbarIcons{
		H1Icon:         fyne.NewStaticResource("H1Icon", h1Icon),
		H2Icon:         fyne.NewStaticResource("H2Icon", h2Icon),
		H3Icon:         fyne.NewStaticResource("H3Icon", h3Icon),
		BoldIcon:       fyne.NewStaticResource("BoldIcon", boldIcon),
		ItalicIcon:     fyne.NewStaticResource("ItalicIcon", italicIcon),
		LinkIcon:       fyne.NewStaticResource("LinkIcon", linkIcon),
		ImageIcon:      fyne.NewStaticResource("ImageIcon", imageIcon),
		QuoteBlockIcon: fyne.NewStaticResource("QuoteBlockIcon", quoteBlockIcon),
		CodeBlockIcon:  fyne.NewStaticResource("CodeBlockIcon", codeBlockIcon),
		PageBreakIcon:  fyne.NewStaticResource("PageBreakIcon", pageBreakIcon),
		UndoIcon:       fyne.NewStaticResource("UndoIcon", undoIcon),
		RedoIcon:       fyne.NewStaticResource("RedoIcon", redoIcon),
	}, nil
}
