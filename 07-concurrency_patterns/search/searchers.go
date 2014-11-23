// Copyright Information.

// Package search : seachers.go contains all the different implementations
// for the existing searchers.
package search

import (
	"log"
	"math/rand"
	"time"
)

// init is called before main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Google provides support for Google searches.
type google struct{}

// Search implements the Searcher interface. It performs a search
// against Google.
func (g google) Search(term string, results chan<- []Result) {
	log.Printf("Google : Search : Started : search term[%s]\n", term)

	// Slice for the results.
	var r []Result

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	r = append(r, Result{
		Engine:      "Google",
		Title:       "The Go Programming Language",
		Description: "The Go Programming Language",
		Link:        "https://golang.org/",
	})

	log.Printf("Google : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

// Bing provides support for Bing searches.
type bing struct{}

// Search implements the Searcher interface. It performs a search
// against Bing.
func (b bing) Search(term string, results chan<- []Result) {
	log.Printf("Bing : Search : Started : search term [%s]\n", term)

	// Slice for the results.
	var r []Result

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	r = append(r, Result{
		Engine:      "Bing",
		Title:       "A Tour of Go",
		Description: "Welcome to a tour of the Go programming language.",
		Link:        "http://tour.golang.org/",
	})

	log.Printf("Bing : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

// Yahoo provides support for Yahoo searches.
type yahoo struct{}

// Search implements the Searcher interface. It performs a search
// against Yahoo.
func (y yahoo) Search(term string, results chan<- []Result) {
	log.Printf("Yahoo : Search : Started : search term [%s]\n", term)

	// Slice for the results.
	var r []Result

	// Simulate an amount of time for the search.
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// Simulate a result for the search.
	r = append(r, Result{
		Engine:      "Yahoo",
		Title:       "Go Playground",
		Description: "The Go Playground is a web service that runs on golang.org's servers",
		Link:        "http://play.golang.org/",
	})

	log.Printf("Yahoo : Search : Completed : Found[%d]\n", len(r))
	results <- r
}
