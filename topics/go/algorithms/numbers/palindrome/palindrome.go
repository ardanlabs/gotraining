package palindrome

import "github.com/ardanlabs/gotraining/topics/go/algorithms/numbers/reverse"

// Is checks if a integer is a Palindrome.
func Is(input int) bool {

	// A negative integer is not a palindrome.
	if input < 0 {
		return false
	}

	// An integer that is only one digit in length is a palindrome.
	if input >= 0 && input < 10 {
		return true
	}

	// Reverse the digits in the integer.
	rev := reverse.Reverse(input)

	return input == rev
}
