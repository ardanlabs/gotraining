// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ArdanStudios/gotraining/12-http/api/routes"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("main : Started : Listening on: http://localhost:" + port)
	http.ListenAndServe(":"+port, routes.API())
}
