// Package service maintains the logic for the web service.
package service

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/braintree/manners"
)

// init binds the routes and handlers for the web service.
func init() {

	// Setup a route for our static files.
	//
	// Because our static directory is set as the root of the FileSystem,
	// we need to strip off the /static/ prefix from the request path
	// before searching the FileSystem for the given file.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup a route for the home page.
	http.HandleFunc("/search", handler)
}

// Run binds the service to a port and starts listening for requests.
func Run() {
	host := "localhost:5000"
	readTimeout := 10 * time.Second
	writeTimeout := 30 * time.Second

	// Create a new server and set timeout values.
	s := manners.NewWithServer(&http.Server{
		Addr:           host,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	})

	// Support for shutting down cleanly.
	go func() {

		// Listen for an interrupt signal from the OS.
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan

		// We have been asked to shutdown the server.
		log.Println("Starting shutdown...")
		s.Close()

		// For now until I deal with manners handling static files.
		go func() {
			time.Sleep(time.Second * 60)
			log.Println("Killed Service")
			os.Exit(1)
		}()
	}()

	log.Println("Listening on:", host)
	s.ListenAndServe()
}
