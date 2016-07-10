package main

import "testing"

func TestLatency(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10000; i++ {
			ch <- i
		}
	}()
	for range ch {
		// do nothing (we're looking at overhead)
	}
}
