// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// To bundle assets into the source code so the binary
// has everything it needs:
// $ rice embed-go
// $ go build

// Sample program to show how to bundle assets, static files, etc
// into web application and access these bundled resources.
package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

// App creates a mux and binds the routes for use in our server.
func App() http.Handler {

	// Create a new mux for this service.
	m := http.NewServeMux()

	// Create a rice box for our static folder. These assest
	// are now cached into memory.
	box := rice.MustFindBox("./static")

	// Bind the rice box into the http FileServer and create a
	// route for these assests under an imaginary assets folder.
	assets := http.StripPrefix("/assets/", http.FileServer(box.HTTPBox()))
	m.Handle("/assets/", assets)

	// Bind the root handler to serve up the home page.
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		b, _ := box.Bytes("index.html")
		res.Write(b)
	})

	return m
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
