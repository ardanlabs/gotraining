// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to have a single route for
// the api but have access to either through configuration.
package main

import (
	"log"
	"net/http"
)

// defaultAPIVersion sets the default version of
// the api we will server.
const defaultAPIVersion = "2"

// apis is a map of different versions of
// the api we can serve.
var apis map[string]http.Handler

func init() {

	// Initialize the map with the different
	// versions of the api we support.
	apis = map[string]http.Handler{
		"v1": V1(),
		"v2": V2(),
	}
}

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

	// This main handler identifies which version of
	// the api to use and uses it.
	r.HandleFunc("/api/", func(res http.ResponseWriter, req *http.Request) {

		// Look for a specified version.
		v := req.Header.Get("x-version")

		// Retrieve that version and validate it exists. If
		// not use the default version.
		h := apis[v]
		if h == nil {
			h = apis[defaultAPIVersion]
		}

		// Strip the duplication of the path and process
		// the route against the api version.
		http.StripPrefix("/api", h).ServeHTTP(res, req)
	})

	return r
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
