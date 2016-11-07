// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to apply basic authentication with the
// standard library for your web request.
package main

import (
	"log"
	"net/http"
)

// indexHandler handles the root route and validates authentication.
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// Set the following header into the response.
	res.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	// BasicAuth returns the username and password provided in the request's
	// Authorization header, if the request uses HTTP Basic Authentication.
	u, p, ok := req.BasicAuth()
	if !ok {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}

	// If the username and password is not what we expect, then
	// respond with not authorized.
	if u != "username" && p != "password" {
		http.Error(res, "Not authorized", http.StatusUnauthorized)
		return
	}

	// Respond that the user is authorized.
	res.Write([]byte("Welcome Authorized User!"))
}

// App loads the API for use.
func App() http.Handler {

	// Create a mux for binding routes.
	m := http.NewServeMux()

	// Bind the root route.
	m.HandleFunc("/", indexHandler)

	return m
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
