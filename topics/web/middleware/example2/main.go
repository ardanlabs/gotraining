// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to apply middleware using negroni.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

// App loads the API and the middleware.
func App() http.Handler {

	// Create a new mux for handling routes.
	m := http.NewServeMux()

	// Bind the root route to the Hello World response.
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	// Create a new negroni to apply middleware.
	n := negroni.New()

	// Add the two middleware handlers.
	n.UseFunc(logger)
	n.UseFunc(fooHeader)

	// Apply the mux to negroni.
	n.UseHandler(m)

	return n
}

// fooHeader returns a handler function that will set
// the `foo` header key. Then call the provided
// handler function.
func fooHeader(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	// Set the `foo` header key.
	res.Header().Set("foo", "bar")

	// Call the handler that was provided.
	next(res, req)
}

// logger returns a handler function that will log info about
// the request. Then it calls the provided handler function.
func logger(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	// Get the current time.
	start := time.Now()

	// Once the handler call proceeding this defer
	// is complete, log how long the request took.
	defer func() {
		d := time.Now().Sub(start)
		log.Printf("%s %s %s", req.Method, req.URL.Path, d)
	}()

	// Call the handler that was provided.
	next(res, req)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
