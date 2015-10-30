// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/Lc3tS8kprX

// Basic command line program that accepts arguments to websites and retrieves
// them and record response times. Uses concurrency
package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Use a waitgroup to make sure all go routines finish
	var wg sync.WaitGroup

	// Lets keep a reference to when we started
	start := time.Now()

	// Set the value for the waitgroup
	wg.Add(len(sites))

	for _, site := range sites {
		// Launch each retrieval in a go routine.  This makes each request concurrent
		go func(site string) {
			defer wg.Done()

			// start a timer for this request
			begin := time.Now()

			// Retrieve the site
			if _, err := http.Get(site); err != nil {
				fmt.Println(site, err)
				return
			}

			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}(site)
	}

	// Block until all routines finish
	wg.Wait()

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
