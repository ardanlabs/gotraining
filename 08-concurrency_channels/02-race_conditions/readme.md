## Race Conditions - Concurrency and Channels

A race condition is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that totally appear random or can be never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

## Notes

* Goroutines need to be coordinated and synchronized.
* When two or more goroutines attempt to access the same resource, we have a race condition.
* Atomic functions and mutexes can provide the support we need.

## Cache Coherency
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (30:09-32:10)](https://youtu.be/WDIkqP4JbkE?t=1809)

## Links

http://blog.golang.org/race-detector

http://www.goinggo.net/2013/09/detecting-race-conditions-with-go.html

https://golang.org/doc/articles/race_detector.html

## Documentation

[Race Condition Diagram](documentation/race_condition.md)

## Code Review

[Race Condition](example1/example1.go) ([Go Playground](https://play.golang.org/p/IwFKbnb1JO))

[Atomic Increments](example2/example2.go) ([Go Playground](https://play.golang.org/p/1Wlq7FDgHd))

[Atomic Store/Load](example3/example3.go) ([Go Playground](https://play.golang.org/p/blXRQXlH8Z))

[Mutex](example4/example4.go) ([Go Playground](https://play.golang.org/p/E7Gj_pSpqf))

[Read/Write Mutex](example5/example5.go) ([Go Playground](https://play.golang.org/p/jVPu5BFuPT))

## Exercises

### Exercise 1
Given the following program, use the race detector to find and correct the race condition.

	// https://play.golang.org/p/0C-mUZGUhE

	// Program for an exercise to fix a race condition.
	package main

	import (
		"fmt"
		"math/rand"
		"sync"
		"time"
	)

	// numbers maintains a set of random numbers.
	var numbers []int

	// wg is used to wait for the program to finish.
	var wg sync.WaitGroup

	// init is called prior to main.
	func init() {
		rand.Seed(time.Now().UnixNano())
	}

	// main is the entry point for all Go programs.
	func main() {
		// Add a count for each goroutine we will create.
		wg.Add(3)

		// Create three goroutines to generate random numbers.
		go random(10)
		go random(10)
		go random(10)

		// Wait for all the goroutines to finish.
		wg.Wait()

		// Display the set of random numbers.
		for i, number := range numbers {
			fmt.Println(i, number)
		}
	}

	// random generates random numbers and stores them into a slice.
	func random(amount int) {
		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}

		// Tell main we are done.
	    wg.Done()
	}

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/yBFA-MDcMw)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/wFTNvVoBpz))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
