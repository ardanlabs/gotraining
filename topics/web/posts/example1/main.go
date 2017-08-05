// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to handle different HTTP verbs.
package main

import (
	"log"
	"net/http"
)

// App provides a handler to handle GET and POST calls
// for every request.
func App() http.Handler {

	// Declare the handler function to handle the GET and POST call.
	h := func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			GetHandler(res, req)

		case "POST":
			PostHandler(res, req)

		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

	// Store and return the handler function within
	// a http.Handler interface value.
	return http.HandlerFunc(h)
}

// GetHandler provides support for the GET reponse.
func GetHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(`
<form action="/" method="POST">
<input type="submit" value="CLICK ME!!" />
</form>`))
}

// PostHandler provides support for the POST reponse.
func PostHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Thank you for clicking me!"))
}

func main() {
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", App()))
}
