// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show to connect and publish/subscribe for messages.
// Message are received asynchronously using a channel.
package main

import (
	"log"

	"github.com/nats-io/nats"
)

// user represents a user in the system.
type user struct {
	Name  string
	Email string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Declare the subject to use for publishing/subscribing.
	const subject = "test"

	// Connect to the local nats server.
	rawConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// Create an encoded connection
	conn, err := nats.NewEncodedConn(rawConn, nats.JSON_ENCODER)
	if err != nil {
		log.Println("Unable to create an encoded connection")
		return
	}

	// Make and bind a channel to receiving user values.
	recv := make(chan user)
	subRecv, err := conn.BindRecvChan(subject, recv)
	if err != nil {
		log.Println("Unable to bind the receive channel")
		return
	}

	// Make and bind a channel to send user values.
	send := make(chan user)
	if err := conn.BindSendChan(subject, send); err != nil {
		log.Println("Unable to bind the send channel")
		return
	}

	// Create a value of type user.
	u1 := user{
		Name:  "bill",
		Email: "bill@ardanlabs.com",
	}

	// Send the value to the message bus.
	send <- u1

	// Receiving the value from the message bus.
	u2 := <-recv

	// Display the user.
	log.Println("Received a user:", u2.Name)

	// Unsubscribe from receiving users.
	if err := subRecv.Unsubscribe(); err != nil {
		log.Println("Error unsubscribing from the receive subscription:", err)
		return
	}

	// Close the connection to the NATS server.
	conn.Close()
}
