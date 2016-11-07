// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to build a very basic chat client using NATS.
package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/nats-io/nats"
)

// channel represents the channel to use.
const channel = "general"

// user represents the internal user name.
var user string

// =============================================================================

func init() {

	// Generate a random user name.
	rand.Seed(time.Now().UnixNano())
	user = strconv.Itoa(rand.Intn(1000))
}

func main() {

	// Connect to the local nats server.
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Unable to connect to NATS")
		return
	}

	// This function gets called when a message is received.
	recv := func(m *nats.Msg) {

		// We are using a simple userid:message format.
		s := strings.Split(string(m.Data), ":")

		// If this was not our message, ignore it.
		if s[0] != user {
			WriteMessage(s[0], s[1])
		}
	}

	// Subscribe to receive messages for the specified key.
	sub, err := conn.Subscribe(channel, recv)
	if err != nil {
		log.Println("Subscribing for specified channel:", err)
		return
	}

	// This function gets called when the enter key is hit.
	event := func(s string) {

		// Not performing perfect length checking.
		if len(s) > 2 && s[:3] == "bot" {
			switch s[4:8] {
			case "name":
				user = s[9:]
				WriteMessage(user, "name set")
			}

			return
		}

		// Write the message to the screen.
		WriteMessage(user, s)

		// Add the user id to the message for delivery.
		send := user + ":" + s

		// Publish the message to NATS.
		if err := conn.Publish(channel, []byte(send)); err != nil {
			WriteMessage("Err", err.Error())
		}
	}

	// Draw the box and set the handler.
	Draw(event)

	// Unsubscribe from receiving these messages.
	if err := sub.Unsubscribe(); err != nil {
		log.Println("Error unsubscribing from NATS:", err)
		return
	}

	// Close the connection to the NATS server.
	conn.Close()
}
