// This is a simple example put together to help a friend with the
// idea of not over-engineering a pubsub pattern. This is entirely
// production ready but more of a prototype and concept.
package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	clients := NewClients()

	publisher := NewPublisher(clients)
	defer publisher.Shutdown()

	clients.Add("1")
	clients.Add("2")
	clients.Add("3")
	time.Sleep(time.Second)
	clients.Remove("2")
	time.Sleep(time.Second)
	clients.Remove("1")
	time.Sleep(time.Second)
	clients.Add("2")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("shutting down")
}
