// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This implementation improves performance by removing unnecessary string
// conversions. Further improvements could be made by:
//
// 1. Preallocating enough capacity in out.
// 2. Remove out and sort the input in place.
// 3. Use a specially designed type like strings.Builder from the stdlib.
// 4. Something much more complicated.
//
// As mentioned in example3 this fails on combining characters like "noël". For
// a correct version see: http://rosettacode.org/wiki/Reverse_a_string#Go
func Reverse(s string) string {
	in := []rune(s)
	var out []rune

	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}

	return string(out)
}
