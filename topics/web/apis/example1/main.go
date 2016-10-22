// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create a simple web api with
// different versions.
package main

import (
	"log"
	"net/http"
)

// V1 contains the version 1 handlers for our API.
func V1() http.Handler {

	// Create a new mux for this version.
	r := http.NewServeMux()

	// Bind the handler function for the user API.
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("v1"))
	})

	return r
}

// V2 contains the version 2 handlers for our API.
func V2() http.Handler {

	// Create a new mux for this version.
	r := http.NewServeMux()

	// Bind the handler function for the user API.
	r.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("v2"))
	})

	return r
}

// App loads the entire API set together for use.
func App() http.Handler {

	// Create a new mux which will process the
	// initial requests.
	r := http.NewServeMux()

	// Load the version 1 routes striping the duplication
	// of the resulting path.
	r.Handle("/api/v1/", http.StripPrefix("/api/v1", V1()))

	// Load the version 2 routes striping the duplication
	// of the resulting path.
	r.Handle("/api/v2/", http.StripPrefix("/api/v2", V2()))

	return r
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
