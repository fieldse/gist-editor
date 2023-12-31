// Modal for configuring the Github API token
package ui

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

// Path under the user's home directory for storing app configuration
var CONFIG_DIR_NAME = "./gist-editor"                                    // FIXME: should not be hardcoded
var GITHUB_CONFIG_FILE = path.Join(userConfigPath(), "github-token.txt") // FIXME: Improve file format

// GithubSettingsWindow is the the Github config settings window
type GithubSettingsWindow struct {
	dialog     *dialog.FormDialog // the main form
	tokenField *widget.Entry      // the Github token field
}

// New returns a new instance of the GithubSettingsWindow
func (g GithubSettingsWindow) New(cfg *AppConfig) *GithubSettingsWindow {
	w, entry := githubSettingsUI(cfg)
	return &GithubSettingsWindow{
		dialog:     w,
		tokenField: entry,
	}
}

// Load reads the github setting from file, and stores to the app config.
func (g *GithubSettingsWindow) Load(cfg *AppConfig) error {
	token, err := ReadGithubToken()
	if err != nil {
		logger.Error("load Github settings failed", err)
		return err
	}
	if token == "" {
		logger.Info("no github token found")
		return nil
	}
	logger.Info("github token loaded from %s", GITHUB_CONFIG_FILE)
	g.tokenField.SetText(token) // update the UI placeholder
	cfg.GithubConfig.GithubAPIToken = token
	return nil
}

// Show shows the Github settings modal
func (g GithubSettingsWindow) Show() {
	g.dialog.Show()
}

// Hide hides the Github settings modal
func (g GithubSettingsWindow) Hide() {
	g.dialog.Hide()
}

// githubSettingsUI generates the form and modal for Githut settings.
// Returns the form dialog and the Github token entry field.
func githubSettingsUI(cfg *AppConfig) (*dialog.FormDialog, *widget.Entry) {
	w := cfg.MainWindow.Window
	input := widget.NewEntry()
	input.PlaceHolder = "Enter your Github API token..."
	var tempVal string = ""
	input.OnChanged = func(s string) {
		tempVal = s
	}
	input.OnSubmitted = func(s string) {
		tempVal = s
	}
	// Onsave for the form dialog
	onSave := func(b bool) {
		if !b { // if the user cancels the dialog
			return
		}
		var originalVal = input.Text
		err := saveToken(tempVal)
		if err != nil {
			d := dialog.NewError(fmt.Errorf("error saving token: %s", err.Error()), w)
			d.Show()
			input.SetText(originalVal) // Reset to original state
			return
		}
		d := dialog.NewInformation("Github token saved", "Github token updated", w)
		d.Show()
		logger.Debug("Github API token saved: %v", tempVal)
		cfg.GithubConfig.GithubAPIToken = tempVal
	}
	var formItems []*widget.FormItem
	formItems = append(formItems, widget.NewFormItem("Github API token", input))
	d := dialog.NewForm("Github API Token", "Save", "Cancel", formItems, onSave, w)
	d.Resize(fyne.NewSize(400, 200))
	return d, input
}

// ReadGithubToken reads and returns the Github API token, if it exists.
func ReadGithubToken() (string, error) {
	if !fileExists(GITHUB_CONFIG_FILE) {
		logger.Info("load github settings: no file found.")
		return "", nil
	}
	token, err := readToken()
	if err != nil {
		return "", err
	}
	return token, nil
}

// Regex to validate the token is alphanumeric
var rgx = regexp.MustCompile("^[A-Za-z0-9]*$")

// saveToken saves the Github token to a local file
func saveToken(token string) error {
	if token == "" {
		return fmt.Errorf("token is empty")
	}
	if !rgx.MatchString(token) {
		return fmt.Errorf("token must be alphanumeric characters")
	}
	// Ensure the user config dir exists
	configDir := userConfigPath()
	if !dirExists(configDir) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			return fmt.Errorf("create user config dir failed: %w", err)
		}
		logger.Info("created user config directory at %s", configDir)
	}
	// Save the token to file
	data := []byte(token)
	err := os.WriteFile(GITHUB_CONFIG_FILE, data, 0644)
	if err != nil {
		return fmt.Errorf("save token to file failed: %w", err)
	}
	logger.Debug("saved Github token: %s", GITHUB_CONFIG_FILE)
	return nil
}

// readToken reads the Github token from a local file
// returns error on empty
func readToken() (string, error) {
	if !fileExists(GITHUB_CONFIG_FILE) {
		return "", fmt.Errorf("config file not found: %s", GITHUB_CONFIG_FILE)
	}
	f, err := os.Open(GITHUB_CONFIG_FILE)
	if err != nil {
		return "", fmt.Errorf("error reading file: %s -- err: %s", GITHUB_CONFIG_FILE, err)
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	return buf.String(), nil
}

// fileExists checks if a file exists at the given path.
func fileExists(fp string) bool {
	data, err := os.Stat(fp)
	return err == nil && !data.IsDir()
}

// dirExists checks if a directory exists at the given path
func dirExists(fp string) bool {
	data, err := os.Stat(fp)
	return err == nil && data.IsDir()
}

// userConfigPath returns the save path for config settings
func userConfigPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return path.Join(user.HomeDir, CONFIG_DIR_NAME)
}
