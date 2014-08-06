// http://play.golang.org/p/ncWam67dS1

// Answer for exercise 1 of Channels.
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

	// number is a channel that will receive random numbers.
	number = make(chan int)
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

	// Create a goroutine to monitor when all the numbers
	// have been generated.
	go func() {
		wg.Wait()
		close(number)
	}()

	// Wait for and collect the numbers.
	for value := range number {
		numbers = append(numbers, value)
	}

	// Display the set of random numbers.
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}

// random generates random numbers and stores them into a slice.
func random(amount int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	// Generate as many random numbers as specified.
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)

		// Send the number into the channel.
		number <- n
	}
}
