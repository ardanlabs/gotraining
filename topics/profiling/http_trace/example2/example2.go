// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the http trace with a unique Client
// and Transport.
package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	log.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func main() {

	// Create a new request for the call.
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Create the transport value we are binding event to.
	var t transport

	// Create a ClientTrace value for the events we care about.
	trace := httptrace.ClientTrace{
		GotConn: t.GotConn,
	}

	// Bind to the request context this new context for tracing.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), &trace))

	// Create a client value with our transport.
	client := http.Client{
		Transport: &t,
	}

	// Make the request call and get the tracing informaion.
	// The program will follow the redirect of google.com to
	// www.google.com and will output:
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
