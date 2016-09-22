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
	"bufio"
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
		input := bytes.NewReader(d.input)
		output.Reset()
		algOne(input, &output)

		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}

	fmt.Println("=======================================\nRunning Algorithm Two")
	for _, d := range data {
		input := bytes.NewReader(d.input)
		output.Reset()
		algTwo(input, &output)

		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}

	fmt.Println("=======================================\nRunning Algorithm Three")
	for _, d := range data {
		input := bytes.NewReader(d.input)
		output.Reset()
		algThree(input, &output)

		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}

	fmt.Println("=======================================\nRunning Algorithm Four")
	for _, d := range data {
		output.Reset()
		output.ReadFrom(NewReplaceReader(bytes.NewReader(d.input), find, repl))

		matched := bytes.Compare(d.output, output.Bytes())
		fmt.Printf("Matched: %v Inp: [%s] Exp: [%s] Got: [%s]\n", matched == 0, d.input, d.output, output.Bytes())
	}
}

// algOne is one way to solve the problem. This approach first
// reads a minimum number of bytes required and then starts processing
// new bytes as they are provided in the stream.
func algOne(r io.Reader, w *bytes.Buffer) {

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	tmp := make([]byte, 1)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(r, buf[:end]); err != nil {
		w.Write(buf[:n])
		return
	}

	for {

		// Read in one byte from the input stream.
		n, err := io.ReadFull(r, tmp)

		// If we have a byte then process it.
		if n == 1 {

			// Add this byte to the end of the buffer.
			buf[end] = tmp[0]

			// If we have a match, replace the bytes.
			if bytes.Compare(buf, find) == 0 {
				copy(buf, repl)
			}

			// Write the front byte since it has been compared.
			w.WriteByte(buf[0])

			// Slice that front byte out.
			copy(buf, buf[1:])
		}

		// Did we hit the end of the stream, then we are done.
		if err != nil {

			// Flush the reset of the bytes we have.
			w.Write(buf[:end])
			break
		}
	}
}

// algTwo is a second way to solve the problem. This approach takes an
// io.Reader to represent an infinite stream. This allows for the algorithm to
// accept input from just about anywhere, thanks to the beauty of Go
// interfaces.
// Provided by Tyler Bunnell https://twitter.com/TylerJBunnell
func algTwo(r io.Reader, w *bytes.Buffer) {

	// Create a byte slice of length 1 into which our byte will be read.
	b := make([]byte, 1)

	// Create an index variable to match bytes.
	idx := 0

	for {

		// Are we re-using the byte from a previous call?
		if b[0] == 0 {

			// Read a single byte from our input.
			n, err := r.Read(b)
			if n == 0 || err != nil {
				break
			}
		}

		// Does this byte match the byte at this offset?
		if b[0] == find[idx] {

			// It matches so increment the index position.
			idx++

			// If every byte has been matched, write
			// out the replacement.
			if idx == size {
				w.Write(repl)
				idx = 0
			}

			// Reset the reader byte to 0 so another byte will be read.
			b[0] = 0
			continue
		}

		// Did we have any sort of match on any given byte?
		if idx != 0 {

			// Write what we've matched up to this point.
			w.Write(find[:idx])

			// NOTE: we are NOT resetting the reader byte to 0 here because we need
			// to re-use it on the next call. This is equivalent to the UnreadByte()
			// call in the other functions.

			// Reset the offset to start matching from the beginning.
			idx = 0

			continue
		}

		// There was no previous match. Write byte and reset.
		w.WriteByte(b[0])

		// Reset the reader byte to 0 so another byte will be read.
		b[0] = 0
	}
}

// algThree is a third way to solve the problem.
// Provided by Bill Hathaway https://twitter.com/billhathaway
func algThree(r io.Reader, w *bytes.Buffer) {

	// This identifies where we are in the match.
	var idx int

	var buf = make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if err != nil || n == 0 {
			break
		}

		// Does newest byte match next byte of find pattern?
		if buf[0] == find[idx] {
			idx++

			// If a full match, write out the replacement pattern.
			if idx == len(find) {
				w.Write(repl)
				idx = 0
			}
			continue
		}

		// If we have matched anything earlier, write it.
		if idx > 0 {
			w.Write(find[:idx])
			idx = 0
		}

		// Match start of pattern?
		if buf[0] == find[0] {
			idx = 1
			continue
		}

		// Write out what we read since no match.
		w.Write(buf)
	}

	// Write out any partial match before returning.
	if idx > 0 {
		w.Write(find[:idx])
	}
}

// algFour is a forth way to solve the problem. This takes a Functional approach.
// Provided by Adam Williams https://twitter.com/acw5
func algFour(r io.Reader, w *bytes.Buffer) {
	buf := bufio.NewReaderSize(r, len(find))

	for {
		peek, err := buf.Peek(len(find))
		if err == nil && bytes.Equal(find, peek) {

			// A match was found. Advance the bufio reader past the match.
			if _, err := buf.Discard(len(find)); err != nil {
				return
			}
			w.Write(repl)
			continue
		}

		// Ignore any peek errors because we may not be able to peek
		// but still be able to read a byte.

		c, err := buf.ReadByte()
		if err != nil {
			return
		}
		w.WriteByte(c)
	}
}

// NewReplaceReader returns an io.Reader that reads from r, replacing
// any occurrence of old with new. Used by algFour.
func NewReplaceReader(r io.Reader, old, new []byte) io.Reader {
	return &replaceReader{
		br:  bufio.NewReaderSize(r, len(old)),
		old: old,
		new: new,
	}
}

type replaceReader struct {
	br       *bufio.Reader
	old, new []byte
}

// Read reads into p the translated bytes.
func (r *replaceReader) Read(p []byte) (int, error) {
	var n int
	for {
		peek, err := r.br.Peek(len(r.old))
		if err == nil && bytes.Equal(r.old, peek) {

			// A match was found. Advance the bufio reader past the match.
			if _, err := r.br.Discard(len(r.old)); err != nil {
				return n, err
			}
			copy(p[n:], r.new)
			n += len(r.new)
			continue
		}

		// Ignore any peek errors because we may not be able to peek
		// but still be able to read a byte.

		p[n], err = r.br.ReadByte()
		if err != nil {
			return n, err
		}
		n++

		// Read reads up to len(p) bytes into p. Since we could potentially add
		// len(r.new) new bytes, we check here that p still has capacity.
		if n+len(r.new) >= len(p) {
			return n, nil
		}
	}
}
