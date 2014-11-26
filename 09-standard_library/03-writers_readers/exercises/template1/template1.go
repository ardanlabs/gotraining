// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/6JAjLK8Tpw

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
	resp_variable_name, error_variable_name := http.Get("http://www.goinggo.net/feeds/posts/default")
	if error_variable_name != nil {
		fmt.Println(error_variable_name)
		return
	}

	// A slice of io.Writers we will write the file to.
	var slice_name []io.writer_interface_type

	// Send the document to stdout.
	slice_name = append(slice_name, os.stdout_variable)

	// Send the document to a file.
	variable_name, error_variable_name := os.create_function("goinggo.rss")
	if error_variable_name != nil {
		fmt.Println(error_variable_name)
		return
	}
	slice_name = append(slice_name, variable_name)
	defer variable_name.close_function()

	// MultiWriter(io.Writer...) returns a single writer which multiplexes its
	// writes across all of the writers we pass in.
	multi_writer_variable := io.MultiWriter(slice_name...)

	// Write to dest the same way as before, copying from the Body.
	io.Copy(multi_writer_variable, resp_variable_name.Body)
	if error_variable_name := resp_variable_name.Body.Close(); error_variable_name != nil {
		fmt.Println(error_variable_name)
	}
}
