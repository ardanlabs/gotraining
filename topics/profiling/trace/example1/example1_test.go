package main

import (
	"sync"
	"testing"
	"time"
)

func TestDelay(t *testing.T) {
	wait := make(chan struct{})
	var mu sync.Mutex
	go func() {
		mu.Lock()
		defer mu.Unlock()
		t.Log("inner acquired, sleeping")
		close(wait)
		time.Sleep(1 * time.Second)
	}()
	<-wait
	t.Log("outer waiting for lock")
	mu.Lock()
	defer mu.Unlock()
	t.Log("outer acquired")
}
