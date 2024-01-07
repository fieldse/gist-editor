package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
