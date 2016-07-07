// Package search manages the searching of results against Google, Yahoo and Bing.
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
	Search(term string, found chan<- []Result)
}

// =============================================================================

// Submit uses goroutines and channels to perform a search against the
// feeds concurrently.
func Submit(options Options) []Result {
	log.Printf("search : Submit : Started : %#v\n", options)

	searchers := make(map[string]Searcher)

	// Create a CNN Searcher if checked.
	if options.CNN {
		log.Println("search : Submit : Info : Adding CNN")
		searchers["cnn"] = NewCNN()
	}

	// Create a NYT Searcher if checked.
	if options.NYT {
		log.Println("search : Submit : Info : Adding NYTimes")
		searchers["nyt"] = NewNYTimes()
	}

	results := make(chan []Result)

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, searcher := range searchers {
		go searcher.Search(options.Term, results)
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
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(found))
			}()
			continue
		}

		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		found := <-results

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(found))
		final = append(final, found...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(final))
	return final
}
