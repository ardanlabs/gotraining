// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to apply middleware.
package main

import (
	"log"
	"net/http"
	"time"
)

// fooHeader returns a handler function that will set the
// `foo` header key then call the provided handler function.
func fooHeader(hf http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Set the `foo` header key.
		res.Header().Set("foo", "bar")

		// Call the handler that was provided.
		hf(res, req)
	}
}

// logger returns a handler function that will log info about
// the request then call the provided handler function.
func logger(hf http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Get the current time.
		start := time.Now()

		// Once the handler call proceeding this defer
		// is complete, log how long the request took.
		defer func() {
			d := time.Now().Sub(start)
			log.Printf("%s %s %s", req.Method, req.URL.Path, d)
		}()

		// Call the handler that was provided.
		hf(res, req)
	}
}

// App loads the API and the middleware.
func App() http.Handler {

	// Create a new mux for handling routes.
	m := http.NewServeMux()

	// For the root route, the logger handler first calls into the
	// fooHeader handler which first calls into the Hello World handler.
	// This chain of calls happen on the processing of the route.
	hello := func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	}
	m.HandleFunc(
		"/",
		logger(
			fooHeader(
				hello,
			),
		),
	)

	return m
}

func main() {

	// Start the http server to handle requests.
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}
