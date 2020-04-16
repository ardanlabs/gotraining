package strings

import (
	"sort"
	"strings"
)

// IsPermutation check if two
// strings are permutations.
func IsPermutation(str1, str2 string) bool {

	// If the length are not equal they cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	// Create a slice for each input string that
	//is split on every character of the string.
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")

	// Sort the the two slices.
	sort.Strings(s1)
	sort.Strings(s2)

	return strings.Join(s1, "") == strings.Join(s2, "")
}
