// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program that implements a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing04/handlers"
)

// main is the entry point for the application.
func main() {
	Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", handlers.SendJSON)
}
