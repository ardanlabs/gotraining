// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how slices allow for efficient linear traversals.
package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	// Given a stream of bytes to be processed.
	x := []byte{0x0A, 0x15, 0x0e, 0x28, 0x05, 0x96, 0x0b, 0xd0, 0x0}

	// Perform a linear traversal across the bytes, never making
	// copies of the actual data but still passing those bytes
	// to the binary function for processing.

	a := x[0]
	b := binary.LittleEndian.Uint16(x[1:3])
	c := binary.LittleEndian.Uint16(x[3:5])
	d := binary.LittleEndian.Uint32(x[5:9])

	// The result is zero allocation data access that is mechanically
	// sympathetic with the hardware.

	fmt.Println(a, b, c, d)
}
