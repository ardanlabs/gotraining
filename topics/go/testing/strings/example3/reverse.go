// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This is the first (mostly) correct implementation. It breaks for strings
// with combining characters; "noël" becomes "l̈eon" instead of "lëon". The
// point of this section is learning testing not reversing strings. For a truly
// correct version see: http://rosettacode.org/wiki/Reverse_a_string#Go
//
// It is still not very efficient but maybe it's good enough?
func Reverse(s string) string {
	var out string

	for _, r := range s {
		out = string(r) + out
	}

	return out
}
