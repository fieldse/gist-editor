package github

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fieldse/gist-editor/internal/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

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

// Response from Github profile request
type GithubUserProfileData struct {
	Viewer struct { // FIXME
		ID    string // JWT identifier
		Email string
	}
}

var (
	// Application ID and secret, registered at Github
	GithubClientID     = os.Getenv("GITHUB_OAUTH_CLIENT_KEY")
	GithubClientSecret = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")

	// Unique state token passed at time of Oauth flow
	stateToken string = generateStateToken()

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

// OAuth struct for Github authorization
// Details on required GitHub oauth scopes:
// https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/scopes-for-oauth-apps
var gitHubOAuthConfig = &oauth2.Config{
	RedirectURL: fmt.Sprintf("http://%s/authorize", HostAndPort),
	Scopes: []string{
		"read:user",  // read user profile data
		"user:email", // read user email
		"gist",       // write access to user gists
	},
	Endpoint: github.Endpoint,
}

// startServer starts the http server and listens for the configured endpoints.
func startServer() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
	http.HandleFunc("/success", loginSuccessful)

	http.ListenAndServe(HostAndPort, nil)
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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<form action="/oauth/github" method="post">
		<input type="submit" value="Login with Github">
	</form>
</body>
</html>`)
}

func loginSuccessful(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<h2>Login successful! You may now return to the application.</h2>
</body>
</html>`)
}

// startGithubOauth is a handler to start the authorization flow.
// This creates a redirect response with the Auth URL and the following parameters:
//   - response_type 	(optional)
//   - client_id 		(required)
//   - redirect_uri 	(optional: this is specified in Github oAuth app settings)
//   - scope 			(optional: default is read public profile info)
//   - state			(required: state token)
func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := gitHubOAuthConfig.AuthCodeURL(stateToken)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// completeGithubOauth finishes the authorization flow
func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	// Check the received state code matches
	if state != stateToken {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	// Exchange the code for a token
	token, err := gitHubOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "login failed", http.StatusInternalServerError)
		return
	}

	// Store the token for use later
	logger.Info("github login successful: token stored %v", token)
	GithubAuthToken = token
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

// newClient returns a new http client with authorization from the given token
func newClient(token *oauth2.Token, r *http.Request) *http.Client {
	ts := gitHubOAuthConfig.TokenSource(r.Context(), token)
	return oauth2.NewClient(r.Context(), ts)
}

// getUserData fetches user basic profile data
func getUserData(client *http.Client) (GithubUserProfileData, error) {

	// GraphQL query
	// 	FIXME: convert this to REST API request
	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		logger.Error("get user profile data failed", err)
		return GithubUserProfileData{}, err
	}
	defer resp.Body.Close()

	var p GithubUserProfileData
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		logger.Error("invalid Github response", err)
		return GithubUserProfileData{}, err
	}
	return p, nil

}

// generate a random state token to prevent XSRF attacks
func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// generateExpiration returns a date one year from the current time
func generateExpiration() time.Time {
	return time.Now().Add(365 * 24 * time.Hour)
}

func storeGithubToken(token string) error {
	// TODO -- store the token somewhere.
	return fmt.Errorf("not yet implemented")
}
