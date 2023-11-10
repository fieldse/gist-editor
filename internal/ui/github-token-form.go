// Modal for configuring the Github API token
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

func GithubTokenModal(cfg *AppConfig, w fyne.Window) *dialog.FormDialog {
	input := widget.NewEntry()
	input.PlaceHolder = "Enter your Github API token..."
	var tempVal string = ""
	input.OnChanged = func(s string) {
		tempVal = s
	}
	input.OnSubmitted = func(s string) {
		tempVal = s
	}
	// Onsave for the form dialog
	onSave := func(b bool) {
		logger.Debug("onSave called: bool value: %t", b)
		if !b {
			return
		}
		logger.Debug("onSave called: Github API token: %v", tempVal)
		cfg.GithubConfig.GithubAPIToken = tempVal
	}
	var formItems []*widget.FormItem
	formItems = append(formItems,
		widget.NewFormItem("Github API token", input),
	)
	d := dialog.NewForm("Github API Token", "Save", "Cancel", formItems, onSave, w)
	d.Resize(fyne.NewSize(400, 200))
	return d
}
