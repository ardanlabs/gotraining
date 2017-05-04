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
	// You provide the host:port to bind to. Here we are binding to port 300 only
	// on the local network by specifying "localhost:3000". If you want this
	// service to hear connections from external machines you can omit the host
	// and just do ":3000".
	// It's also a good idea to log that the server is starting.
	log.Print("Listening on localhost:3000")
	log.Panic(http.ListenAndServe("localhost:3000", nil))
}
