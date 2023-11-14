// Functions for open, close, and save files
package ui

import (
	"io"
	"path"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/fieldse/gist-editor/internal/github"
	"github.com/fieldse/gist-editor/internal/logger"
)

// A GistFile represents a currently open local or remote markdown file
type GistFile struct {
	Gist      *github.Gist
	isLocal   bool   // true if this is a local file from disk
	localURI  string // path to file resource: may differ on different OSs
	isOpen    bool
	isDirty   bool
	lastSaved time.Time
}

// Openable filetypes  filter
var filter = storage.NewExtensionFileFilter([]string{".md", ".txt"})

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
		filePath := read.URI().Path()
		fileName := read.URI().Name()

		// Initialize a new Gist from the data
		g := github.Gist{}.New(fileName, string(data))
		cfg.CurrentFile = &GistFile{
			Gist:     &g,
			isLocal:  true,
			isOpen:   true,
			localURI: path.Join(filePath, fileName),
		}
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
func closeFile(cfg *AppConfig) {
	cfg.CurrentFile.Gist = &github.Gist{}
	cfg.CurrentFile.isOpen = false
	cfg.CurrentFile.isLocal = false
	cfg.CurrentFile.isDirty = false
	cfg.Editor.SetText("") // clear the editor text
}
