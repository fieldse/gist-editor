// Data types for the Gists
package github

import (
	"time"

	"github.com/avelino/slugify"
)

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
func (g Gist) New(fileName string) *Gist {
	slug := slugify.Slugify(fileName)
	return &Gist{
		ID:        slug,
		Slug:      slug,
		Filename:  fileName,
		CreatedAt: time.Now(),
	}
}
