// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create and use your own mux.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Create a new mux for handling routes.
	m := http.NewServeMux()

	// Bind a handler to the root route.
	f := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World!")
	}
	m.HandleFunc("/", f)

	// Bind a handler to the /foo route.
	f = func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello Foo!")
	}
	m.HandleFunc("/foo", f)

	// Start the http server to handle the request.
	log.Panic(http.ListenAndServe(":3000", m))
}
