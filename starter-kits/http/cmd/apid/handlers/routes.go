// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"net/http"
	"path"
	"runtime"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/middleware"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
)

// MasterDB contains the master database connections.
var MasterDB *db.DB

// API returns a handler for a set of routes.
func API(masterDB *db.DB) http.Handler {

	// Set the global master database connections.
	MasterDB = masterDB

	// Create the web handler for setting routes and middleware.
	app := web.New(middleware.RequestLogger, middleware.ErrorHandler)

	// Create the file server to serve static content such as
	// the index.html page.
	views := http.FileServer(http.Dir(viewsDir()))
	app.TreeMux.NotFoundHandler = views.ServeHTTP

	// Initialize the routes for the API binding the route to the
	// handler code for each specified verb.
	app.Handle("GET", "/v1/users", UserList)
	app.Handle("POST", "/v1/users", UserCreate)
	app.Handle("GET", "/v1/users/:id", UserRetrieve)
	app.Handle("PUT", "/v1/users/:id", UserUpdate)
	app.Handle("DELETE", "/v1/users/:id", UserDelete)

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
