// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that takes a stream of bytes and looks for the bytes
// “elvis” and when they are found, replace them with “Elvis”. The code
// cannot assume that there are any line feeds or other delimiters in the
// stream and the code must assume that the stream is of any arbitrary length.
// The solution cannot meaningfully buffer to the end of the stream and
// then process the replacement.
package main

import (
	"bytes"
	"fmt"
	"io"
)

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// Declare what needs to be found and its replacement.
var find = []byte("elvis")
var repl = []byte("Elvis")

// Calculate the number of bytes we need to locate.
var size = len(find)

func main() {

	// Range over the table testing the algorithm against each input/output.
	for _, d := range data {

		// Use the bytes package to provide a stream to process.
		input := bytes.NewBuffer(d.input)
		var output bytes.Buffer

		// Declare the buffers we need to process the stream.
		buf := make([]byte, size)
		tmp := make([]byte, 1)
		end := size - 1

		// Read in an initial number of bytes we need to get started.
		if n, err := io.ReadFull(input, buf[:end]); err != nil {
			fmt.Printf("Match: 0 Inp: [%s] Exp: [%s] Got: [%s]\n", d.input, d.output, buf[:n])
			continue
		}

		for {

			// Read in one byte from the input stream.
			n, err := io.ReadFull(input, tmp)

			// If we have a byte then process it.
			if n == 1 {

				// Add this byte to the end of the buffer.
				buf[end] = tmp[0]

				// If we have a match, replace the bytes.
				if bytes.Compare(buf, find) == 0 {
					copy(buf, repl)
				}

				// Write the front byte since it has been compared.
				output.WriteByte(buf[0])

				// Slice that front byte out.
				copy(buf, buf[1:])
			}

			// Did we hit the end of the stream, then we are done.
			if err != nil {

				// Flush the reset of the bytes we have.
				output.Write(buf[:end])
				break
			}
		}

		// Create strings from the bytes.
		matched := bytes.Compare(d.output, output.Bytes())

		// Display the results.
		fmt.Printf("Match: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched, d.input, d.output, output.Bytes())
	}
}
