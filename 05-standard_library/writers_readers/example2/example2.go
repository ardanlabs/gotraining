/*
http://golang.org/pkg/bytes/#NewBufferString
NewBufferString creates and initializes a new Buffer using string
s as its initial contents. It is intended to prepare a buffer to
read an existing string.

func NewBufferString(s string) *Buffer

*******

http://golang.org/pkg/encoding/base64/#NewDecoder
NewDecoder constructs a new base64 stream decoder.

func NewDecoder(enc *Encoding, r io.Reader) io.Reader

*******

http://golang.org/pkg/io/#Copy
Copy copies from src to dst until either EOF is reached on src or
an error occurs. It returns the number of bytes copied and the first
error encountered while copying, if any.

func Copy(dst Writer, src Reader) (written int64, err error)
*/

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
