package github

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GithubTokenResponse is the JSON response from the authorize request
// Expected format:
//
//	{
//	  "access_token":"gho_16C7e42F292c6912E7710c838347Ae178B4a",
//	  "scope":"repo,gist",
//	  "token_type":"bearer"
//	}
type GithubTokenResponse struct {
	AccessToken string
	Scope       string
	TokenType   string
}

var (
	// Application ID and secret, registered at Github
	GithubClientID     = os.Getenv("GITHUB_OAUTH_CLIENT_KEY")
	GithubClientSecret = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")

	// Unique state token passed at time of Oauth flow
	stateToken string = ""

	// GithubAuthToken is the token we get back from the GitHub authorization request
	GithubAuthToken *oauth2.Token

	IsAuthorized bool // true if we have completed the oAuth flow
	UserProfile  GithubUserProfileData

	// Our logged-in HTTP client with authorization
	Client *http.Client
)

const (
	// Callback listener URL and port
	HostIP   string = "127.0.0.1"
	HostPort string = "9090"
)

var HostAndPort string = fmt.Sprintf("%s:%s", HostIP, HostPort)

// The permission scopes we need for the app
// Details on GitHub oauth scopes:
// https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/scopes-for-oauth-apps
var requiredGithubScopes []string = []string{
	"read:user",  // read user profile data
	"user:email", // read user email
	"gist",       // write access to user gists
}

// OAuth struct for Github authorization
var githubConfig = &oauth2.Config{
	RedirectURL: fmt.Sprintf("http://%s/%s", HostAndPort, urls.Authenticate),
	Scopes:      requiredGithubScopes,
	Endpoint:    github.Endpoint,
}

// StartServer starts the http server and listens for the configured endpoints.
func StartServer() {
	http.HandleFunc(urls.Index, index)
	http.HandleFunc(urls.Authenticate, authenticate)
	http.HandleFunc(urls.Callback, completeGithubOauth)
	http.HandleFunc(urls.Success, loginSuccessful)

	http.ListenAndServe(HostAndPort, nil)
}

// newClient returns a new http client with authorization from the given token
func NewClient(token *oauth2.Token, r *http.Request) *http.Client {
	ts := githubConfig.TokenSource(r.Context(), token)
	return oauth2.NewClient(r.Context(), ts)
}

// OpenBrowser opens a new browser in the default system browser
func OpenBrowser(url string) error {
	return browser.OpenURL(url)
}

// generate a random state token to prevent XSRF attacks
func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
