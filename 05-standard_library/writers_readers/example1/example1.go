// The io.Reader and io.Writer interfaces allow you to compose all of these different bits together.

// Sample program to show how different functions from
// the standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main is the entry point for the application.
func main() {
	// Create a Buffer value and write a string
	// to the buffer.
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to write the string to the Buffer.
	// This is possible because Buffer implements the io.Writer interface.
	fmt.Fprintf(&b, "world!")

	// Write the content of the Buffer to the stdout device.
	// os.File types implement the io.Writer interface.
	b.WriteTo(os.Stdout)
}
