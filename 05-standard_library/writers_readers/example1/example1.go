/*
The io.Reader and io.Writer interfaces allow you to compose all of these different bits together.

http://golang.org/pkg/io
Package io provides basic interfaces to I/O primitives. Its primary job
is to wrap existing implementations of such primitives, such as those in
package os, into shared public interfaces that abstract the functionality,
plus some other related primitives.

http://golang.org/pkg/io/#Reader
type Reader interface {
        Read(p []byte) (n int, err error)
}

http://golang.org/pkg/io/#Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}

http://golang.org/pkg/io/#ReadWriter
type ReadWriter interface {
    Reader
    Writer
}

*******

http://golang.org/pkg/fmt/#Fprintf
Fprint formats using the default formats for its operands and writes to w.
Spaces are added between operands when neither is a string. It returns the
number of bytes written and any write error encountered.

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

*******

http://golang.org/pkg/bytes/#Buffer.WriteTo
WriteTo writes data to w until the buffer is drained or an error occurs.
The return value n is the number of bytes written; it always fits into
an int, but it is int64 to match the io.WriterTo interface. Any error
encountered during the write is also returned.

func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

*******

http://golang.org/pkg/os/#pkg-variables
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
*/

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
