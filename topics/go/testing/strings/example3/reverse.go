// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package strings provides some useful string manipulation functions.
package strings

// Reverse gives the reversed form of s.
//
// This is the first correct implementation. It may not behave how you want for
// things like unicode combining characters. If you can break it let me know :)
//
// It is still not very efficient but maybe it's good enough?
func Reverse(s string) string {
	in := []rune(s)
	var r string

	for i := len(in) - 1; i >= 0; i-- {
		r = r + string(in[i])
	}

	return r
}
