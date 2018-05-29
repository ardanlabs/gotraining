// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the finder interface for hockey.
// When implementing the find method for hockey, call into the find method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of finders and perform the search.
package main

import (
	"fmt"
	"strings"
)

// finder defines the behavior required for performing matching.
type finder interface {
	find(needle string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// find checks the value for the specified term.
func (s *sport) find(needle string) bool {
	return strings.Contains(s.team, needle) || strings.Contains(s.city, needle)
}

// hockey represents specific hockey information.
type hockey struct {
	sport
	country string
}

// find checks the value for the specified term.
func (h *hockey) find(needle string) bool {
	return h.sport.find(needle) || strings.Contains(h.country, needle)
}

func main() {

	// Define the term to find.
	needle := "Miami"

	// Create a slice of finder values and assign values
	// of the concrete hockey type.
	haystack := []finder{
		&hockey{sport{"Panthers", "Miami"}, "USA"},
		&hockey{sport{"Canadiens", "Montreal"}, "Canada"},
	}

	// Display what we are trying to find.
	fmt.Println("Searching For:", needle)

	// Range of each haystack value and check the term.
	for _, hs := range haystack {
		if hs.find(needle) {
			fmt.Printf("FOUND: %+v\n", hs)
		}
	}
}
