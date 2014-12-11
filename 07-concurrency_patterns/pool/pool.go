// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This example is provided with help by Gabriel Aszalos.

// Package pool manages a user defined set of resources.
// Based on the work by Fatih Arslan with his pool package.
package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// Pool manages a set of resources that can be shared safely by multiple goroutines.
// The resource being managed must implement to io.Closer interface.
type Pool struct {
	sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrInvalidCapacity is returned when there has been an attempt to create an
// unbuffered pool.
var ErrInvalidCapacity = errors.New("Capacity needs to be greater than zero.")

// New creates a pool that manages resources. A pool requires a function
// that can allocate a new resource and the number of resources that can
// be allocated.
func New(fn func() (io.Closer, error), capacity uint) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, capacity),
	}, nil
}

// Acquire retrieves a resource	from the pool.
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// Check for a free resource.
	case r, ok := <-p.resources:
		fmt.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, errors.New("Pool has been closed.")
		}
		return r, nil

	// Provide a new resource since there are none available.
	default:
		fmt.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places a new resource onto the pool.
func (p *Pool) Release(r io.Closer) {
	// Secure this operation with the Close operation.
	p.Lock()
	defer p.Unlock()

	// If the pool is closed, discard the resource.
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue.
	case p.resources <- r:
		fmt.Println("Release:", "In Queue")

	// If the queue is already at capacity we close the resource.
	default:
		fmt.Println("Release:", "Closing")
		r.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	// Secure this operation with the Release operation.
	p.Lock()
	defer p.Unlock()

	// If the pool is already close, don't do anything.
	if p.closed {
		return
	}

	// Toggle the flag
	p.closed = true

	// Close the channel before we drain the channel of its
	// resources. If we don't do this, we will have a deadlock.
	close(p.resources)

	// Close the resources
	for r := range p.resources {
		r.Close()
	}
}
