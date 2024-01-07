package github

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Client ID for github app
const GithubAppID string = "b9cbbaa5e7c0f0644796"

// Callback listener URL and port
const TokenCallbackHostPort string = "http://127.0.0.1:9090"

var oAuthGithubLogin = &oauth2.Config{
	RedirectURL:  TokenCallbackHostPort + "/authorize",
	ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_KEY"),
	ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     github.Endpoint,
}

// startCallbackListener starts the listener for the Github authorization callback
func startCallbackListener() {
	http.HandleFunc("/authorize", callbackHandler)

	// Listen and serve on a specific port
	log.Fatal(http.ListenAndServe(TokenCallbackHostPort, nil))
}

// callbackHandler handles the Github authorization request callback
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Process the callback request here
	fmt.Fprintf(w, "Callback Received!")
}
