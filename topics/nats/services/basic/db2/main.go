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
	getUsers = "req.users" // Subject for all users
	group    = "db2"       // QueueGroup name to avoid duplicate processing
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

	// Function to satisfy nats.Handler interface and pass conn to GetUsers
	f := func(m *nats.Msg) {
		GetUsers(conn, m)
	}

	// QueueSubscribe to getUsers
	sub, err := conn.QueueSubscribe(getUsers, group, f)
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
