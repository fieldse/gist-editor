package github

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fieldse/gist-editor/internal/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// UserGithubAuthToken is the token we get back from the GitHub authorization request
var UserGithubAuthToken string = ""

// Client ID for github app
const GithubAppID string = "b9cbbaa5e7c0f0644796"

// Callback listener URL and port
const TokenCallbackHostPort string = "127.0.0.1:9090"

// OAuth struct for Github authorization
// Details on required GitHub oauth scopes:
// https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/scopes-for-oauth-apps
var oAuthGithubLogin = &oauth2.Config{
	RedirectURL:  fmt.Sprintf("http://%s/authorize", TokenCallbackHostPort),
	ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_KEY"),
	ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
	Scopes: []string{
		"read:user",  // read user profile data
		"user:email", // read user email
		"gist",       // write access to user gists
	},
	Endpoint: github.Endpoint,
}

// StartCallbackListener starts the listener for the Github authorization callback
func StartCallbackListener() {
	http.HandleFunc("/authorize", callbackHandler)

	// Listen and serve on a specific port
	log.Fatal(http.ListenAndServe(TokenCallbackHostPort, nil))
}

// callbackHandler handles the Github authorization request callback
// Stores the token to UserGithubAuthToken on success.
func callbackHandler(w http.ResponseWriter, r *http.Request) {

	// FIXME -- get the response somehow
	var response = &http.Response{}
	token, err := parseResponse(response)
	if err != nil {
		logger.Error("callback handler failed: %v", err)
	}
	// Fixme: do something with this
	UserGithubAuthToken = token
	logger.Info("callback handler succeeded: token received: %s", token)
}

// GithubTokenResponse is the JSON response from the authorize request
// expected format:
// Accept: application/json
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

// parseResponse parses the API response JSON and returns the access token
func parseResponse(r *http.Response) (string, error) {
	if r.StatusCode != 200 {
		return "", fmt.Errorf("invalid API response code: %d", r.StatusCode)
	}
	if r.ContentLength <= 0 {
		return "", fmt.Errorf("invalid API response length: %d", r.ContentLength)
	}
	// Read response to byte slice
	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %v", err)
	}
	// Unmarshal JSON to struct
	asJson := GithubTokenResponse{}
	err = json.Unmarshal(body, &asJson)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response JSON: %v", err)
	}
	return asJson.AccessToken, nil
}

func createOAuthLoginURL(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	u := oAuthGithubLogin.AuthCodeURL(oauthState)

	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

// generate a random state token to prevent XSRF attacks
func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// TODO - make the OAuth request to GitHub
// copied from https://github.com/douglasmakey/oauth2-example
// Unsure what the code parameter is supposed to be
func getAuthToken(code string) (string, error) {

	// Use code to get token and get user info from Google.
	token, err := oAuthGithubLogin.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("oauth exchange failed: %v", err)
	}
	return token.AccessToken, nil
}

// generate a cookie to attach to the OAuth request
func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	var state = generateStateToken() // random state token to prevent XSRF attacks
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
