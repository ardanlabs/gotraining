// http://play.golang.org/p/Tmt7v3fIQF

// https://github.com/extemporalgenome/watchpost/blob/master/main.go
// Sample code provided by Kevin Gillette

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

		// Create a Reader that writes to the hash what it reads
		// from the file.
		hashReader := io.TeeReader(file, hash)

		// Create the multipart writer to put everything together.
		mpWriter := multipart.NewWriter(pipeWriter)
		fileWriter, err := mpWriter.CreateFormFile("file", "data.json")

		// Write the contents of the file to the multipart form.
		_, err = io.Copy(fileWriter, hashReader)
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
