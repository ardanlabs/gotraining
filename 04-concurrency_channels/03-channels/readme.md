# Channels - Concurrency and Channels
Channels are a reference type that provide a safe mechanism to share data between goroutines. Unbuffered channel give a guarantee of delivery that data has passed from one goroutine to the other. Buffered channels allow for data to pass through the channel without such guarantees. Unbuffered channels require both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can complete. Buffered channels don't force goroutines to be ready at the same instant to perform sends and receives.

### Code Review

[Unbuffered channels - Tennis game](example1/example1.go) ([Go Playground](http://play.golang.org/p/7WO_eOJx_G))

[Unbuffered channels - Relay race](example2/example2.go) ([Go Playground](http://play.golang.org/p/5B1MxmDuZI))

[Buffered channels - Control concurrency](example3/example3.go) ([Go Playground](http://play.golang.org/p/G9Gfy1drox))

### Advanced Code Review

[Timers](advanced/timer/timer.go)

[Semaphores](advanced/semaphore/semaphore.go)

[Pooling](advanced/pool/pool.go)

### Exercises

#### Exercise 1
Given the following program, convert it to use a channel. [Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/ncWam67dS1))

	// http://play.golang.org/p/vW-48gPin1

	// Answer for exercise 1 of Race Conditions.
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

		// mutex will help protect the slice.
		mutex sync.Mutex
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
			mutex.Lock()
			{
				numbers = append(numbers, n)
			}
			mutex.Unlock()
		}
	}

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)