// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	// Create the first goroutine and manage its lifecycle here.
	go func() {
		printPrime("A")
		wg.Done()
	}()

	// Create the second goroutine and manage its lifecycle here.
	go func() {
		printPrime("B")
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime logs and displays prime numbers for the first 10 numbers.
func printPrime(prefix string) {
	file, err := os.Create(prefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { file.Close(); os.Remove(prefix) }()
	mw := io.MultiWriter(os.Stdout, file)

next:
	for outer := 2; outer < 10; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Fprintf(mw, "%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
