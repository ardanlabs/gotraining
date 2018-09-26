// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This is the first real implementation.
// It is not very efficient and it will fail on multibyte input.
func Reverse(s string) string {
	var r string

	for i := len(s) - 1; i >= 0; i-- {
		r = r + string(s[i])
	}

	return r
}
