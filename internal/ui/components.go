// Miscellaneous small UI components
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func TitleText(msg string) fyne.Widget {
	l := widget.NewLabel("Your gists")
	l.TextStyle.Bold = true
	return l
}
