// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show to connect and publish/subscribe requests.
// Message are received asynchronously using a handler function.
package main

import (
	"log"
	"time"

	"github.com/nats-io/nats"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Declare the subject to use for publishing/subscribing.
	const subject = "help"

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Function is called when new requests are received.
	f := func(m *nats.Msg) {
		log.Println("Received a request:", string(m.Data))

		// Send a response to the message.
		if err := conn.Publish(m.Reply, []byte("I can help!")); err != nil {
			log.Println("Publishing a message for specified subject:", err)
			return
		}

		log.Println("Sent the response")
	}

	// Subscribe to receive messages for the specified subject.
	sub, err := conn.Subscribe(subject, f)
	if err != nil {
		log.Println("Subscribing for specified subject:", err)
		return
	}

	// Send a request waiting a second for a response.
	m, err := conn.Request(subject, []byte("help me"), time.Second)
	if err != nil {
		log.Println("Sending a request:", err)
		return
	}

	// Display the response.
	log.Println("Received a response:", string(m.Data))

	// Unsubscribe from receiving these messages.
	if err := sub.Unsubscribe(); err != nil {
		log.Println("Error unsubscribing from the bus:", err)
		return
	}

	// Close the connection to the NATS server.
	conn.Close()
}
