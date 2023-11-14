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
	App               *fyne.App
	BaseWindow        *fyne.Window
	ListWindow        *fyne.Window
	Editor            *widget.Entry // the content editor
	EditWindow        *fyne.Window
	editWindowVisible bool
	setCanSave        func(bool) // Toggle function to allow saving file
	GithubTokenModal  *dialog.FormDialog
	RunUI             func()
	CurrentFile       *GistFile
	GithubConfig      *github.GithubConfig
}

// New initializes a new AppConfig instance
func (AppConfig) New() AppConfig {
	// Initialize a new Fyne app
	a := app.New()
	return AppConfig{
		App:          &a,
		GithubConfig: &github.GithubConfig{},
		CurrentFile: &GistFile{
			Gist: &github.Gist{},
		},
	}
}

var cfg AppConfig

// Initialize a new app at load time
func init() {
	cfg = AppConfig{}.New()
}

// Generate and store the basic UI components
func (cfg *AppConfig) MakeUI() {

	// Create base view UI.
	// This is initialized last, because the buttons require the List and Edit views
	// to be initialized and stored, to attach their Show functions.
	w := BaseWindow(cfg)

	// Create Gists list window
	l := ListWindow(cfg)

	// Create Edit view window, and get the entry editor
	e, editor := EditWindow(cfg)

	// Create Github token modal
	g := GithubTokenModal(cfg, w)

	// Create the main menu
	m, setCanSave := FileMenu(cfg)
	w.SetMainMenu(m)

	// Store the windows to cfg
	cfg.BaseWindow = &w
	cfg.ListWindow = l
	cfg.EditWindow = e
	cfg.Editor = editor
	cfg.GithubTokenModal = g
	cfg.setCanSave = setCanSave

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
	cfg.editWindowVisible = true
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

// NewFile opens a new empty markdown editor
func (cfg *AppConfig) NewFile() {
	cfg.setCanSave(true)
	cfg.ShowEditWindow()
}

// OpenFile opens a local markdown file
func (cfg *AppConfig) OpenFile() {
	openFile(cfg)
	cfg.setCanSave(true)
	cfg.ShowEditWindow()
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

// CloseFile closes the currently open markdown file and closes the editor window
func (cfg *AppConfig) CloseFile() {
	closeFile(cfg)
	cfg.setCanSave(false)
	w := *cfg.EditWindow
	w.Hide()
}

func StartUI() {
	cfg.MakeUI()
	cfg.RunUI()
}
