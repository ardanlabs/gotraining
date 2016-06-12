// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show to using the queuing functionality to allow
// a round robin of services to handle messages.
package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Declare the subject and group to use for publishing/subscribing.
	const subject = "test"
	const group = "test_group"

	// Declare the number of subscriptions we will use and the
	// number of messages to send.
	const subcriptions = 2
	const msgs = 10

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Used to wait for the messages to be received.
	var wg sync.WaitGroup
	wg.Add(msgs)

	// Create a slice of the subscriptions we will create.
	subs := make([]*nats.Subscription, subcriptions)

	// Create subscriptions to receive messages. The messages
	// will be delivered round robin across these subcriptions.
	for i := 0; i < subcriptions; i++ {

		// Create a local variable for the id.
		id := i + 1

		// Function is called when a new message is received.
		f := func(m *nats.Msg) {
			log.Println(id, "Received a message:", string(m.Data))
			wg.Done()
		}

		// Subscribe to receive messages for the specified
		// subject and group.
		var err error
		if subs[i], err = conn.QueueSubscribe(subject, group, f); err != nil {
			log.Println("Subscribing for specified subject:", err)
			return
		}
	}

	// Send messages for the specified subject. The queue of handlers
	// will take turns handling the messages.
	for i := 1; i <= msgs; i++ {

		// Generate the message to send.
		s := fmt.Sprintf("%d: Hello World", i)

		// Publish the message for the specified subject.
		if err := conn.Publish(subject, []byte(s)); err != nil {
			log.Println("Publishing a message for specified subject:", err)
			return
		}
	}

	// Wait to be told all the messages were received.
	wg.Wait()

	// Rangle over the subscriptions and unsubscribe from all of them.
	for _, sub := range subs {

		// Unsubscribe from receiving messages for this subscription.
		if err := sub.Unsubscribe(); err != nil {
			log.Println("Error unsubscribing from the bus:", err)
			return
		}
	}

	// Close the connection to the NATS server.
	conn.Close()
}
