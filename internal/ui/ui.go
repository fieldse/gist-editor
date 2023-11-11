package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	App              *fyne.App
	BaseWindow       *fyne.Window
	ListWindow       *fyne.Window
	EditWindow       *fyne.Window
	Editor           *widget.Entry // the content editor
	GithubTokenModal *dialog.FormDialog
	RunUI            func()
	CurrentFile      GistFile
	GithubConfig     *github.GithubConfig
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

	// Create Edit view window, and get the entry editor
	e, editor := EditWindow(cfg)

	// Create Github token modal
	g := GithubTokenModal(cfg, w)

	// Create the main menu
	m := FileMenu(cfg)
	w.SetMainMenu(m)

	// Store the windows to cfg
	cfg.BaseWindow = &w
	cfg.ListWindow = l
	cfg.EditWindow = e
	cfg.Editor = editor
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

// OpenFile opens a local markdown file
func (cfg *AppConfig) OpenFile() {
	openFile(cfg)
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
	closeFile(cfg)
	cfg.Editor.SetText("") // clear the content of the editor
	w := *cfg.EditWindow
	w.Hide()
}

func StartUI() {
	cfg.MakeUI()
	cfg.RunUI()
}
