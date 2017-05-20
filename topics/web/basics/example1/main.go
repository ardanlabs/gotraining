// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create a handler for a basic web app.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// App provides application level context for our handler.
type App struct{}

// ServeHTTP implements the http.Handler interface. It gives the same
// greeting to every request
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

func main() {

	// Create a value of our App. It implements http.Handler so we can
	// pass it to http.ListenAndServe
	var a App

	// Log that the server is starting so we see it's alive.
	log.Print("Listening on localhost:3000. Ctrl-c to cancel.")

	// Start the http server to handle the request.
	log.Fatal(http.ListenAndServe("localhost:3000", a))

	// You provide the host:port to bind to. Here we are binding to port
	// 3000 and only on the local network by specifying
	// "localhost:3000". If you want this service to hear connections
	// from external machines you can omit the host and just do ":3000".
}
