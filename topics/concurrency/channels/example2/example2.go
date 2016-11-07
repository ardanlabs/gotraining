// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use an unbuffered channel to
// simulate a relay race between four goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {

	// Create an unbuffered channel.
	track := make(chan int)

	// Add a count of one for the last runner.
	wg.Add(1)

	// First runner to his mark.
	go Runner(track)

	// Start the race.
	track <- 1

	// Wait for the race to finish.
	wg.Wait()
}

// Runner simulates a person running in the relay race.
func Runner(track chan int) {

	// The number of exchanges of the baton.
	const maxExchanges = 4

	var exchange int

	// Wait to receive the baton.
	baton := <-track

	// Start running around the track.
	fmt.Printf("Runner %d Running With Baton\n", baton)

	// New runner to the line.
	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("Runner %d To The Line\n", exchange)
		go Runner(track)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if baton == maxExchanges {
		fmt.Printf("Runner %d Finished, Race Over\n", baton)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		baton,
		exchange)

	track <- exchange
}
