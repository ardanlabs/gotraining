package routes

import (
	"net/http"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/handlers"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	a := app.New()

	a.TreeMux.NotFoundHandler = http.FileServer(http.Dir("views")).ServeHTTP

	a.Handle("GET", "/v1/users", handlers.Users.List)
	a.Handle("POST", "/v1/users", handlers.Users.Create)
	a.Handle("GET", "/v1/users/:id", handlers.Users.Retrieve)
	a.Handle("PUT", "/v1/users/:id", handlers.Users.Update)
	a.Handle("DELETE", "/v1/users/:id", handlers.Users.Delete)

	return a
}
