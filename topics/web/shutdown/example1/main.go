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
	"sync"
	"time"
)

// app is our application handler. We log requests when they start so we can
// see them happening then log again when they're over. We have a random sleep
// from 800-1200 milliseconds so everything goes slow enough to see.
func app(res http.ResponseWriter, req *http.Request) {
	id := time.Now().Nanosecond()
	log.Printf("app : Start %d", id)

	sleep := rand.Intn(400) + 800
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	log.Printf("app : End   %d", id)
}

func main() {

	log.Println("main : Started")

	// Create a new server and set timeout values.
	server := http.Server{
		Addr:           "localhost:3000",
		Handler:        http.HandlerFunc(app),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// We want to report the listener is closed.
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the listener.
	go func() {
		log.Println("listener : Listening on localhost:3000")
		log.Println("listener :", server.ListenAndServe())
		wg.Done()
	}()

	// Listen for an interrupt signal from the OS. Use a buffered
	// channel because of how the signal package is implemented.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)

	// Wait for a signal to shutdown.
	<-osSignals

	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt the graceful shutdown by closing the listener and
	// completing all inflight requests.
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown : Graceful shutdown did not complete in %v : %v", timeout, err)

		// Looks like we timedout on the graceful shutdown. Kill it hard.
		if err := server.Close(); err != nil {
			log.Printf("shutdown : Error killing server : %v", err)
		}
	}

	// Wait for the listener to report it is closed.
	wg.Wait()
	log.Println("main : Completed")
}
