// This program provides a sample web service implements a RESTFul CRUD
// related API against a MongoDB database.
package main

import (
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/11-http/api/routes"
)

// init is called before main. We are using init to
// set the logging package.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main : Started : Listing on: http://localhost:9000")

	http.ListenAndServe(":9000", routes.TM)
}
