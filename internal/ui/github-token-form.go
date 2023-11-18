// Modal for configuring the Github API token
package ui

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fieldse/gist-editor/internal/logger"
)

// Path under the user's home directory for storing app configuration
var CONFIG_DIR_NAME = "./gist-editor"                                    // FIXME: should not be hardcoded
var GITHUB_CONFIG_FILE = path.Join(userConfigPath(), "github-token.txt") // FIXME: Improve file format

func GithubTokenModal(cfg *AppConfig) *dialog.FormDialog {
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
		if !b {
			return
		}
		err := saveToken(tempVal)
		if err != nil {
			d := dialog.NewError(err, w)
			d.Show()
			return
		}
		d := dialog.NewInformation("Github token saved", "Github token updated", w)
		d.Show()
		logger.Debug("Github API token saved: %v", tempVal)
		cfg.GithubConfig.GithubAPIToken = tempVal
	}
	var formItems []*widget.FormItem
	formItems = append(formItems,
		widget.NewFormItem("Github API token", input),
	)
	d := dialog.NewForm("Github API Token", "Save", "Cancel", formItems, onSave, w)
	d.Resize(fyne.NewSize(400, 200))
	return d
}

// ReadGithubToken reads and returns the Github API token, if it exists.
func ReadGithubToken() (string, error) {
	if !fileExists(GITHUB_CONFIG_FILE) {
		return "", nil
	}
	token, err := readToken()
	if err != nil {
		logger.Error("read github config file failed", err)
		return "", err
	}
	return token, nil
}

// saveToken saves the Github token to a local file
func saveToken(token string) error {
	if token == "" {
		return fmt.Errorf("token is empty")
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
	logger.Info("saved Github token: %s", GITHUB_CONFIG_FILE)
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
