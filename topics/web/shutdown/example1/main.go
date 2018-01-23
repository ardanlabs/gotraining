// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program shows how to launch a web server then shut it down gracefully.
package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// app is our application handler. We log requests when they start so we can
// see them happening then log again when they're over. We have a random sleep
// from 800-1200 milliseconds so everything goes slow enough to see.
func app(res http.ResponseWriter, req *http.Request) {
	id := time.Now().Nanosecond()
	log.Printf("Request Start %d", id)

	sleep := rand.Intn(400) + 800
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	log.Printf("Request End   %d", id)
}

func main() {

	// Create a new server and set timeout values.
	server := http.Server{
		Addr:           "localhost:3000",
		Handler:        http.HandlerFunc(app),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the listener.
	go func() {
		log.Println("Listening on", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	// Block waiting for a receive on either channel
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-osSignals:

		// Create a context to attempt a graceful 5 second shutdown.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Attempt the graceful shutdown by closing the listener and
		// completing all inflight requests.
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop server gracefully: %v", err)
			log.Print("Initiating hard shutdown")
			if err := server.Close(); err != nil {
				log.Fatalf("Could not stop http server: %v", err)
			}
		}
	}

	log.Println("Shut down successful")
}
