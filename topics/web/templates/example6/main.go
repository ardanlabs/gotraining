// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to serve up static files from
// a web application and deliver a home page.
package main

import (
	"log"
	"net/http"
	"path"
	"runtime"
)

// App creates a mux and binds the root route for processing
// static files.
func App() http.Handler {

	// Create a new mux for this service.
	m := http.NewServeMux()

	// Bind the route for serving static files using the
	// default FileServer. This will load the home page.
	m.Handle("/", http.FileServer(http.Dir(staticDir())))

	return m
}

// staticDir builds a full path to the 'static' directory
// that is relative to this file.
func staticDir() string {

	// Locate from the runtime the location of
	// the apps static files.
	_, filename, _, _ := runtime.Caller(1)

	// Return a path to the static folder.
	return path.Join(path.Dir(filename), "static")
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
