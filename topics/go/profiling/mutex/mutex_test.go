// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to profile mutexes.
package mutex

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	// data is a slice that will be shared.
	data = make([]string, 1000)

	// rwMutex is used to define a critical section of code.
	rwMutex sync.RWMutex
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// TestMutexProfile creates goroutines that will content.
func TestMutexProfile(t *testing.T) {
	t.Log("Starting Test")

	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			writer()
			wg.Done()
		}()

		go func() {
			reader()
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log("Test Complete")
}

// writer adds 10 new strings to the slice in random intervals.
func writer() {
	for i := 0; i < 10; i++ {
		rwMutex.Lock()
		{
			data = append(data, "A")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
		rwMutex.Unlock()
	}
}

// reader wakes up and iterates over the data slice.
func reader() {
	for i := 0; i < 10; i++ {
		rwMutex.RLock()
		{
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
		rwMutex.RUnlock()
	}
}
