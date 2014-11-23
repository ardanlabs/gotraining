// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package pool manages a user defined set of resources.
// Based on the work by Fatih Arslan with his pool package.
// https://github.com/fatih/pool/tree/v1.0.0
package pool

import (
	"errors"
	"fmt"
	"sync"
)

type (
	// Resource must be implemented by types to use the pool.
	Resource interface {
		Close()
	}

	// Factory is a user supplied function that creates resources.
	Factory func() (Resource, error)
)

type (
	// resources is a named type for the channel of resources.
	queue chan Resource

	// Pool contains a queue of resources and methods that provide
	// support to share resources between goroutines.
	Pool struct {
		mutex     sync.Mutex
		resources queue
		factory   Factory
		closed    bool
	}
)

// New creates a pool for managing resources.
func New(factory Factory, capacity int) (*Pool, error) {
	// Check the capacity is greater than zero else
	// we could create an unbuffered channel or panic.
	if capacity <= 0 {
		return nil, fmt.Errorf("Invalid Capacity Value: %d", capacity)
	}

	return &Pool{
		factory:   factory,
		resources: make(queue, capacity),
	}, nil
}

// Acquire retrieves a resource	from the pool.
func (p *Pool) Acquire() (Resource, error) {
	select {
	case resource, ok := <-p.resources:
		fmt.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, errors.New("Pool has been closed.")
		}
		return resource, nil

	default:
		fmt.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release places the resource back into the queue or closes
// the resource for good.
func (p *Pool) Release(resource Resource) {
	// Secure this operation with the Close operation.
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// If the poll is closed, close the resource before returning.
	if p.closed {
		resource.Close()
		return
	}

	// The queue is still open so release the resource.

	select {
	// Attempt to place the resource back into the queue first. If the queue
	// is full, then the default case  will be executed.
	case p.resources <- resource:
		fmt.Println("Release:", "In Queue")

	// The queue is full so just close this resource.
	default:
		fmt.Println("Release:", "Closing")
		resource.Close()
	}
}

// Close will shutdown the pool and close all existing resources.
func (p *Pool) Close() {
	// Secure this operation with the Release operation.
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// If the pool is not closed, the close it down.
	if !p.closed {
		p.closed = true

		// Close the queue and close all existing resources.
		close(p.resources)
		for resource := range p.resources {
			resource.Close()
		}
	}
}
