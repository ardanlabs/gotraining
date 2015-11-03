// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/CUnPy2CKqI

// Program to show how to run a basic web server with routing.
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from %q", html.EscapeString(r.URL.Path))
}
