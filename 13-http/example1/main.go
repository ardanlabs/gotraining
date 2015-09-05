// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/vB_ZytmqC1

// Sample program to show how to implement a handler function
// with the http package.
package main

import (
	"fmt"
	"net/http"
)

// HelloHandler shows how the http Handlers are the universal interface for
// handling http stuff, you will see that many Go developers rally around the
// http.Handler interface.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Fprintln is a nice way to write formatted text out to a io.Writer
	fmt.Fprintln(w, "Hello world")
}

func main() {
	// Like many packages in the standard library, the http package has
	// convenience methods that are declared on the package level. These methods
	// actually delegate to an underlying structure that you can initialize
	// yourself. In this example, we are adding a handler to the http.DefaultServeMux
	http.HandleFunc("/", HelloHandler)

	// ListenAndServe will listen on the specified host:port combo. Passing nil
	// as the second argument means we will be using the http.DefaultServeMux as
	// our http handler
	http.ListenAndServe(":4000", nil)
}
