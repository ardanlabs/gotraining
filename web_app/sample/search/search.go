// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package search manages the searching of results against Google, Blekko and Bing.
package search

import (
	"html/template"
	"log"
)

// Options provides the search options for performing searches.
type Options struct {
	SearchTerm string
	Google     bool
	Bing       bool
	Blekko     bool
	First      bool
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

// Searcher declares an interface used to leverage different
// search engines to find results.
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

// Submit uses goroutines and channels to perform a search against the three
// leading search engines concurrently.
func Submit(options *Options) []Result {
	log.Printf("search : Submit : Started : %#v\n", options)

	var final []Result
	searchers := make(map[string]Searcher)
	searchResults := make(chan []Result)

	// Create a Google Searcher if checked.
	if options.Google {
		log.Println("search : Submit : Info : Adding Google")
		searchers["google"] = NewGoogle()
	}

	// Create a Bing Searcher if checked.
	if options.Bing {
		log.Println("search : Submit : Info : Adding Bing")
		searchers["bing"] = NewBing()
	}

	// Create a Bing Searcher if checked.
	if options.Blekko {
		log.Println("search : Submit : Info : Adding Blekko")
		searchers["blekko"] = NewBlekko()
	}

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, searcher := range searchers {
		go searcher.Search(options.SearchTerm, searchResults)
	}

	// Wait for the results to come back.
	for search := 0; search < len(searchers); search++ {
		// If we just want the first result, don't wait any longer by
		// concurrently discarding the remaining searchResults.
		// Failing to do so will leave the Searchers blocked forever.
		if options.First && search > 0 {
			go func() {
				results := <-searchResults
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(results))
			}()
			continue
		}

		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		results := <-searchResults

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(results))
		final = append(final, results...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(final))
	return final
}
