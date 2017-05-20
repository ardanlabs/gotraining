// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create handlers for different routes
// utilizing the default mux.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// App is the default Handler for our service.
type App struct{}

// ServeHTTP implements the http.Handler interface.
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

// FooApp handles greeting requests under the /foo route.
type FooApp struct{}

// ServeHTTP implements the http.Handler interface.
func (a FooApp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Foo!")
}

// BarApp handles greeting requests under the /bar route.
type BarApp struct{}

// ServeHTTP implements the http.Handler interface.
func (a BarApp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Bar!")
}

func main() {

	// Use http.Handle to register our handlers. It is a shortcut for
	// the Handle method on http.DefaultServeMux
	http.Handle("/", App{})
	http.Handle("/foo", FooApp{})
	http.Handle("/bar", BarApp{})

	// Start the server passing nil for the Handler. This tells
	// ListenAndServe to use http.DefaultServeMux
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
