// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/routes"
	"github.com/braintree/manners"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main : Started")

	// Check the environment for a configured port value.
	host := os.Getenv("HOST")
	if host == "" {
		host = ":3000"
	}

	// Create a new server and set timeout values.
	server := manners.NewWithServer(&http.Server{
		Addr:           host,
		Handler:        routes.API(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	// Create this goroutine to run the web server.
	go func() {
		log.Println("listener : Started : Listening on" + host)
		server.ListenAndServe()
	}()

	// Listen for an interrupt signal from the OS.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("main : Shutting down...")
	server.Close()

	log.Println("main : Completed")
}
