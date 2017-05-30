// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to apply middleware using negroni.
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
)

// App loads the API and the middleware.
func App() http.Handler {

	// Create a middleware chain of useful gorilla handlers
	chain := alice.New(
		logger,                             // An adapter to use the gorilla logger
		fooHeader,                          // A custom middleware
		handlers.RecoveryHandler(),         // Serve 500 status on panics
		handlers.CompressHandler,           // Gzip responses
		handlers.ProxyHeaders,              // Interpret headers from reverse proxy like nginx
		handlers.HTTPMethodOverrideHandler, // Support PUT/DELETE for some older clients
		handlers.CORS(),                    // Allow cross origin requests
	)

	// A second chain that is only used for certain "secured" routes
	secure := alice.New(decoderCheck)

	// Create a mux and bind our handlers to routes. The default route
	// does not use any special middleware. The /secret route uses the
	// additional "secure" middleware chain.
	m := http.NewServeMux()
	m.HandleFunc("/", helloWorld)
	m.Handle("/secret", secure.ThenFunc(secret))

	// Wrap the chain around our mux. All requests use this chain first.
	return chain.Then(m)
}

func main() {
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}
