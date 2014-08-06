# Race Conditions - Concurrency and Channels

### Code Review

[Race Condition](example1/example1.go) ([Go Playground](http://play.golang.org/p/DCkt7qIzB8))

[Atomic Increments](example2/example2.go) ([Go Playground](http://play.golang.org/p/4CaNFVZaXn))

[Atomic Store/Load](example3/example3.go) ([Go Playground](http://play.golang.org/p/zW9bjwLzzH))

[Mutex](example4/example4.go) ([Go Playground](http://play.golang.org/p/4ZPCUnWyV5))

### Exercise 1
Given the following program, find and correct the race condition. [Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/vW-48gPin1))

	// http://play.golang.org/p/rxBadio1I2

	// Program for an exercise to fix a race condition.
	package main

	import (
		"fmt"
		"math/rand"
		"sync"
		"time"
	)

	var (
		// numbers maintains a set of random numbers.
		numbers []int

		// wg is used to wait for the program to finish.
		wg sync.WaitGroup
	)

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
		for index, number := range numbers {
			fmt.Println(index, number)
		}
	}

	// random generates random numbers and stores them into a slice.
	func random(amount int) {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}
	}

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)