// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ardanlabs/gotraining/13-http/api/routes"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main : Started")

	// Check the environment for a configured port value.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Create this goroutine to run the web server.
	go func() {
		log.Println("listener : Started : Listening on: http://localhost:" + port)
		http.ListenAndServe(":"+port, routes.API())
	}()

	// Listen for an interrupt signal from the OS.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("main : Completed")
}
