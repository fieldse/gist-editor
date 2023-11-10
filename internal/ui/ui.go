package ui

import (
	"io"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/fieldse/gist-editor/internal/github"
)

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	App              *fyne.App
	BaseWindow       *fyne.Window
	ListWindow       *fyne.Window
	EditWindow       *fyne.Window
	GithubTokenModal *dialog.FormDialog
	RunUI            func()
	CurrentFile      GistFile
	GithubConfig     *github.GithubConfig
}

// A GistFile represents a currently open markdown file
type GistFile struct {
	gist          github.Gist
	localFilepath string
	isOpen        bool
	isDirty       bool
	lastSaved     time.Time
}

var cfg AppConfig

// Generate and store the basic UI components
func (cfg *AppConfig) MakeUI() {

	// Create app
	a := app.New()
	cfg.App = &a
	cfg.GithubConfig = &github.GithubConfig{}

	// Create base view UI.
	// This is initialized last, because the buttons require the List and Edit views
	// to be initialized and stored, to attach their Show functions.
	w := BaseWindow(cfg)

	// Create Gists list window
	l := ListWindow(a)

	// Create Edit view window
	e := EditWindow(cfg)

	// Create Github token modal
	g := GithubTokenModal(cfg, w)

	// Create the main menu
	m := FileMenu(cfg)
	w.SetMainMenu(m)

	// Store the windows to cfg
	cfg.BaseWindow = &w
	cfg.ListWindow = l
	cfg.EditWindow = e
	cfg.GithubTokenModal = g

	// Store the show window functions
	cfg.RunUI = func() { w.ShowAndRun() }
}

// Show the All Gists list view
func (cfg *AppConfig) ShowListWindow() {
	w := *cfg.ListWindow
	w.Show()
}

// Show the Edit Gists view
func (cfg *AppConfig) ShowEditWindow() {
	w := *cfg.EditWindow
	w.Show()
}

// Show the Github Token modal
func (cfg *AppConfig) ShowGithubTokenModal() {
	w := *cfg.GithubTokenModal
	w.Show()
}

// Exit the application
func (cfg *AppConfig) Exit() {
	w := *cfg.BaseWindow
	w.Close()
}

// Openable filetypes  filter
var filter = storage.NewExtensionFileFilter([]string{".md", ".txt"})

// OpenFile opens a local markdown file
func (cfg *AppConfig) OpenFile() {
	// TODO
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

// SaveFile saves the currently open markdown file locally to disk
func (cfg *AppConfig) SaveFile() {
	// TODO
	saveFile()
	cfg.CurrentFile.lastSaved = time.Now()
	cfg.CurrentFile.isDirty = false
}

// SaveFileAs saves the currently open markdown file locally to disk with a new filename
func (cfg *AppConfig) SaveFileAs() {
	// TODO
	saveFileAs()
	cfg.CurrentFile.lastSaved = time.Now()
	cfg.CurrentFile.isDirty = false
}

// CloseFile closes the currently open markdown file
func (cfg *AppConfig) CloseFile() {
	// TODO
	cfg.CurrentFile.gist = github.Gist{}
	cfg.CurrentFile.isOpen = false
	closeFile()
}

func StartUI() {
	cfg.MakeUI()
	cfg.RunUI()
}
