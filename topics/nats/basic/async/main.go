// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show to connect and publish/subscribe for messages.
// Message are received asynchronously using a handler function.
package main

import (
	"log"
	"sync"

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

	// Used to wait for the message to be received.
	var wg sync.WaitGroup
	wg.Add(1)

	// Function is called when new messages are received.
	f := func(m *nats.Msg) {
		log.Println("Received a message:", string(m.Data))
		wg.Done()
	}

	// Subscribe to receive messages for the specified subject.
	sub, err := conn.Subscribe(subject, f)
	if err != nil {
		log.Println("Subscribing for specified subject:", err)
		return
	}

	// Publish the message for the specified subject.
	if err := conn.Publish(subject, []byte("Hello World")); err != nil {
		log.Println("Publishing a message for specified subject:", err)
		return
	}

	// Wait to be told the message was received.
	wg.Wait()

	// Unsubscribe from receiving these messages.
	if err := sub.Unsubscribe(); err != nil {
		log.Println("Error unsubscribing from the bus:", err)
		return
	}

	// Close the connection to the NATS server.
	conn.Close()
}
