package routes

import (
	"net/http"

	"github.com/ardanlabs/gotraining/13-http/api/app"
	"github.com/ardanlabs/gotraining/13-http/api/handlers"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	a := app.New()

	// Setup the file server to serve up static content such as
	// the index.html page.
	a.TreeMux.NotFoundHandler = http.FileServer(http.Dir("views")).ServeHTTP

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	a.Handle("GET", "/v1/users", handlers.Users.List)
	a.Handle("POST", "/v1/users", handlers.Users.Create)
	a.Handle("GET", "/v1/users/:id", handlers.Users.Retrieve)
	a.Handle("PUT", "/v1/users/:id", handlers.Users.Update)
	a.Handle("DELETE", "/v1/users/:id", handlers.Users.Delete)

	return a
}
