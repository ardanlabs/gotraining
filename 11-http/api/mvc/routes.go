package mvc

import (
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// Run binds the service to a port and starts listening for requests.
func Run() {
	log.Println("main : mvc : Run : Started : Listing on: http://localhost:9000")

	http.ListenAndServe(":9000", routes())
}

// routes binds the routes and handlers for the web service.
func routes() *httptreemux.TreeMux {
	r := httptreemux.New()

	// Custom user routes
	AddRoute(r, "/user", GetUser, "GET")
	AddRoute(r, "/user", InsUser, "POST")
	AddRoute(r, "/user", UpdUser, "PUT")
	AddRoute(r, "/user", DelUser, "DELETE")

	return r
}
