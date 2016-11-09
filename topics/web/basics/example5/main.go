// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to run the server using a goroutine and
// create goroutines to run multiple requests concurrently.
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

		// Create a new mux for handling routes.
		m := http.NewServeMux()

		// Bind a handler to the root route that
		m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			log.Println(req.URL.Path)

			// Create some fake latency.
			time.Sleep(1 * time.Second)

			// Any attempt to brew coffee with a teapot should result
			// in the HTTP error code 418 I'm a teapot and the resulting
			// entity body MAY be short and stout.
			res.WriteHeader(http.StatusTeapot)
		})

		// Start the http server to handle the request.
		http.ListenAndServe(":3000", m)
	}()

	// Get the current time so we can time how long this request takes.
	start := time.Now()

	// Call the handler function 100 times.
	call(100, func(i int) {

		// Call into the running service we started above.
		res, err := http.Get(fmt.Sprintf("http://localhost:3000/%d", i))
		if err != nil {
			log.Fatal(err)
		}

		// We should get the Teapot status.
		if res.StatusCode != http.StatusTeapot {
			log.Fatal("Oops!")
		}
	})

	// Display how long the request took.
	fmt.Printf("\nduration: %s\n", time.Now().Sub(start))
}

// call executes the provided function count number
// of times concurrently.
func call(count int, f func(index int)) {
	var w sync.WaitGroup
	w.Add(count)

	for i := 0; i < count; i++ {
		go func(index int) {
			f(index)
			w.Done()
		}(i)
	}

	w.Wait()
}
