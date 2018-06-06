// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the finder interface for hockey.
// When implementing the find method for hockey, call into the find method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of haystacks and perform the search.
package main

import "strings"

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

// Declare a struct type named hockey that represents specific
// hockey information:
// - Have it embed the sport type first.
// - Have it include a field with the country of the team.

// find checks the value for the specified term.
func ( /* receiver type */ ) find(needle string) bool {

	// Make sure you call into find method for the embedded sport type.

	// Implement the search for the new fields.
	return false
}

func main() {

	// Define the term to find.

	// Create a slice of finder values and assign values
	// of the concrete hockey type.

	// Display what we are trying to find.

	// Range over each finder value and check the term.
}
