// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the match method for hockey, call into the match method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import "strings"

// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// match checks the value for the specified term.
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.

// match checks the value for the specified term.
func ( /* receiver type */ ) match(searchTerm string) bool {

	// Make sure you call into match method for the embedded sport type.

	// Implement the search for the new fields.
	return false
}

func main() {

	// Define the term to match.

	// Create a slice of matcher values and assign values
	// of the concrete hockey type.

	// Display what we are trying to match.

	// Range of each matcher value and check the term.
}
