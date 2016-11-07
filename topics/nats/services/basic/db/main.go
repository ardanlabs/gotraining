// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show what a basic web service might look like.
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats"
)

const (
	requestSubject = "dbReq" // Requests come in under this subject.
	group          = "db"    // This is the group this service is listen under.
)

func main() {
	log.Println("Starting service")

	// Connect to the local nats server.
	rawConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("ERROR: Unable to connect to NATS")
		return
	}

	// Create an encoded connection.
	conn, err := nats.NewEncodedConn(rawConn, nats.JSON_ENCODER)
	if err != nil {
		log.Println("ERROR: Unable to create an encoded connection")
		return
	}

	// Function is called when a new message is received.
	f := func(m *nats.Msg) {
		ResponseHandler(conn, m)
	}

	// Subscribe to receive messages for the specified subject and group.
	sub, err := conn.QueueSubscribe(requestSubject, group, f)
	if err != nil {
		log.Println("ERROR: Subscribing for specified subject:", err)
		return
	}

	// Bind a channel to receive OS signal events.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig

	log.Println("Shutting down service")

	// Unsubscribe from receiving messages for this subscription.
	if err := sub.Unsubscribe(); err != nil {
		log.Println("ERROR: Unsubscribing from the bus:", err)
	}

	// Close the connection to the NATS server.
	conn.Close()
}
