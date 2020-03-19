package strings

import strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"

// IsPalindrome checks if a string is a Palindrome.
func IsPalindrome(input string) bool {

	// If the input string is empty or as a length of 1 return true.
	if input == "" || len(input) == 1 {
		return true
	}

	// Create a reverse string from input string.
	rev := strings.ReverseString(input)

	// Check if input and rev strings are equal.
	if input == rev {
		return true
	}

	return false
}
