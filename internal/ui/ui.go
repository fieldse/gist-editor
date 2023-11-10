package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/fieldse/gist-editor/internal/github"
)

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	App         *fyne.App
	BaseWindow  *fyne.Window
	ListWindow  *fyne.Window
	EditWindow  *fyne.Window
	RunUI       func()
	CurrentGist github.Gist
	CurrentFile GistFile
}

// A GistFile represents a currently open markdown file
type GistFile struct {
	gist      github.Gist
	isOpen    bool
	isDirty   bool
	lastSaved time.Time
}

var cfg AppConfig

// Generate and store the basic UI components
func (cfg *AppConfig) MakeUI() {

	// Create app
	a := app.New()
	cfg.App = &a

	// Create base view UI.
	// This is initialized last, because the buttons require the List and Edit views
	// to be initialized and stored, to attach their Show functions.
	w := BaseWindow(cfg)

	// Create Gists list window
	l := ListWindow(a)

	// Create Edit view window
	e := EditWindow(cfg)

	// Create the main menu
	m := FileMenu(cfg)
	w.SetMainMenu(m)

	// Store the windows to cfg
	cfg.BaseWindow = &w
	cfg.ListWindow = l
	cfg.EditWindow = e

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

// Exit the application
func (cfg *AppConfig) Exit() {
	w := *cfg.BaseWindow
	w.Close()
}

// OpenFile opens a local markdown file
func (cfg *AppConfig) OpenFile() {
	// TODO
	openFile()
	cfg.CurrentFile.gist = github.ExampleGist
	cfg.CurrentFile.isOpen = true
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
