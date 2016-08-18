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
	var output bytes.Buffer

	fmt.Println("=======================================\nRunning Algorithm One")
	for _, d := range data {
		output.Reset()
		algOne(d.input, &output)
		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}

	fmt.Println("=======================================\nRunning Algorithm Two")
	for _, d := range data {
		output.Reset()
		algTwo(d.input, &output)
		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}
}

// algOne is one way to solve the problem.
func algOne(data []byte, output *bytes.Buffer) {

	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	tmp := make([]byte, 1)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
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
}

// algTwo is a second way to solve the problem.
// Provided by Tyler Bunnell https://twitter.com/TylerJBunnell
func algTwo(data []byte, output *bytes.Buffer) {

	// Use the bytes Reader to provide a stream to process.
	input := bytes.NewReader(data)

	// Create an index variable to match bytes.
	idx := 0

	for {

		// Read a single byte from our input.
		b, err := input.ReadByte()
		if err != nil {
			break
		}

		// Does this byte match the byte at this offset?
		if b == find[idx] {

			// It matches so increment the index position.
			idx++

			// If every byte has been matched, write
			// out the replacement.
			if idx == size {
				output.Write(repl)
				idx = 0
			}

			continue
		}

		// Did we have any sort of match on any given byte?
		if idx != 0 {

			// Write what we've matched up to this point.
			output.Write(find[:idx])

			// Unread the unmatched byte so it can be processed again.
			input.UnreadByte()

			// Reset the offset to start matching from the beginning.
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset.
		output.WriteByte(b)
		idx = 0
	}
}
