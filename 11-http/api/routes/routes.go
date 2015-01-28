package routes

import (
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/11-http/api/context"
	ctrl "github.com/ArdanStudios/gotraining/11-http/api/controllers"
	"github.com/sqs/mux"
)

// Run binds the service to a port and starts listening for requests.
func Run() {
	log.Println("Listing on: http://localhost:9000")

	http.ListenAndServe(":9000", routes())
}

// routes binds the routes and handlers for the web service.
func routes() *mux.Router {
	r := mux.NewRouter()

	// Custom routes
	context.AddRoute(r, "/search", ctrl.Search)
	return r
}
