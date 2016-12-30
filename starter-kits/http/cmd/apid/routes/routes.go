// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package routes

import (
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/handlers"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/middleware"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	app := web.New(middleware.RequestLogger, middleware.ErrorHandler, middleware.Mongo())
	app.Use(middleware.CORS(app, "*", "GET, POST, PUT, PATCH, DELETE, OPTIONS"))

	// Setup the file server to serve up static content such as
	// the index.html page.
	app.TreeMux.NotFoundHandler = http.FileServer(http.Dir("views")).ServeHTTP

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	app.Handle("GET", "/v1/users", handlers.UserList)
	app.Handle("POST", "/v1/users", handlers.UserCreate)
	app.Handle("GET", "/v1/users/:id", handlers.UserRetrieve)
	app.Handle("PUT", "/v1/users/:id", handlers.UserUpdate)
	app.Handle("DELETE", "/v1/users/:id", handlers.UserDelete)

	return app
}
