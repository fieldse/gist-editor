package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Startui() {
	a := app.New()
	w := a.NewWindow("Gist editor")
	title := widget.NewLabel("Gist editor")
	text := widget.NewLabel("Hello Gist!")
	content := container.New(layout.NewVBoxLayout(), title, text)
	w.Resize(fyne.NewSize(600, 400))

	w.SetContent(content)
	w.ShowAndRun()

}
