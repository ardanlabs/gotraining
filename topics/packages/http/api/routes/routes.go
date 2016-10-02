// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package routes

import (
	"net/http"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/handlers"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/middleware"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	a := app.New(middleware.RequestLogger, middleware.Mongo)

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
