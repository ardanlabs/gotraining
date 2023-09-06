package main

import (
	"sync"
	"testing"
	"time"
)

func TestCouter(t *testing.T) {
	counter := 0
	n, m := 10, 1000

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < m; i++ {
				time.Sleep(time.Microsecond)
				counter++
			}
		}()
	}

	wg.Wait()
	if counter != m*n {
		t.Fatal(counter)
	}
}
