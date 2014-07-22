/*
http://golang.org/pkg/io/#Pipe
Pipe creates a synchronous in-memory pipe. It can be used to connect code
expecting an io.Reader with code expecting an io.Writer. Reads on one end
are matched with writes on the other, copying data directly between the two;
there is no internal buffering. It is safe to call Read and Write in parallel
with each other or with Close. Close will complete once pending I/O is done.
Parallel calls to Read, and parallel calls to Write, are also safe: the
individual calls will be gated sequentially.

func Pipe() (*PipeReader, *PipeWriter)

*******

http://golang.org/pkg/mime/multipart/#NewWriter
NewWriter returns a new multipart Writer with a random boundary, writing to w.

func NewWriter(w io.Writer) *Writer

*******

http://golang.org/pkg/crypto/sha1/#New
New returns a new hash.Hash computing the SHA1 checksum.

func New() hash.Hash

http://golang.org/pkg/hash/#Hash
Hash is the common interface implemented by all hash functions.

type Hash interface {
        // Write (via the embedded io.Writer interface) adds more data to the running hash.
        // It never returns an error.
        io.Writer

        // Sum appends the current hash to b and returns the resulting slice.
        // It does not change the underlying hash state.
        Sum(b []byte) []byte

        // Reset resets the Hash to its initial state.
        Reset()

        // Size returns the number of bytes Sum will return.
        Size() int

        // BlockSize returns the hash's underlying block size.
        // The Write method must be able to accept any amount
        // of data, but it may operate more efficiently if all writes
        // are a multiple of the block size.
        BlockSize() int
}

*******

http://golang.org/pkg/io/#TeeReader
TeeReader returns a Reader that writes to w what it reads from r. All reads
from r performed through it are matched with corresponding writes to w. There
is no internal buffering - the write must complete before the read completes.
Any error encountered while writing is reported as a read error.

func TeeReader(r Reader, w Writer) Reader

*******

http://golang.org/pkg/mime/multipart/#Writer.CreateFormFile
CreateFormFile is a convenience wrapper around CreatePart. It creates a new
form-data header with the provided field name and file name.

func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

*******

http://golang.org/pkg/io/#Copy
Copy copies from src to dst until either EOF is reached on src or
an error occurs. It returns the number of bytes copied and the first
error encountered while copying, if any.

func Copy(dst Writer, src Reader) (written int64, err error)
*/

// Sample program to show how io.Writes can be embedded within
// other Writer calls to perform complex writes.
package main

import (
	"crypto"
	_ "crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

// main is the entry point for the application.
func main() {
	// Open the file.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Open File", err)
		return
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Create a synchronous in-memory pipe. This pipe will
	// allow us to use a Writer as a Reader.
	pipeReader, pipeWriter := io.Pipe()

	// Create a goroutine to perform the write since
	// the Reader will block until the Writer is closed.
	go func() {
		// Create an SHA1 hash value which implements io.Writer.
		hash := crypto.SHA1.New()

		// Stream the file directly through the hash function to
		// create a hash key.
		io.TeeReader(file, hash)

		// Create the multipart writer to put everything together.
		mpWriter := multipart.NewWriter(pipeWriter)
		fw, err := mpWriter.CreateFormFile("file", "data.json")

		// Write the contents of the file to the multipart form.
		_, err = io.Copy(fw, file)
		if err != nil {
			fmt.Println("Write File", err)
			return
		}

		// Add the SHA hash key we generated.
		mpWriter.WriteField("sha1", hex.EncodeToString(hash.Sum(nil)))

		// Close the Writer which will cause the Reader to unblock.
		mpWriter.Close()
		pipeWriter.Close()
	}()

	// Wait until the Writer is closed, then write the
	// Pipe to Stdout.
	io.Copy(os.Stdout, pipeReader)
}
