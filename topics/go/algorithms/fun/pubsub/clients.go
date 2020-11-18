package main

import (
	"fmt"
	"log"
	"sync"
)

// Clients manage clients who are looking to receive messages.
type Clients struct {
	capacity int
	clients  map[string]chan string
	mu       sync.Mutex
}

// NewClients returns a clients management value.
func NewClients() *Clients {
	return &Clients{
		capacity: 1024,
		clients:  make(map[string]chan string),
	}
}

// Add places a client in the list for receiving messages.
func (c *Clients) Add(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[id]; exists {
		return fmt.Errorf("client id already exists: %s", id)
	}

	c.clients[id] = make(chan string, c.capacity)

	return nil
}

// Remove takes a client out of the list.
func (c *Clients) Remove(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[id]; !exists {
		return fmt.Errorf("client id doesn't exist: %s", id)
	}

	delete(c.clients, id)
	return nil
}

// Send will deliver the message to all existing clients.
func (c *Clients) Send(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the other side is not ready, they lose the message.
	// There is a capacity so messages would only be lost if
	// the client is not responding.

	for id, ch := range c.clients {
		select {
		case ch <- message:
			log.Println("client: sent: to: ", id)
		default:
			log.Println("client: timeout: to: ", id)
		}
	}
}
