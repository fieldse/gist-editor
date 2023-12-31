// Data types for the Gists
package github

import (
	"time"

	"github.com/avelino/slugify"
)

type GithubConfig struct {
	GithubAPIToken string
}

// Example structure for a Github Gist.
// This is just a placeholder
type Gist struct {
	ID        string
	Slug      string
	Filename  string
	Content   string
	AuthorId  string
	CreatedAt time.Time
}

// Generate a new Gist
func (g Gist) New(fileName string, content string) Gist {
	slug := slugify.Slugify(fileName)
	return Gist{
		ID:        slug,
		Slug:      slug,
		Filename:  fileName,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
