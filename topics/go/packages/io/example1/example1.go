// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// The io.Reader and io.Writer interfaces allow you to compose all of these different bits together.

// Sample program to show how different functions from the
// standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	// Create a Buffer value and write a string to the buffer.
	// Using the Write method that implements io.Writer.
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to concatenate a string to the Buffer.
	// Passing the address of a bytes.Buffer value for io.Writer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to the stdout device.
	// Passing the address of a os.File value for io.Writer.
	b.WriteTo(os.Stdout)
}
