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

// Save saves a Gist file to local storage
func (g *GistFile) Save() {
	logger.Debug("todo -- save file...")
	g.lastSaved = time.Now()
	g.isDirty = false
}

// SaveAs saves a Gist file to local storage with a new filename
func (g *GistFile) SaveAs() {
	logger.Debug("todo -- save as...")
	g.lastSaved = time.Now()
	g.isDirty = false
}

// Close clears a Gist file to empty and marks as closed
func (g *GistFile) Close() {
	g.Gist = &github.Gist{}
	g.isOpen = false
	g.isLocal = false
	g.isDirty = false
}

// Openable filetypes  filter
var filter = storage.NewExtensionFileFilter([]string{".md", ".txt"})

// openFile is the opener function passed to the Open File dialog
func openFile(read fyne.URIReadCloser, err error) {
	w := cfg.MainWindow.Window
	if err != nil {
		logger.Error("open file failed", err)
		dialog.ShowError(err, w)
		return
	}
	if read == nil {
		logger.Debug("open file was canceled")
		return
	}
	defer read.Close()
	data, err := io.ReadAll(read)
	if err != nil {
		logger.Error("read file failed", err)
		dialog.ShowError(err, w)
		return
	}
	filePath := read.URI().Path()
	fileName := read.URI().Name()

	logger.Debug("open file succeeded: filename: %s", fileName)

	// Initialize a new Gist from the data
	g := github.Gist{}.New(fileName, string(data))
	cfg.CurrentFile = &GistFile{
		Gist:     &g,
		isLocal:  true,
		isOpen:   true,
		localURI: path.Join(filePath, fileName),
	}

	// Update the content of the editor window
	cfg.Editor.SetContent(g.Content)
	cfg.Editor.Title = fileName

	// Show the edit window
	cfg.MainWindow.SetCanSave(true)
	cfg.ShowEditWindow()
}
