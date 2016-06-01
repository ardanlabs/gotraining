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

	// Declare the key to use for publishing/subscribing.
	const key = "test"

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Subscribe to receive messages for the specified key.
	// Passing nil for the handler so everything is a manual pull.
	sub, err := conn.SubscribeSync(key)
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

	// Pull the message we published from the message bus.
	// NOTE: This is a bad call. It creates a timer on each call :(
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
