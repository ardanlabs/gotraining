// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create a simple web service.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Write a handler function to accept and process the request.
	f := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World!")
	}

	// Bind the handler function to the root route.
	http.HandleFunc("/", f)

	// Start the http server to handle the request.
	log.Panic(http.ListenAndServe(":3000", nil))
}
