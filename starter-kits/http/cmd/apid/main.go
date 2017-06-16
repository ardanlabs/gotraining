// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/cmd/apid/handlers"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/db"
)

// init is called before main. We are using init to customize logging output.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	log.Println("main : Started")

	// Check the environment for a configured port value.
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "got:got2015@ds039441.mongolab.com:39441/gotraining"
	}

	// Register the Master Session for the database.
	log.Println("main : Started : Capturing Master DB...")
	masterDB, err := db.NewMGO(dbHost, 25*time.Second)
	if err != nil {
		log.Fatalf("startup : Register DB : %v", err)
	}
	defer masterDB.MGOClose()

	// Check the environment for a configured port value.
	host := os.Getenv("HOST")
	if host == "" {
		host = ":3000"
	}

	// Create a new server and set timeout values.
	server := http.Server{
		Addr:           host,
		Handler:        handlers.API(masterDB),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// We want to report the listener is closed.
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the listener.
	go func() {
		log.Printf("startup : Listening %s", host)
		log.Printf("shutdown : Listener closed : %v", server.ListenAndServe())
		wg.Done()
	}()

	// Listen for an interrupt signal from the OS.
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
