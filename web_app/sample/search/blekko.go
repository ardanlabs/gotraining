// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package search : bing performs searches against the bing search engine.
package search

import (
	"log"
	"strings"
)

// Blekko provides support for Blekko searches.
type Blekko struct{}

// NewBlekko returns a Blekko Searcher value.
func NewBlekko() Searcher {
	return Blekko{}
}

// Search implements the Searcher interface. It performs a search
// against Blekko.
func (b Blekko) Search(searchTerm string, searchResults chan<- []Result) {
	log.Printf("Blekko : Search : Started : searchTerm[%s]\n", searchTerm)

	// Build a proper search url.
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	uri := "http://blekko.com/ws/?q=" + searchTerm + "+%2Frss+%2Fps=8"

	// Perform a RSS search.
	rssSearch("Blekko", uri, searchResults)

	log.Println("Blekko : Search : Completed")
}
