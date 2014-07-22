// Package pool manages a user defined set of resources.
// Based on the work by Fatih Arslan with his pool package.
// https://github.com/fatih/pool/tree/v2.0.0
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

// getQueue returns a safe copy of the resource queue.
func (p *Pool) getQueue() queue {
	var resources queue
	p.mutex.Lock()
	{
		resources = p.resources
	}
	p.mutex.Unlock()

	return resources
}

// Acquire retrieves a resource	from the pool.
func (p *Pool) Acquire() (Resource, error) {
	// Get a safe copy of the queue and check
	// if the queue has been closed.
	resources := p.getQueue()
	if resources == nil {
		return nil, errors.New("Pool has been closed.")
	}

	// The queue is still open so acquire a resource.

	select {
	case resource := <-resources:
		fmt.Println("Acquire:", "Shared Resource")
		if resource == nil {
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
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// If the queue is closed, close the resource before returning.
	if p.resources == nil {
		resource.Close()
		return
	}

	// The queue is still open so release the resource.

	select {
	// Attempt to place the resource back into the queue first.
	// If the queue is full, this will block and the default case
	// will be executed.
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
	// Get a safe copy of the queue and set the queue to nil.
	var resources queue
	p.mutex.Lock()
	{
		resources = p.resources
		p.resources = nil
	}
	p.mutex.Unlock()

	// Is the queue already closed.
	if resources == nil {
		return
	}

	// Close the queue and close all existing resources.
	close(resources)
	for resource := range resources {
		resource.Close()
	}
}
