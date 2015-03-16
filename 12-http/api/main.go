// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/handlers"
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
	http.ListenAndServe(":"+port, API())
}

// API returns a handler for a set of routes.
func API() http.Handler {
	a := app.New()
	a.Handle("GET", "/v1/users", handlers.Users.List)
	a.Handle("POST", "/v1/users", handlers.Users.Create)
	a.Handle("GET", "/v1/users/:id", handlers.Users.Retrieve)
	a.Handle("PUT", "/v1/users/:id", handlers.Users.Update)
	a.Handle("DELETE", "/v1/users/:id", handlers.Users.Delete)
	return a
}
