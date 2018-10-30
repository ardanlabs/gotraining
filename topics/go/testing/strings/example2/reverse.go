// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This is the first real attempt. It fails on multibyte input.
func Reverse(s string) string {
	var out string

	for i := 0; i < len(s); i++ {
		out = string(s[i]) + out
	}

	return out
}
