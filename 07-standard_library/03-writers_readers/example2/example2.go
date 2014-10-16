// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/W3YoitIiT-

// https://gist.github.com/jmoiron/e9f72720cef51862b967#file-01-curl-go
// Sample code provided by Jason Moiron

// ./example2 http://www.goinggo.net/feeds/posts/default

// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main is the entry point for the application.
func main() {
	// r here is a response, and r.Body is an io.Reader
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// io.Copy(dst io.Writer, src io.Reader), copies from the Body to Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
