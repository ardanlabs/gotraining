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
		Addr:           ":3000",
		Handler:        http.HandlerFunc(app),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start listening for requests.
	go func() {
		log.Println("listener : Listening on :3000")
		log.Println("listener :", server.ListenAndServe())
	}()

	// Block until there's an interrupt then shut the server down. The main
	// func must not return before this process is complete or in-flight
	// requests will be aborted.
	blockUntilShutdown(&server)

	log.Println("main : Completed")
}

// blockUntilShutdown holds the goroutine waiting for a signal to shutdown.
func blockUntilShutdown(server *http.Server) {

	// Set up channel to receive interrupt signals.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c
	log.Println("shutdown : Attempting graceful shut down...")

	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt the graceful shutdown.
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown : Graceful shutdown did not complete in %v : %v", timeout, err)

		// Looks like we timedout on the graceful shutdown. Kill it hard.
		if err := server.Close(); err != nil {
			log.Printf("shutdown : Error killing server : %v", err)
		}
	}
}
