// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This implementation attempts to improve performance by removing unnecessary
// string conversions. Further improvements could be made by:
//
// 1. Preallocating enough capacity in r.
// 2. Remove r and sort the input in place.
// 3. Something much more complicated.
func Reverse(s string) string {
	in := []rune(s)
	var r []rune

	for i := len(in) - 1; i >= 0; i-- {
		r = append(r, in[i])
	}

	return string(r)
}
