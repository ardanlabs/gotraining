// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package work manages a pool of goroutines to perform work.
// Example provided with help from Jason Waldrip.
package work

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Worker must be implemented by types that want to use
// this worker processes.
type Worker interface {
	Work()
}

// Work provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Work struct {
	tasks        chan Worker
	log          chan struct{}
	wg           sync.WaitGroup
	countPending uint64
	countActive  uint64
}

// New creates a new Worker.
func New(goroutines int) *Work {
	w := Work{
		tasks: make(chan Worker),
		log:   make(chan struct{}),
	}

	w.wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for task := range w.tasks {
				atomic.AddUint64(&w.countActive, 1)
				task.Work()
				atomic.AddUint64(&w.countActive, ^uint64(0))
			}
			w.wg.Done()
		}()
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				log.Printf("HTTP request pool, pending=%v, active=%v\n", atomic.LoadUint64(&w.countPending), atomic.LoadUint64(&w.countActive))
			case <-w.log:
				return
			}
		}
	}()

	return &w
}

// RunTask to the worker queue
func (w *Work) RunTask(task Worker) {
	atomic.AddUint64(&w.countPending, 1)
	w.tasks <- task
	atomic.AddUint64(&w.countPending, ^uint64(0))
}

// Shutdown waits for all the workers to finish
func (w *Work) Shutdown() {
	close(w.tasks)
	close(w.log)
	w.wg.Wait()
}
