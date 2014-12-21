// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// export GODEBUG=schedtrace=50
// runqueue=1 [9]	1 in global queue and 9 in the context.

// export GOMAXPROCS=2
// runqueue=2 [2 2]	2 in global queue and 2 in each context.

// export GODEBUG=schedtrace=50,scheddetail=1

// Sample program to review scheduler stats.
package main

import (
	"sync"
	"time"
)

// Create a waitgroup.
var wg sync.WaitGroup

var m sync.Mutex

// main is the entry point for the application.
func main() {
	// We are going to create 10 goroutines.
	wg.Add(10)

	// Create those 10 goroutines.
	for i := 0; i < 10; i++ {
		go goroutine(i)
	}

	wg.Wait()
	time.Sleep(3 * time.Second)
}

func goroutine(i int) {
	time.Sleep(time.Second)

	for i := 0; i < 1e10; i++ {
		i++
	}

	wg.Done()
}
