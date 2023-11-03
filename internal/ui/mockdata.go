// Contains mock data for the Gists view
package ui

import (
	"fmt"
	"time"
)

// Example structure for a Github Gist.
// This is just a placeholder
type Gist struct {
	id       string
	filename string
	slug     string
	content  string
	authorId string
	createAt time.Time
}

var exampleGist = Gist{
	id:       "example-gist",
	slug:     "example-gist",
	filename: "Example Gist.md",
	content:  "This is an example Gist placeholder",
	authorId: "example-author",
	createAt: time.Now(),
}

func newExampleGist(id int) Gist {
	e := exampleGist
	e.id = fmt.Sprintf("example-gist-%d", id)
	e.content = e.content + " -- gist ID: " + e.id
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
