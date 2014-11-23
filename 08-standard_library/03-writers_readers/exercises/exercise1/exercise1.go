// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/3s-weLqNZC

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// main is the entry point for the application.
func main() {
	// r here is a response, and r.Body is an io.Reader
	r, err := http.Get("http://www.goinggo.net/feeds/posts/default")
	if err != nil {
		fmt.Println(err)
		return
	}

	// A slice of io.Writers we will write the file to.
	var writers []io.Writer

	// Send the document to stdout.
	writers = append(writers, os.Stdout)

	// Send the document to a file.
	file, err := os.Create("goinggo.rss")
	if err != nil {
		fmt.Println(err)
		return
	}
	writers = append(writers, file)
	defer file.Close()

	// MultiWriter(io.Writer...) returns a single writer which multiplexes its
	// writes across all of the writers we pass in.
	dest := io.MultiWriter(writers...)

	// Write to dest the same way as before, copying from the Body.
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
