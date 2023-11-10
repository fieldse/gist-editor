// Modal for configuring the Github API token
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// FIXME: placeholder save function
func onSave(bool) {}

func GithubTokenModal(w fyne.Window) *dialog.FormDialog {
	input := widget.NewEntry()
	var formItems []*widget.FormItem
	formItems = append(formItems,
		widget.NewFormItem("Enter Github API token", input),
	)
	d := dialog.NewForm("Github API Token", "Save", "Cancel", formItems, onSave, w)
	return d
}
