// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package routes

import (
	"net/http"
	"path"
	"runtime"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/routes/handlers"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/middleware"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	app := web.New(middleware.RequestLogger, middleware.ErrorHandler, middleware.Mongo())

	// Create the file server to serve static content such as
	// the index.html page.
	views := http.FileServer(http.Dir(viewsDir()))
	app.TreeMux.NotFoundHandler = views.ServeHTTP

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	app.Handle("GET", "/v1/users", handlers.UserList)
	app.Handle("POST", "/v1/users", handlers.UserCreate)
	app.Handle("GET", "/v1/users/:id", handlers.UserRetrieve)
	app.Handle("PUT", "/v1/users/:id", handlers.UserUpdate)
	app.Handle("DELETE", "/v1/users/:id", handlers.UserDelete)

	return app
}

// viewsDir builds a full path to the 'views' directory
// that is relative to this file. It uses a trick of the
// runtime package to get the path of the file that calls
// this function.
func viewsDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "../views")
}
