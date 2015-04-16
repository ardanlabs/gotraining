# Concurrency

In this quick start guide, we will cover the basics of working concurrently.

## Retrieving Websites

Sometimes, we need to do a lot of operations that will take an unkown length of time.
We want to do this as fast and effeciently as possible, so we can use Go's built in concurrecy.

## The problem

To keep things simple, let's pretend we have a list of websites we want to monitor response times.


## Basic CLI

We will start with a basic command line program that accepts arguments as the websites to retrieve.

```go
// https://play.golang.org/p/GDaMvunNMZ

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	// flag.Args contains all non-flag arguments
	fmt.Println(flag.Args())
}
```

Running this program with the following arguments:

```sh
go run monitor.go http://google.com http://yahoo.com
```

You will see that we output the following:

```sh
[http://google.com http://yahoo.com]
```

## Retreiving and recording response times

Now we want to retrieve them and record response times.

Change the program as such:

```go
// https://play.golang.org/p/I-gUNt3biw

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Lets keep a reference to when we started
	start := time.Now()

	for _, site := range sites {
		// start a timer for this request
		begin := time.Now()
		
		// Retreive the site
		if _, err := http.Get(site); err != nil {
			fmt.Println(site, err)
			continue
		}
		
		fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
	}

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
```

Now run the program:

```sh
go run monitor.go http://google.com http://yahoo.com
```

Times may vary, but this is the output I received:

```sh
Site "http://google.com" took 119.178774ms to retrieve.
Site "http://yahoo.com" took 581.01092ms to retrieve.
Entire process took 700.245713ms
```

## Make it concurrent

Ok, as we can see, we should be able to speed this up quite a bit by making the program concurrent.

Make the following changes:

```go
// https://play.golang.org/p/Lc3tS8kprX

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
			
			// Retreive the site
			if _, err := http.Get(site); err != nil {
				fmt.Println(site, err)
				continue
			}
			
			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}(site)
	}

	// Block until all routines finish
	wg.Wait()

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
```

Run it again:

```sh
go run monitor.go http://google.com http://yahoo.com
```

Now you should see an output something like this when you run it:

```sh
Site "http://google.com" took 131.805645ms to retrieve.
Site "http://yahoo.com" took 541.624279ms to retrieve.
Entire process took 541.743794ms
```

Notice how it not only takes slightly longer than the longest request?  This is due to everything running concurrently.

The big difference with the final program is the use of a `WaitGroup` and launching everything
inside of a go routine by using the `go func` signature.

You can read more on wait groups in the [Sync Package](http://golang.org/pkg/sync).

## Summary

Congratulations, you just wrote your first concurrenty program!
