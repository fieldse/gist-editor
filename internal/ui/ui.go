package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/fieldse/gist-editor/internal/github"
)

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	App               *fyne.App
	MainWindow        MainWindow
	ListWindow        *ListView
	Editor            *Editor
	editWindowVisible bool
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

	// Create base app window.
	cfg.MainWindow = MainWindow{}.New(cfg)

	// Create Gists list window
	cfg.ListWindow = ListView{}.New(cfg)

	// Create Edit view window, and get the entry editor
	cfg.Editor = Editor{}.New(cfg)

	// Create Github token modal
	g := GithubTokenModal(cfg)

	cfg.GithubTokenModal = g

	// Store the show window functions
	cfg.RunUI = cfg.MainWindow.ShowAndRun
}

// Show the All Gists list view
func (cfg *AppConfig) ShowListWindow() {
	cfg.ListWindow.Show()
}

// Show the Edit Gists view
func (cfg *AppConfig) ShowEditWindow() {
	cfg.editWindowVisible = true
	cfg.Editor.Show()
}

// Show the Github Token modal
func (cfg *AppConfig) ShowGithubTokenModal() {
	w := *cfg.GithubTokenModal
	w.Show()
}

// Exit the application
func (cfg *AppConfig) Exit() {
	cfg.MainWindow.Close()
}

// NewFile opens a new empty markdown editor
func (cfg *AppConfig) NewFile() {
	cfg.MainWindow.SetCanSave(true)
	g := github.Gist{}.New("New Gist.md", "Enter your content here...")
	cfg.CurrentFile = &GistFile{
		isLocal:  true,
		isOpen:   true,
		isDirty:  false,
		localURI: "",
		Gist:     &g,
	}
	cfg.Editor.SetContent(g.Content)
	cfg.Editor.Title = "New Gist"
	cfg.ShowEditWindow()
}

// OpenFile opens a local markdown file
func (cfg *AppConfig) OpenFile() {
	d := dialog.NewFileOpen(openFile, cfg.MainWindow.Window)
	d.SetFilter(filter)
	d.Resize(fyne.NewSize(800, 600))
	d.Show()
}

// SaveFile saves the currently open markdown file locally to disk
func (cfg *AppConfig) SaveFile() {
	cfg.CurrentFile.Save()
}

// SaveFileAs saves the currently open markdown file locally to disk with a new filename
func (cfg *AppConfig) SaveFileAs() {
	cfg.CurrentFile.SaveAs()
}

// CloseFile closes the currently open markdown file and closes the editor window
func (cfg *AppConfig) CloseFile() {
	cfg.MainWindow.SetCanSave(false)
	cfg.Editor.Clear() // clear the editor text and title
	cfg.CurrentFile.Close()
	cfg.Editor.Hide()
}

func StartUI() {
	cfg.MakeUI()
	cfg.RunUI()
}
