// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/6CAkumo0HI

// Sample program to show how strings have a UTF-8 encoded byte array.
package main

import (
	"fmt"
	"unicode/utf8"
)

// main is the entry point for the application.
func main() {
	// Declare a string with both chinese and english charaters.
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	// Iterate over each character in the string.
	for i, r := range s {
		// Capture the number of bytes for this character.
		k := utf8.RuneLen(r)

		// Calculate the slice offset to slice the character out.
		j := i + k

		// Copy of character from the string to our buffer.
		copy(buf[:], s[i:j])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:k])
	}
}
