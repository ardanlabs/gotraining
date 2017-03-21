// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

/*
	https://blog.golang.org/strings

	Go source code is always UTF-8.
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	Those sequences represent Unicode code points, called runes.
	No guarantee is made in Go that characters in strings are normalized.

	----------------------------------------------------------------------------

	Multiple runes can represent different characters:

	The lower case grave-accented letter à is a character, and it's also a code
	point (U+00E0), but it has other representations.

	We can use the "combining" grave accent code point, U+0300, and attach it to
	the lower case letter a, U+0061, to create the same character à.

	In general, a character may be represented by a number of different sequences
	of code points (runes), and therefore different sequences of UTF-8 bytes.
*/

// Sample program to show how strings have a UTF-8 encoded byte array.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Declare a string with both chinese and english characters.
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	// Iterate over the string.
	for i, r := range s {

		// Capture the number of bytes for this rune.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated
		// with this rune.
		si := i + rl

		// Copy of rune from the string to our buffer.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
