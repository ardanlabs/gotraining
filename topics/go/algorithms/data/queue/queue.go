// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package queue implements of a circular queue.
package queue

import (
	"errors"
)

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
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}

	q := Queue{
		front: 0,
		end:   0,
		data:  make([]*Data, cap),
	}
	return &q, nil
}

// Enqueue inserts data into the queue if there
// is available capacity.
func (q *Queue) Enqueue(data *Data) error {

	// If the front of the queue is right behind the end or
	// if the front is at the end of the capacity and the end
	// is at the beginning of the capacity, the queue is full.
	//  F  E  - Enqueue (Full) |  E        F - Enqueue (Full)
	// [A][B][C]               | [A][B][C]
	if q.front+1 == q.end ||
		q.front == len(q.data) && q.end == 0 {
		return errors.New("queue at capacity")
	}

	switch {
	case q.front == len(q.data):

		// If we are at the end of the capacity, then
		// circle back to the beginning of the capacity by
		// moving the front pointer to the beginning.
		q.front = 0
		q.data[q.front] = data
	default:

		// Add the data to the current front position
		// and then move the front pointer.
		q.data[q.front] = data
		q.front++
	}

	q.Count++

	return nil
}

// Dequeue removes data into the queue if data exists.
func (q *Queue) Dequeue() (*Data, error) {

	// If the front and end are the same, the
	// queue is empty
	//  EF - (Empty)
	// [  ][ ][ ]
	if q.front == q.end {
		return nil, errors.New("queue is empty")
	}

	var data *Data
	switch {
	case q.end == len(q.data):

		// If we are at the end of the capacity, then
		// circle back to the beginning of the capacity by
		// moving the end pointer to the beginning.
		q.end = 0
		data = q.data[q.end]
	default:

		// Remove the data from the current end position
		// and then move the end pointer.
		data = q.data[q.end]
		q.end++
	}

	q.Count--

	return data, nil
}

// Operate accepts a function that takes data and calls
// the specified function for every piece of data found.
func (q *Queue) Operate(f func(d *Data) error) error {
	end := q.end
	for {
		if end == q.front {
			break
		}

		if end == len(q.data) {
			end = 0
		}

		if err := f(q.data[end]); err != nil {
			return err
		}

		end++
	}
	return nil
}
