// Package heap implements a heap interface.
package heap

import (
	"container/heap"
	"errors"
)

var ErrEmptyHeap = errors.New("heap is empty")

// Data represents what is being stored on the queue.
type Data struct {
	Value int
	Index int
}

// Queue represents a list of data.
type Heap struct {
	size     int
	capacity int
	dataHeap *dataHeap
}

// New returns a new Heap with given capacity in which we will store our data.
func New(cap int) (*Heap, error) {
	if cap <= 0 {
		return  nil, errors.New("invalid capacity")
	}

	h := Heap{
		size:     0,
		capacity: cap,
		dataHeap: &dataHeap{},
	}

	return &h, nil
}

// Store is storing the element of type Data in the heap.
func (h *Heap) Store(data *Data) error {
	// If we try to add data to full Heap, we will return an error.
	if h.size >= h.capacity {
		return errors.New("getting out of heap capacity")
	}

	// The complexity of the heap.Push operation is O(log n) where n = h.dataHeap.Len().
	// The element which we will be store in the heap will be added closely to the end of the heap and after what it
	// will be sift up, until it not get in the right position in the heap according it's value.
	heap.Push(h.dataHeap, data)
	h.size++
	return nil
}

// Extract returns the minimum element of type Data from the heap or error error heap is empty.
// Since we have minimum heap implementation. The size of the heap will be decreased.
func (h *Heap) Extract() (*Data, error) {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return nil, ErrEmptyHeap
	}

	// heap.Pop return the element from the heap. the type of element is empty interface.
	element := heap.Pop(h.dataHeap)
	h.size--

	// Since the returned value we got has an empty interface type, let's convert it to our Data type.
	data := element.(*Data)

	return data, nil
}

// Remove removes given element of type Data from the heap. The size value of the Heap will be decreased.
// Returns error when heap is empty.
func (h *Heap) Remove(data *Data) error {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return ErrEmptyHeap
	}

	// Remove given data from the heap
	heap.Remove(h.dataHeap, data.Index)
	h.size--

	return nil
}

// GetRoot returns the element which is stored in the root of the heap and does not change the heap.
// Returns error if heap is empty.
func (h *Heap) GetRoot() (data *Data, err error) {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return nil, ErrEmptyHeap
	}

	return (*h.dataHeap)[0], nil
}

// Size return the current size of the Heap
func (h *Heap) Size() int {
	return h.size
}

// dataHeap is a type which implements a minimum heap interface from the container/heap package which means that data
// with minimum values will be stored at the top of the heap.
type dataHeap []*Data

func (dh *dataHeap) Len() int {
	return len(*dh)
}

func (dh *dataHeap) Less(i, j int) bool {
	return (*dh)[i].Value < (*dh)[j].Value
}

func (dh *dataHeap) Swap(i, j int) {
	(*dh)[i], (*dh)[j] = (*dh)[j], (*dh)[i]
	(*dh)[i].Index = j
	(*dh)[j].Index = i
}

func (dh *dataHeap) Push(x interface{}) {
	item, ok := x.(*Data)
	if !ok {
		return
	}
	item.Index = len(*dh)
	*dh = append(*dh, item)
}

func (dh *dataHeap) Pop() interface{} {
	item := (*dh)[len(*dh)-1]
	*dh = (*dh)[0:len(*dh)-1]
	return item
}
