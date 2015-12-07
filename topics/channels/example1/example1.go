// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/kuxUFMqy-9

// Sample program to show how to use an unbuffered channel to
// simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup

	// Create an unbuffered channel.
	court := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two players.
	go func() {
		player("Serena", court)
		wg.Done()
	}()

	go func() {
		player("Venus", court)
		wg.Done()
	}()

	// Start the set.
	court <- 1

	// Wait for the game to finish.
	wg.Wait()
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
