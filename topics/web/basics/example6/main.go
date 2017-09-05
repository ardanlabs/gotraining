// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how http servers already handle concurrent
// requests. Individual requests being slow do not block others.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	// Launch a goroutine to run the web service.
	go func() {

		m := http.NewServeMux()

		// The main handler for our service
		m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			log.Println(req.URL.Path)

			// Create some fake latency.
			time.Sleep(4 * time.Second)

			// Any attempt to brew coffee with a teapot should result in the
			// HTTP error code 418 I'm a teapot and the resulting entity
			// body MAY be short and stout.
			res.WriteHeader(http.StatusTeapot)
		})

		// A route that returns 200 OK when the service is up
		m.HandleFunc("/ready", func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
		})

		log.Print("Listening on localhost:3000")
		log.Fatal(http.ListenAndServe("localhost:3000", m))
	}()

	// Call /ready until we get a 200 OK
	waitForReady()

	// Get the current time so we can measure how long this all takes.
	start := time.Now()

	// Call the handler function 100 times.
	process(100)

	// Display how long all of the requests took.
	fmt.Printf("\nduration: %s\n", time.Now().Sub(start))
}

// waitForReady calls the /ready endpoint until it gets a successful response.
// It waits for 100ms after the first attempt, 200ms after the second, and so
// on for up to 20 attempts. If it does not succeed after that it kills the
// application.
func waitForReady() {
	max := 20
	for attempts := 1; attempts <= max; attempts++ {
		res, err := http.Get("http://localhost:3000/ready")
		if err == nil && res.StatusCode == http.StatusOK {
			return
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}

	log.Fatal("application did not start in time")
}

// process makes n concurrent requests against our service.
func process(n int) {
	var w sync.WaitGroup
	w.Add(n)

	for i := 0; i < n; i++ {
		go func(index int) {
			call(index)
			w.Done()
		}(i)
	}

	w.Wait()
}

// call makes a single request to the service we started.
func call(i int) {

	res, err := http.Get(fmt.Sprintf("http://localhost:3000/%d", i))
	if err != nil {
		log.Fatal(err)
	}

	// We should get the Teapot status.
	if res.StatusCode != http.StatusTeapot {
		log.Fatal("Oops!")
	}
}
