// Package service maintains the logic for the web service.
package service

import (
	"net/http"
)

// init binds the routes and handlers for the web service.
func init() {
	// Setup a route for our static files.
	//
	// Because our static directory is set as the root of the FileSystem,
	// we need to strip off the /static/ prefix from the request path
	// before searching the FileSystem for the given file.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup a route for our search.
	http.HandleFunc("/search", index)
}

// Run binds the service to a port and starts listening
// for requests.
func Run() {
	// Listen for our HTTP requests.
	http.ListenAndServe("localhost:9999", nil)
}
