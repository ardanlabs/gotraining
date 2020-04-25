package palindrome

import "github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/reverse"

// Is checks if a integer is a Palindrome.
func Is(input int) bool {

	// Integer below zero cannot be palindrome.
	if input < 0 {
		return false
	}

	// A integer is a palindrome if it has the value between 0 and 9.
	if input >= 0 && input < 10 {
		return true
	}

	// Get the reverse integer.
	rev := reverse.Reverse(input)

	return input == rev
}
