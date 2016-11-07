// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement your own App Handler
// that can use any provided handler function.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// App provides application level context for our handler.
type App struct {
	h http.HandlerFunc
}

// ServeHTTP implements the http.Handler interface.
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// Execute the handler that was configured for
	// this custom App handler.
	a.h(res, req)
}

// myHandler handles the implementation of the request.
func myHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

// wrap takes a handler function and configures the custom
// App handler to use it.
func wrap(h http.HandlerFunc) http.Handler {
	return App{h: h}
}

func main() {

	// Create a new mux for handling routes.
	m := http.NewServeMux()

	// Bind a new App handler to the root route using
	// the provided handler function to process requests.
	m.Handle("/", wrap(myHandler))

	// Start the http server to handle the request.
	log.Panic(http.ListenAndServe(":3000", m))
}
