// Jason Waldrip [6:44 PM]
// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package work manages a pool of goroutines to perform work.
// Example provided with help from Jason Waldrip.
package work

import "sync"

// Worker must be implemented by types that want to use
// this worker processes.
type Worker interface {
	Work()
}

// Work provides a pool of goroutines that can execute any Worker
// tasks that are submitted.
type Work struct {
	tasks chan Worker
	wg    sync.WaitGroup
}

// New creates a new Worker.
func New(goroutines int) *Work {
	w := Work{
		tasks: make(chan Worker),
	}

	w.wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for task := range w.tasks {
				task.Work()
			}
			w.wg.Done()
		}()
	}

	return &w
}

// RunTask to the worker queue
func (w *Work) RunTask(task Worker) {
	w.tasks <- task
}

// Shutdown waits for all the workers to finish
func (w *Work) Shutdown() {
	close(w.tasks)
	w.wg.Wait()
}
