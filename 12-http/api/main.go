// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/handlers"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main : Started : Listening on: http://localhost:9000")
	http.ListenAndServe(":9000", Api())
}

func Api() http.Handler {
	a := app.New()
	a.Handle("GET", "/v1/users", handlers.UsersList)
	a.Handle("POST", "/v1/users", handlers.UsersCreate)
	a.Handle("GET", "/v1/users/:id", handlers.UsersRetrieve)
	a.Handle("PUT", "/v1/users/:id", handlers.UsersUpdate)
	a.Handle("DELETE", "/v1/users/:id", handlers.UsersDelete)
	return a
}
