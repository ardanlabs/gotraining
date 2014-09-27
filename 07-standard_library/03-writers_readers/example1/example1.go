// http://play.golang.org/p/jlTo1IV1RQ

// The io.Reader and io.Writer interfaces allow you to compose all of these different bits together.

// Sample program to show how different functions from the
// standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main is the entry point for the application.
func main() {
	// Create a Buffer value and write a string to the buffer.
	// Using the io.Writer implementation for Buffer.
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to write the string to the Buffer.
	// Using the io.Writer implementation for Buffer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to the stdout device.
	// Using the io.Writer implementation for os.File.
	b.WriteTo(os.Stdout)
}
