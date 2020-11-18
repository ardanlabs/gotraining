package main

import (
	"fmt"
	"log"
	"time"
)

// Publisher is consuming messages and publishing them.
type Publisher struct {
	clients *Clients
}

// NewPublisher connects to the publisher can recieves messages.
func NewPublisher(clients *Clients) *Publisher {
	pub := Publisher{
		clients: clients,
	}

	// This is not production ready since it can't be shutdown.
	// For now we just want a stream of messages.

	go func() {
		var counter int
		for {
			time.Sleep(100 * time.Millisecond)
			counter++
			log.Println("publisher: mesage received : sending to clients")
			clients.Send(fmt.Sprintf("message %d", counter))
		}
	}()

	return &pub
}

// Shutdown disconnects the publisher and stop messages.
func (p *Publisher) Shutdown() {

	// TO BE IMPLEMENTED
}
