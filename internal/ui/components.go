// Miscellaneous small UI components
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// A text component with bold styling
// TODO: Figure out how to increase font size
func TitleText(msg string) fyne.Widget {
	l := widget.NewLabel(msg)
	l.TextStyle.Bold = true
	return l
}

// Returns a title text container with some spacing
func TitleBox(msg string) *fyne.Container {
	t := TitleText(msg)
	spacer := layout.NewSpacer()
	c := container.NewGridWithRows(2, t, spacer)
	return c
}

// Button container with padding and right alignment
func ButtonContainer(numItems int, objects ...fyne.CanvasObject) *fyne.Container {
	return container.NewPadded(container.NewGridWithColumns(numItems, objects...))
}
