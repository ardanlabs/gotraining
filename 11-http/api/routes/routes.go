// Package routes maintains the start and routes logic.
package routes

import (
	"log"
	"net/http"

	ctrl "github.com/ArdanStudios/gotraining/11-http/api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// routes binds the routes and handlers for the web service.
func routes() *mux.Router {
	r := mux.NewRouter()

	// 404 processing
	r.NotFoundHandler = http.HandlerFunc(ctrl.NotFound)

	// Custom routes
	r.HandleFunc("/search", ctrl.Search)
	return r
}

// Run binds the service to a port and starts listening for requests.
func Run() {
	log.Println("Listing on: http://localhost:9000")

	n := negroni.New(ctrl.Authentication{}, ctrl.BeforeRequest{})
	n.UseHandler(routes())
	n.Use(ctrl.AfterRequest{})

	http.ListenAndServe(":9000", n)
}
