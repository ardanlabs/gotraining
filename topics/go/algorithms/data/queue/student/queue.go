// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package queue implements of a circular queue.
package queue

// Data represents what is being stored on the queue.
type Data struct {
	Name string
}

// Queue represents a list of data.
type Queue struct {
	Count int
	data  []*Data
	front int
	end   int
}

// New returns a queue with a set capacity.
func New(cap int) (*Queue, error) {
	return nil, nil
}

// Enqueue inserts data into the queue if there
// is available capacity.
func (q *Queue) Enqueue(data *Data) error {
	return nil
}

// Dequeue removes data into the queue if data exists.
func (q *Queue) Dequeue() (*Data, error) {
	return nil, nil
}

// Operate accepts a function that takes data and calls
// the specified function for every piece of data found.
func (q *Queue) Operate(f func(d *Data) error) error {
	return nil
}
