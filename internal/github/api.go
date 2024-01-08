// Functions to call Github API endpoints
package github

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fieldse/gist-editor/internal/logger"
)

// Response from Github profile request
type GithubUserProfileData struct {
	Viewer struct { // FIXME
		ID    string // JWT identifier
		Email string
	}
}

// githubApiUrls are the REST API endpoints for the required Github data
var githubApiUrls = struct {
	UserProfile string
}{
	UserProfile: "https://api.github.com/graphql", // FIXME: this is graphql, not REST
}

// GetUserData fetches user basic profile data
func GetUserData(client *http.Client) (GithubUserProfileData, error) {

	// GraphQL query
	// 	FIXME: convert this to REST API request
	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)
	resp, err := client.Post(githubApiUrls.UserProfile, "application/json", requestBody)
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
