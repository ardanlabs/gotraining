// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show to connect and publish/subscribe for messages.
// Message are received on demand using NextMsg.
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
	const subject = "test"

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Subscribe to receive messages for the specified subject.
	sub, err := conn.SubscribeSync(subject)
	if err != nil {
		log.Println("Subscribing for specified subject:", err)
		return
	}

	// Publish the message for the specified subject.
	if err := conn.Publish(subject, []byte("Hello World")); err != nil {
		log.Println("Publishing a message for specified subject:", err)
		return
	}

	// Pull the message we just published.
	m, err := sub.NextMsg(time.Second)
	if err != nil {
		log.Println("Error pulling a message from the bus:", err)
		return
	}

	// Display the message.
	log.Println("Received a message:", string(m.Data))

	// Unsubscribe from receiving these messages.
	if err := sub.Unsubscribe(); err != nil {
		log.Println("Error unsubscribing from the bus:", err)
		return
	}

	// Close the connection to the NATS server.
	conn.Close()
}
