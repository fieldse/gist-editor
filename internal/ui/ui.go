package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/github"
)

// Basic app structure, with windows and other data to be passed around
type AppConfig struct {
	App          *fyne.App
	BaseWindow   *fyne.Window
	ListWindow   *fyne.Window
	EditWindow   *fyne.Window
	showEditView func()
	showListView func()
	// openFileFunc  func()
	// closeFileFunc func()
	exit        func()
	RunUI       func()
	CurrentFile github.Gist
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
	e := EditWindow(a)

	// Store the windows to cfg
	cfg.BaseWindow = &w
	cfg.ListWindow = &l
	cfg.EditWindow = &e

	// Connect the Show window functions
	cfg.showListView = l.Show
	cfg.showEditView = e.Show

	// Store the show window functions
	cfg.RunUI = func() { w.ShowAndRun() }
}

// Intitialize and return the base window
// TODO: find a better way to pass the Show window functions --
// not sure if it can be gotten directly from cfg before it's stored.
func BaseWindow(cfg *AppConfig) fyne.Window {
	a := *cfg.App

	// Generate new master window, set size, center it on screen
	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))
	w.SetMaster() // master window, when closed closes all other windows
	w.CenterOnScreen()

	// Store the exit function to Cfg
	cfg.exit = w.Close

	// Generate window content UI
	content := BaseView(cfg)
	w.SetContent(content)

	return w
}

// Base view upon opening the app
func BaseView(cfg *AppConfig) *fyne.Container {

	// Title
	title := TitleText("Welcome to the Gist editor!")
	subLabel := widget.NewLabel("Click View Gists to see all your gists, or create a new one with New Gist.")

	// Title and welcome text container
	titleContainer := container.NewVBox(title, subLabel)

	// Buttons for "View Gists" and "New Gist"
	viewGistsButton := widget.NewButton("View Gists", func() {
		cfg.showListView()
	})
	newGistButton := widget.NewButton("New Gist", func() {
		cfg.showEditView()
	})
	closeBtn := widget.NewButton("Exit", func() {
		cfg.exit()
	})

	// Centered buttons grid
	buttons := container.NewGridWithColumns(3, newGistButton, viewGistsButton, closeBtn)

	spacer := layout.NewSpacer()

	// Vertical grid layout
	content := container.New(
		layout.NewGridLayoutWithRows(5), titleContainer, spacer, spacer, spacer, buttons,
	)
	return content
}

func StartUI() {
	cfg.MakeUI()
	cfg.RunUI()
}
