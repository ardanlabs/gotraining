// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that implements a web request with a context that is
// used to timeout the request if it takes too long.
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Create a new request.
	req, err := http.NewRequest("GET", "https://www.goinggo.net/post/index.xml", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create a context with a timeout of 50 milliseconds.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// Perform the call which much finish in 50 milliseconds.
	if err = httpDo(ctx, req, handler); err != nil {
		log.Println(err)
		return
	}
}

// handler is used to hande the response from the request.
func handler(resp *http.Response, err error) error {

	// If we received an error, just return it.
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)
	return nil
}

// httpDo performs a request and will return when the request completes
// or the specified time based on the context runs out.
func httpDo(ctx context.Context, req *http.Request, h func(*http.Response, error) error) error {

	// Declare a new transport and client for the call.
	var tr http.Transport
	client := http.Client{
		Transport: &tr,
	}

	// Make the request call in a separate goroutine.
	ch := make(chan error, 1)
	go func() {
		log.Println("Starting Request")

		// Call the provided handler function which processes
		// the request and returns error information.
		ch <- h(client.Do(req))
	}()

	// Wait the request or timeout.
	select {
	case <-ctx.Done():
		log.Println("Context Timedout")

		// Cancel the request.
		tr.CancelRequest(req)

		log.Println("Waiting For Request To Cancel")

		// Wait for request to return.
		<-ch
		return ctx.Err()

	case err := <-ch:

		// The request is complete.
		return err
	}
}
