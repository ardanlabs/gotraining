// Copyright Information.

// Package search : search.go manages the searching of results against Google, Yahoo and Bing.
package search

import "log"

// Result represents a search result that was found.
type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

// Searcher declares an interface used to leverage different
// search engines to find results.
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

// searchSession holds information about the current search submission.
// It contains options, searchers and a channel down which we will receive
// results.
type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

// Google search will be added to the search session if this option
// is provided.
func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

// Bing search will be added to this search session if this option
// is provided.
func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

// Yahoo search will be enabled if this option is provided as an argument
// to Submit.
func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

// OnlyFirst is an option that will restrict the search session to just the
// first result.
func OnlyFirst(s *searchSession) { s.first = true }

// Submit uses goroutines and channels to perform a search against the three
// leading search engines concurrently.
func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result

	// Wait for the results to come back.
	for search := 0; search < len(session.searchers); search++ {
		// If we just want the first result, don't wait any longer by
		// concurrently discarding the remaining searchResults.
		// Failing to do so will leave the Searchers blocked forever.
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		result := <-session.resultChan

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		results = append(results, result...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}
