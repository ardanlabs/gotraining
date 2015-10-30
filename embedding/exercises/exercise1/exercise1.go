// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/B7VVYyaA21

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the Search method for hockey, call into the Search method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing searches.
type matcher interface {
	search(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// Search checks the value for the specified term.
func (s sport) search(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// hockey represents specific hockey information.
type hockey struct {
	sport
	country string
}

// Search checks the value for the specified term.
func (h hockey) search(searchTerm string) bool {
	return h.sport.search(searchTerm) || strings.Contains(h.country, searchTerm)
}

// main is the entry point for the application.
func main() {
	// Define the term to search.
	searchTerm := "Miami"

	// Create a slice of matcher values to search.
	matchers := []matcher{
		hockey{sport{"Panthers", "Miami"}, "USA"},
		hockey{sport{"Canadiens", "Montreal"}, "Canada"},
	}

	// Display what we are searching for.
	fmt.Println("Searching For:", searchTerm)

	// Range of each matcher value and check the search term.
	for _, m := range matchers {
		if m.search(searchTerm) {
			fmt.Printf("FOUND: %+v", m)
		}
	}
}
