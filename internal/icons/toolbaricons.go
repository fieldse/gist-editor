// bundled icons for the toolbar

package icons

import "fyne.io/fyne/v2"

type toolbarIcons = struct {
	BoldIcon          *fyne.StaticResource
	ChecklistIcon     *fyne.StaticResource
	CodeBlockIcon     *fyne.StaticResource
	DeleteIcon        *fyne.StaticResource
	EraserIcon        *fyne.StaticResource
	FileAddIcon       *fyne.StaticResource
	FileCopyIcon      *fyne.StaticResource
	FileEditIcon      *fyne.StaticResource
	FolderStarIcon    *fyne.StaticResource
	FolderIcon        *fyne.StaticResource
	H1Icon            *fyne.StaticResource
	H2Icon            *fyne.StaticResource
	H3Icon            *fyne.StaticResource
	H4Icon            *fyne.StaticResource
	ImageIcon         *fyne.StaticResource
	ItalicIcon        *fyne.StaticResource
	LinkIcon          *fyne.StaticResource
	ListBulletIcon    *fyne.StaticResource
	ListNumberedIcon  *fyne.StaticResource
	PageBreakIcon     *fyne.StaticResource
	QuoteBlockIcon    *fyne.StaticResource
	RedoIcon          *fyne.StaticResource
	StrikethroughIcon *fyne.StaticResource
	UnderlineIcon     *fyne.StaticResource
	UndoIcon          *fyne.StaticResource
}

var ToolbarIcons toolbarIcons = toolbarIcons{
	BoldIcon:          resourceBoldPng,
	ChecklistIcon:     resourceChecklistPng,
	CodeBlockIcon:     resourceCodeBlockPng,
	DeleteIcon:        resourceDeletePng,
	EraserIcon:        resourceEraserPng,
	FileAddIcon:       resourceFileAddPng,
	FileCopyIcon:      resourceFileCopyPng,
	FileEditIcon:      resourceFileEditPng,
	FolderStarIcon:    resourceFolderStarPng,
	FolderIcon:        resourceFolderPng,
	H1Icon:            resourceH1Png,
	H2Icon:            resourceH2Png,
	H3Icon:            resourceH3Png,
	H4Icon:            resourceH4Png,
	ImageIcon:         resourceImagePng,
	ItalicIcon:        resourceItalicPng,
	LinkIcon:          resourceLinkPng,
	ListBulletIcon:    resourceListBulletPng,
	ListNumberedIcon:  resourceListNumberedPng,
	PageBreakIcon:     resourcePageBreakPng,
	QuoteBlockIcon:    resourceQuoteBlockPng,
	RedoIcon:          resourceRedoPng,
	StrikethroughIcon: resourceStrikethroughPng,
	UnderlineIcon:     resourceUnderlinePng,
	UndoIcon:          resourceUndoPng,
}
