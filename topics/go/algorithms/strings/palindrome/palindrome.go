package palindrome

import "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"

// Is checks if a string is a Palindrome.
func Is(input string) bool {

	// If the input string is empty or as a length of 1 return true.
	if input == "" || len(input) == 1 {
		return true
	}

	// Create a reverse string from input string.
	rev := reverse.String(input)

	// Check if input and rev strings are equal.
	if input == rev {
		return true
	}

	return false
}
