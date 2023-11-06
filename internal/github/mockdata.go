// Contains mock data for the Gists view
package github

import (
	"fmt"
	"time"
)

var ExampleGist = Gist{
	ID:        "example-gist",
	Slug:      "example-gist",
	Filename:  "Example Gist.md",
	Content:   "## Example Gist\n\nThis is an example Gist placeholder.\n\nA list:\n- item 1\n- item 2\n- item 3",
	AuthorId:  "example-author",
	CreatedAt: time.Now(),
}

func newExampleGist(id int) Gist {
	e := ExampleGist
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
