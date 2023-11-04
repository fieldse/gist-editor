package ui

import (
	// mockdata "github.com/fieldse/gist-editor/internal/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// var data = mockdata.MockGistData

var data = []string{"foo", "bar", "baz"}

// FIXME: implement
func createItem() fyne.CanvasObject {
	var item fyne.CanvasObject
	return item
	// TODO
}

// FIXME: implement updates
func updateItem(itemId widget.ListItemID, o fyne.CanvasObject) {
	o.(*widget.Label).SetText(data[itemId])
}

func listLength() int {
	return len(data)
}

// Returns a list view widget of all user gists
func ListWidget() *fyne.Container {
	title := widget.NewLabel("Your gists")
	title.TextStyle.Bold = true

	// List widget
	var exampleList = widget.NewList(
		listLength,
		createItem,
		updateItem,
	)

	// Example content widget
	vBox := container.NewVBox(title, exampleList)

	content := container.NewHBox(title, vBox)
	return content
}
