// handlers for the oAuth endpoints
package github

import (
	"fmt"
	"net/http"

	"github.com/fieldse/gist-editor/internal/logger"
)

// urls are url paths for the various handler
var urls = struct {
	Index        string
	Authenticate string
	Callback     string
	Success      string
}{
	Index:        "/",
	Authenticate: "/authorize",
	Callback:     "/callbacks",
	Success:      "/success",
}

// index serves the base link prompting the user to go authenticate.
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<form action="%s" method="post">
		<input type="submit" value="Login with Github">
	</form>
</body>
</html>`, urls.Authenticate)
}

// loginSuccessful displays a "login success!" message and directs user back
// to the application
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

// authenticate is a handler to start the authorization flow.
// This creates a redirect response with the Auth URL and the following parameters:
//   - response_type 	(optional)
//   - client_id 		(required)
//   - redirect_uri 	(optional: this is specified in Github oAuth app settings)
//   - scope 			(optional: default is read public profile info)
//   - state			(required: state token)
func authenticate(w http.ResponseWriter, r *http.Request) {
	stateToken = generateStateToken()
	redirectURL := githubConfig.AuthCodeURL(stateToken)
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
	token, err := githubConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "login failed", http.StatusInternalServerError)
		return
	}

	// Store the token for use later
	logger.Info("github login successful: token stored %v", token)
	GithubAuthToken = token
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}
