// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/trace"
)

// LoadWrite reads a file from the network into memory and then
// writes it to disk.
func LoadWrite() {

	// Download the tar file.
	r, err := http.Get("https://ftp.gnu.org/gnu/binutils/binutils-2.7.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Read in the entire contents of the file.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	// Create a new file.
	f, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	// Defer the close and removal of the file.
	defer os.Remove(f.Name())
	defer f.Close()

	// Write the data to the file.
	_, err = f.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

// StreamWrite streams a file from the network, writing it to disk.
func StreamWrite() {

	// Download the tar file.
	r, err := http.Get("https://ftp.gnu.org/gnu/binutils/binutils-2.7.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new file.
	f, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	// Defer the close and removal of the file.
	defer os.Remove(f.Name())
	defer f.Close()

	// Stream the file to disk.
	if _, err = io.Copy(f, r.Body); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
}

func main() {

	// Create a file to hold tracing data.
	tf, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()

	// Start gathering the tracing data.
	trace.Start(tf)
	defer trace.Stop()

	// Perform the work.
	//LoadWrite()
	//StreamWrite()
}
