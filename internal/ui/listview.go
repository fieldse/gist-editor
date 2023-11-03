package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Adds an item to the list
func addItem() fyne.CanvasObject {
	// TODO
	var item fyne.CanvasObject
	return item
}

// Updates a single item in the list
func updateItem(itemId widget.ListItemID, item fyne.CanvasObject) {
	// TODO:
}

// Placeholder function for length() argument to NewList
func listLength() int { return 14 }

// Returns a list view widget of all user gists
func ListWidget() *fyne.Container {
	title := widget.NewLabel("Your gists")
	title.TextStyle.Bold = true

	// l := widget.NewLabel("Example Gist")
	exampleList := widget.NewList(listLength, addItem, updateItem)

	// Example content widget
	vBox := container.NewVBox(title, exampleList)

	content := container.NewHBox(title, vBox)
	return content
}
