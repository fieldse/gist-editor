// Contains icons for the Markdown editor toolbar buttons
package icons

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/fieldse/gist-editor/assets"
)

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
	// FIXME: find a way to avoid this ugly redundancy
	h1Icon, err := loadIconResource("h1.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h2Icon, err := loadIconResource("h2.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	h3Icon, err := loadIconResource("h3.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	boldIcon, err := loadIconResource("bold.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	italicIcon, err := loadIconResource("italic.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	linkIcon, err := loadIconResource("link.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	imageIcon, err := loadIconResource("image.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	quoteBlockIcon, err := loadIconResource("quote-block.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	codeBlockIcon, err := loadIconResource("code-block.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	pageBreakIcon, err := loadIconResource("page-break.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	undoIcon, err := loadIconResource("undo.png")
	if err != nil {
		return ToolbarIcons{}, err
	}
	redoIcon, err := loadIconResource("redo.png")
	if err != nil {
		return ToolbarIcons{}, err
	}

	return ToolbarIcons{
		H1Icon:         h1Icon,
		H2Icon:         h2Icon,
		H3Icon:         h3Icon,
		BoldIcon:       boldIcon,
		ItalicIcon:     italicIcon,
		LinkIcon:       linkIcon,
		ImageIcon:      imageIcon,
		QuoteBlockIcon: quoteBlockIcon,
		CodeBlockIcon:  codeBlockIcon,
		PageBreakIcon:  pageBreakIcon,
		UndoIcon:       undoIcon,
		RedoIcon:       redoIcon,
	}, nil
}

// Load the asset data as a byte array, returning a Fyne resource
func loadIconResource(filename string) (*fyne.StaticResource, error) {
	data, ok := assets.IconMap[filename]
	if !ok {
		return &fyne.StaticResource{}, fmt.Errorf("failed to load icon data for %s", filename)
	}
	return fyne.NewStaticResource(filename, data), nil
}
