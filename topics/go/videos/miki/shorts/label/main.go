// Use labels to break from nested loops and switch cases.

package main

import (
	"context"
	"log"
	"time"
)

func processEvents(ctx context.Context, ch <-chan Event) {
	count, start := 0, time.Now()
loop:
	for {
		select {
		case e, ok := <-ch:
			if !ok {
				break loop
			}
			handleEvent(e)
			count++
		case <-ctx.Done():
			break loop
		}
	}
	log.Printf("processed %d events in %v", count, time.Since(start))
}

func main() {
	ch := make(chan Event)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for i := 0; i < 7; i++ {
			ch <- Event{}
		}
		// cancel()
		close(ch)
	}()

	processEvents(ctx, ch)
}

// ---

type Event struct{}

func handleEvent(e Event) {
	time.Sleep(17 * time.Millisecond)
	log.Printf("event")
}
