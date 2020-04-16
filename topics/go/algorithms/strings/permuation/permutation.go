package strings

import (
	strings "github.com/ardanlabs/gotraining/topics/go/algorithms/strings/types"
	"sort"
)

// IsPermutation check if two strings are permutations.
func IsPermutation(str1, str2 string) bool {

	// If the length are not equal they cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	// Create a rune for each input string.
	s1 := []rune(str1)
	s2 := []rune(str2)

	// Sort the the two runes
	sort.Sort(strings.RuneSlice(s1))
	sort.Sort(strings.RuneSlice(s2))

	return string(s1) == string(s2)
}
