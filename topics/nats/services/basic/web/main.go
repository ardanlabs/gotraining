// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show what a basic web service might look like.
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
	"github.com/nats-io/nats"
)

var rawConn *nats.Conn
var conn *nats.EncodedConn

func main() {
	var err error

	// Connect to the local nats server.
	rawConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("ERROR: Unable to connect to NATS")
		return
	}

	// Create an encoded connection
	conn, err = nats.NewEncodedConn(rawConn, nats.JSON_ENCODER)
	if err != nil {
		log.Println("ERROR: Unable to create an encoded connection")
		return
	}

	// Support for shutting down cleanly.
	go func() {

		// Listen for an interrupt signal from the OS.
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan

		log.Println("Starting shutdown...")
		log.Println("Waiting on requests to complete...")

		// We have been asked to shutdown the server.
		manners.Close()
	}()

	// Bind routes.
	http.HandleFunc("/users", GetUsers)

	// Start the web service.
	const host = "localhost:8080"
	log.Printf("Listening on: %s\n", host)
	manners.ListenAndServe(host, http.DefaultServeMux)

	// Close the connection to the NATS server.
	log.Println("Waiting on NATS to close...")
	conn.Close()
}
