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
	BaseWindow   *fyne.Window
	ListWindow   *fyne.Window
	EditWindow   *fyne.Window
	showListView func()
	exit         func()
	RunUI        func()
	CurrentFile  github.Gist
}

var cfg AppConfig

func newGist() {
	// TODO
}

// Generate and store the basic UI components
func (cfg *AppConfig) MakeUI() {

	// Create app and base window
	a := app.New()
	w := a.NewWindow("GistEdit")
	w.Resize(fyne.NewSize(600, 400))
	w.SetMaster() // master window, when closed closes all other windows

	// Store the exit function
	cfg.exit = w.Close

	// Create Gists list window
	l := ListWindow(a)

	// Create Edit view window
	e := EditWindow(a)

	// Create base view UI
	content := BaseView(cfg, l.Show)
	w.SetContent(content)

	// Store the app windows to config
	cfg.BaseWindow = &w
	cfg.ListWindow = &l
	cfg.EditWindow = &e

	// Store the Show function for main window
	cfg.RunUI = func() { w.ShowAndRun() }

	// Store the showListView function
	cfg.showListView = func() { l.Show() }
}

// Base view upon opening the app
func BaseView(cfg *AppConfig, showList func()) *fyne.Container {

	// Title
	title := TitleText("Welcome to the Gist editor!")
	subLabel := widget.NewLabel("Click View Gists to see all your gists, or create a new one with New Gist.")

	// Title and welcome text container
	titleContainer := container.NewVBox(title, subLabel)

	// Buttons for "View Gists" and "New Gist"
	b1 := widget.NewButton("View Gists", showList)
	b2 := widget.NewButton("New Gist", newGist)
	closeBtn := widget.NewButton("Exit", func() {
		cfg.exit()
	})

	// Centered buttons grid
	buttons := container.NewGridWithColumns(3, b1, b2, closeBtn)

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
