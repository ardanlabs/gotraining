package strings

import (
	"sort"
)

type RuneSlice []rune

// For sorting an RuneSlice
func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// IsPermutation check if two strings are permutations.
func IsPermutation(str1, str2 string) bool {

	// If the length are not equal they cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	// Create a rune for each input string.
	s1 := []rune(str1)
	s2 := []rune(str2)

	// Sort the the two runes.
	sort.Sort(RuneSlice(s1))
	sort.Sort(RuneSlice(s2))

	return string(s1) == string(s2)
}
