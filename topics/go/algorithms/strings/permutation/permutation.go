package permutation

import (
	"sort"
)

// RuneSlice a custom type of a slice of runes.
type RuneSlice []rune

// For sorting an RuneSlice.
func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Is check if two strings are permutations.
func Is(str1, str2 string) bool {

	// If the length are not equal they cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	// Convert each string into a collection of runes.
	s1 := []rune(str1)
	s2 := []rune(str2)

	// Sort each collection of runes.
	sort.Sort(RuneSlice(s1))
	sort.Sort(RuneSlice(s2))

	// Convert the collection of runes back to a string
	// and compare.
	return string(s1) == string(s2)
}
