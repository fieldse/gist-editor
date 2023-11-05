// Contains mock data for the Gists view
package mockdata

import (
	"fmt"
	"time"
)

// Example structure for a Github Gist.
// This is just a placeholder
type Gist struct {
	ID       string
	Slug     string
	Filename string
	Content  string
	AuthorId string
	CreateAt time.Time
}

var exampleGist = Gist{
	ID:       "example-gist",
	Slug:     "example-gist",
	Filename: "Example Gist.md",
	Content:  "This is an example Gist placeholder",
	AuthorId: "example-author",
	CreateAt: time.Now(),
}

func newExampleGist(id int) Gist {
	e := exampleGist
	e.ID = fmt.Sprintf("example-gist-%d", id)
	e.Filename = fmt.Sprintf("Example Gist-%d", id)
	e.Content = e.Content + " -- gist ID: " + e.ID
	return e
}

// Some mock Gist data for the Gists view
var MockGistData = []Gist{
	newExampleGist(1),
	newExampleGist(2),
	newExampleGist(3),
	newExampleGist(4),
	newExampleGist(5),
	newExampleGist(5),
}
