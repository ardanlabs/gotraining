// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/routes"
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
	server := http.Server{
		Addr:           host,
		Handler:        routes.API(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {

		// Listen for an interrupt signal from the OS.
		osSignals := make(chan os.Signal)
		signal.Notify(osSignals, os.Interrupt)

		<-osSignals

		// Create a context to attempt a graceful 5 second shutdown.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Attempt the graceful shutdown.
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("shutdown : Graceful shutdown did not complete in %v : %v", timeout, err)

			// Looks like we timedout on the graceful shutdown. Kill it hard.
			if err := server.Close(); err != nil {
				log.Printf("shutdown : Error killing server : %v", err)
			}
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Printf("shutdown : %v", err)
	}

	log.Println("main : Completed")
}
