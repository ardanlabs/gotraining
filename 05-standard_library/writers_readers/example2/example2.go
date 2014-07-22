// Sample program to show how a string can be converted to an
// io.Reader and used with functions that support the inteface.
package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

// main is the entry point for the application.
func main() {
	// Convert the string into an io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")

	// Decode the base64 string which returns another io.Reader
	dec := base64.NewDecoder(base64.StdEncoding, buf)

	// Write the decoded string to standard out.
	io.Copy(os.Stdout, dec)

	// Output:
	// Gophers rule!
}
