// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/yBFA-MDcMw

// Fix the race condition in this program.
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
