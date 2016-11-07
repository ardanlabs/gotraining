// Package search manages the searching of results against different
// news feeds.
package search

import (
	"html/template"
	"log"
)

// Options provides the search options for performing searches.
type Options struct {
	Term  string
	CNN   bool
	NYT   bool
	BBC   bool
	First bool
}

// Result represents a search result that was found.
type Result struct {
	Engine  string
	Title   string
	Link    string
	Content string
}

// TitleHTML fixes encoding issues.
func (r *Result) TitleHTML() template.HTML {
	return template.HTML(r.Title)
}

// ContentHTML fixes encoding issues.
func (r *Result) ContentHTML() template.HTML {
	return template.HTML(r.Content)
}

// =============================================================================

// Searcher declares an interface used to leverage different
// search engines to find results.
type Searcher interface {
	Search(uid string, term string, found chan<- []Result)
}

// =============================================================================

// Submit uses goroutines and channels to perform a search against the
// feeds concurrently.
func Submit(uid string, options Options) []Result {
	log.Printf("%s : Submit : Started : %#v\n", uid, options)

	searchers := make(map[string]Searcher)

	// Create a CNN Searcher if checked.
	if options.CNN {
		log.Printf("%s : Submit : Info : Adding CNN\n", uid)
		searchers["cnn"] = NewCNN()
	}

	// Create a NYT Searcher if checked.
	if options.NYT {
		log.Printf("%s : Submit : Info : Adding NYT\n", uid)
		searchers["nyt"] = NewNYT()
	}

	// Create a BBC Searcher if checked.
	if options.BBC {
		log.Printf("%s : Submit : Info : Adding BBC\n", uid)
		searchers["bbc"] = NewBBC()
	}

	results := make(chan []Result)

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, searcher := range searchers {
		go searcher.Search(uid, options.Term, results)
	}

	var final []Result

	// Wait for the results to come back.
	for search := 0; search < len(searchers); search++ {

		// If we just want the first result, don't wait any longer by
		// concurrently discarding the remaining results.
		// Failing to do so will leave the Searchers blocked forever.
		if options.First && (search > 0 && len(final) > 0) {
			go func() {
				found := <-results
				log.Printf("%s : Submit : Info : Results Discarded : Results[%d]\n", uid, len(found))
			}()
			continue
		}

		// Wait to recieve results.
		log.Printf("%s : Submit : Info : Waiting For Results...\n", uid)
		found := <-results

		// Save the results to the final slice.
		log.Printf("%s : Submit : Info : Results Used : Results[%d]\n", uid, len(found))
		final = append(final, found...)
	}

	log.Printf("%s : Submit : Completed : Found [%d] Results\n", uid, len(final))
	return final
}
