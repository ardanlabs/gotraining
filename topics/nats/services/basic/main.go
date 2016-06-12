// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show what a basic web service might look like.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

const version = "1.0.0"

func main() {

	// Parse the command line arguments that will be provided.
	var (
		httpAddr   = flag.String("http", "0.0.0.0:8080", "HTTP service address.")
		healthAddr = flag.String("health", "0.0.0.0:8081", "Health service address.")
	)
	flag.Parse()

	// Log the system is starting.
	log.Println("Starting server...")

	// Create a new server mux and bind the routes for
	// health and readiness.
	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", HealthzHandler)
	hmux.HandleFunc("/readiness", ReadinessHandler)
	hmux.HandleFunc("/healthz/status", HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", ReadinessStatusHandler)

	// Create a manners server binding the specified host
	// and the mux for the routes.
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = LoggingHandler(hmux)

	// Start the server and if an error occurs send it
	// on the channel.
	errHealthChan := make(chan error, 1)
	go func() {
		log.Printf("Health service listening on %s", *healthAddr)
		errHealthChan <- healthServer.ListenAndServe()
	}()

	// Create a new server mux and bind the routes for
	// our service.
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.Handle("/secure", JWTAuthHandler(HelloHandler))
	mux.Handle("/version", VersionHandler(version))

	// Create a manners server binding the specified host
	// and the mux for the routes.
	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = LoggingHandler(mux)

	// Start the server and if an error occurs send it
	// on the channel.
	errServiceChan := make(chan error, 1)
	go func() {
		log.Printf("HTTP service listening on %s", *httpAddr)
		errServiceChan <- httpServer.ListenAndServe()
	}()

	// Bind a channel to receive OS signal events.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Keep the server running until we get an error or
	// we are asked to shutdown by receiving a signal event.
	for {
		select {
		case err := <-errHealthChan:
			log.Fatal("Health:", err)

		case err := <-errServiceChan:
			log.Fatal("Service:", err)

		case <-signalChan:
			log.Println("Shutting down service on request...")
			SetReadinessStatus(http.StatusServiceUnavailable)
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
