// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// go test -v search_test.go

// Tests for the search package
package tests

import (
	"testing"

	"github.com/ArdanStudios/gotraining/feed_app/sample/search"
)

// testMatcher implements the test matcher.
type testMatcher struct{}

// Search implements the behavior for the test matcher.
func (tm testMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	result := search.Result{
		Field:   "Test Field",
		Content: "Test Content",
	}

	return []*search.Result{&result}, nil
}

// init registers the test matcher with the program.
func init() {
	var matcher testMatcher
	search.Register("test", matcher)
}

// TestRegistration tests the registration of matchers is functioning.
func TestRegistration(t *testing.T) {
	matcher := search.FindMatcher("test")

	t.Log("The test Matcher should be returned.")
	if _, ok := matcher.(testMatcher); !ok {
		t.Fatalf("ERROR : Invalid Matcher : %T", matcher)
	}
}

// TestSearch tests the search mechanism for using a Matcher
// with the channel response.
func TestSearch(t *testing.T) {
	matcher := search.FindMatcher("test")
	results := make(chan *search.Result, 1)
	var feed search.Feed
	var searchTerm string

	search.Match(matcher, &feed, searchTerm, results)
	result := <-results

	t.Log("The search should return the test results.")
	if result.Field != "Test Field" {
		t.Errorf("ERROR : Expecting[Test Field] Received[%s]", result.Field)
	}

	if result.Content != "Test Content" {
		t.Errorf("ERROR : Expecting[Test Content] Received[%s]", result.Content)
	}
}

// TestRetrieveFeeds tests the application can read the JSON document file
// of feeds.
func TestRetrieveFeeds(t *testing.T) {
	feeds, err := search.RetrieveFeeds("../")

	t.Log("No error should be reported after the call.")
	if err != nil {
		t.Fatalf("ERROR : %s", err)
	}

	t.Log("Feed documents should be returned.")
	if len(feeds) == 0 {
		t.Fatalf("ERROR : Expecting[>0] Received[%d]", len(feeds))
	}

	t.Log("Feed documents with empty strings should not exist.")
	for _, feed := range feeds {
		if feed.Name == "" {
			t.Error("ERROR : Expecting Name to not be blank.")
		}

		if feed.Type == "" {
			t.Error("ERROR : Expecting Type to not be blank.")
		}

		if feed.URI == "" {
			t.Error("ERROR : Expecting URI to not be blank.")
		}
	}
}
