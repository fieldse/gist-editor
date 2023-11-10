// Functions for open, close, and save files
package ui

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/fieldse/gist-editor/internal/logger"
)

// Open a file from disk in the editor
func openFile(cfg *AppConfig) {

	w := *cfg.BaseWindow // parent window
	openFileFunc := func(read fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if read == nil {
			return
		}
		defer read.Close()
		data, err := io.ReadAll(read)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		cfg.CurrentFile.gist.Content = string(data)
		cfg.CurrentFile.localFilepath = read.URI().Path()
		cfg.CurrentFile.isOpen = true
	}
	openFileDialog := dialog.NewFileOpen(openFileFunc, w)
	openFileDialog.SetFilter(filter)
	openFileDialog.Show()
}

// Save the currently open file
func saveFile() {
	logger.Debug("todo -- save file...")
}

// Save the currently open file with new filename
func saveFileAs() {
	logger.Debug("todo -- save as...")
}

// Close the currently open file
func closeFile() {
	logger.Debug("todo -- close file...")
}
