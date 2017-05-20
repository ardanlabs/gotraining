// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use methods as
package main

import (
	"fmt"
	"log"
	"net/http"
)

// App is an application level context for our service.
type App struct{}

// Default is the default greeting.
func (a App) Default(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

// Foo greets requests at the /foo route.
func (a App) Foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Foo!")
}

// Bar greets requests at the /bar route.
func (a App) Bar(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Bar!")
}

func main() {

	// Create our app
	var a App

	// Use http.HandleFunc instead of http.Handle. Instead of requiring a
	// full-blown type implementing http.Handler this just wants any function
	// that accepts a response writer and request.
	http.HandleFunc("/", a.Default)
	http.HandleFunc("/foo", a.Foo)
	http.HandleFunc("/bar", a.Bar)

	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
