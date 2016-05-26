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

	// Declare the key to use for publishing/subscribing.
	const key = "test"

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
	// NOTE: This is bad. Not receiving a value of type Msg.
	f := func(m *nats.Msg) {
		log.Println("Received a message:", string(m.Data))
		wg.Done()
	}

	// Subscribe to receive messages for the specified key.
	sub, err := conn.Subscribe(key, f)
	if err != nil {
		log.Println("Subscribing for specified key:", err)
		return
	}

	// Publish the string into the message bus under the
	// specified key.
	if err := conn.Publish(key, []byte("Hello World")); err != nil {
		log.Println("Publishing a message for specified key:", err)
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
