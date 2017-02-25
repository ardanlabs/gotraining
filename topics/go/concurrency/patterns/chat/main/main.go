// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This example is provided with help by Gabriel Aszalos.

// This sample program demonstrates how to create a simple chat system.
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ardanlabs/gotraining/topics/concurrency_patterns/chat"
)

func main() {
	cr := chat.New()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting Down Started")
	cr.Close()
	log.Println("Shutting Down Completed")
}
