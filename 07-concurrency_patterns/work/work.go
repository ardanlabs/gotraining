// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package work manages a pool of goroutines to perform work.
package work

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Worker must be implemented by types that want to use
// this worker processes.
type Worker interface {
	Work()
}

// Stats contains information about the work pool.
type Stats struct {
	Goroutines int64
	Pending    int64
	Active     int64
}

// Work provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Work struct {
	tasks    chan Worker    // Unbuffered channel that work is sent into.
	kill     chan struct{}  // Unbuffered channel to signal for a goroutine to die.
	shutdown chan struct{}  // Closed when the Work pool is being shutdown.
	wg       sync.WaitGroup // Manages the number of goroutines for shutdown.
	mutex    sync.Mutex     // Provides synchronization for managing the pool.
	stats    Stats          // Stats about the health of the pool.
}

// New creates a new Worker.
func New(goroutines int) *Work {
	w := Work{
		tasks:    make(chan Worker),
		kill:     make(chan struct{}),
		shutdown: make(chan struct{}),
	}

	w.Add(goroutines)

	return &w
}

// LogStats display work pool stats on the specified duration.
func (w *Work) LogStats(d time.Duration) {
	w.wg.Add(1)

	go func() {
		for {
			select {
			case <-w.shutdown:
				w.wg.Done()
				return

			case <-time.After(d):
				s := w.Stats()
				fmt.Printf("G[%d] P[%d] A[%d]\n", s.Goroutines, s.Pending, s.Active)
			}
		}
	}()
}

// Stats returns the current status for the work pool.
func (w *Work) Stats() Stats {
	var s Stats
	s.Goroutines = atomic.LoadInt64(&w.stats.Goroutines)
	s.Pending = atomic.LoadInt64(&w.stats.Pending)
	s.Active = atomic.LoadInt64(&w.stats.Active)

	return s
}

// Add creates goroutines to process work or sets a count for
// goroutines to terminate.
func (w *Work) Add(goroutines int) {
	if goroutines == 0 {
		return
	}

	w.mutex.Lock()
	{
		if goroutines > 0 {
			// We are adding goroutines to the pool.
			w.wg.Add(goroutines)
			atomic.AddInt64(&w.stats.Goroutines, int64(goroutines))

			for i := 0; i < goroutines; i++ {
				go w.work()
			}
		} else {
			// We are removing goroutines from the pool.
			goroutines = goroutines * -1
			current := int(atomic.LoadInt64(&w.stats.Goroutines))
			if goroutines > current {
				goroutines = current
			}

			// Send the kill signal and wait for these goroutines
			// to die.
			for i := 0; i < goroutines; i++ {
				w.kill <- struct{}{}
			}
		}
	}
	w.mutex.Unlock()
}

// work performs the users work and keeps stats.
func (w *Work) work() {
done:
	for {
		select {
		case t, ok := <-w.tasks:
			if !ok {
				break done
			}

			atomic.AddInt64(&w.stats.Active, 1)
			t.Work()
			atomic.AddInt64(&w.stats.Active, -1)

		case <-w.kill:
			break done
		}
	}

	atomic.AddInt64(&w.stats.Goroutines, -1)
	w.wg.Done()
}

// Run wait for the goroutine pool to take the work
// to be executed.
func (w *Work) Run(work Worker) {
	atomic.AddInt64(&w.stats.Pending, 1)
	w.tasks <- work
	atomic.AddInt64(&w.stats.Pending, -1)
}

// Shutdown waits for all the workers to finish.
func (w *Work) Shutdown() {
	close(w.tasks)
	close(w.shutdown)
	w.wg.Wait()
}
