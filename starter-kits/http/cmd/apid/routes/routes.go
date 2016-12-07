// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package routes

import (
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/handlers"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/app"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/app/middleware"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	a := app.New(middleware.RequestLogger, middleware.Mongo())
	a.Use(middleware.CORS(a, "*", "GET, POST, PUT, PATCH, DELETE, OPTIONS"))

	// Setup the file server to serve up static content such as
	// the index.html page.
	a.TreeMux.NotFoundHandler = http.FileServer(http.Dir("views")).ServeHTTP

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	a.Handle("GET", "/v1/users", handlers.UserList)
	a.Handle("POST", "/v1/users", handlers.UserCreate)
	a.Handle("GET", "/v1/users/:id", handlers.UserRetrieve)
	a.Handle("PUT", "/v1/users/:id", handlers.UserUpdate)
	a.Handle("DELETE", "/v1/users/:id", handlers.UserDelete)

	return a
}
