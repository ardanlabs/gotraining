// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package search : bing performs searches against the bing search engine.
package search

import (
	"log"
	"strings"
)

// Bing provides support for Bing searches.
type Bing struct{}

// NewBing returns a Bing Searcher value.
func NewBing() Searcher {
	return Bing{}
}

// Search implements the Searcher interface. It performs a search
// against Bing.
func (b Bing) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Bing : Seared : Started : searchTerm[%s]\n", searchTerm)

	// Build a proper search url.
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	uri := "http://www.bing.com/search?q=" + searchTerm + "&format=rss"

	// Perform a RSS search.
	rssSearch("Bing", uri, searchResults)

	log.Println("Bing : Search : Completed")
}
