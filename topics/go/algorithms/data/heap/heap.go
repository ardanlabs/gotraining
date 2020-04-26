// Package heap implements a heap interface, and add the ability to store values
// of type Data in the minimum Heap.
package heap

import (
	"container/heap"
	"errors"
)

// Data represents what is being stored on the heap.
type Data struct {
	Value int
	Index int
}

// Heap represents a list of data.
type Heap struct {
	size     int
	capacity int
	dataHeap *dataHeap
}

// New returns a new Heap with a given capacity in which we will store our data.
func New(cap int) (*Heap, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}

	h := Heap{
		capacity: cap,
		dataHeap: &dataHeap{},
	}

	return &h, nil
}

// Store is storing the element of type Data in the heap.
func (h *Heap) Store(data *Data) error {

	// If we try to add data to full Heap, we will return an error.
	if h.size >= h.capacity {
		return errors.New("out of heap capacity")
	}

	// The element which we will be store in the heap will be added closely to
	// the end of the heap and after what it will be sift up, until it not get
	// in the right position in the heap according it's value.
	heap.Push(h.dataHeap, data)
	h.size++
	return nil
}

// Extract returns the minimum element of type Data from the heap or error then
// heap is empty. Since we have minimum heap implementation.
// The size of the heap will be decreased.
func (h *Heap) Extract() (*Data, error) {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return nil, errors.New("heap is empty")
	}

	// Pop Return the element from the heap.
	element := heap.Pop(h.dataHeap)
	h.size--

	// Since the returned value we got has an empty interface type, let's
	// convert it to our Data type.
	data, ok := element.(*Data)
	if !ok {
		return nil, errors.New("wrong data type")
	}
	return data, nil
}

// Remove removes given element of type Data from the heap or return error when
// heap is empty. The size value of the Heap will be decreased.
func (h *Heap) Remove(data *Data) error {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return errors.New("heap is empty")
	}

	// Remove given data from the heap
	heap.Remove(h.dataHeap, data.Index)
	h.size--

	return nil
}

// GetRoot returns the element which is stored in the root of the heap and does
// not change the heap. Returns error if heap is empty.
func (h *Heap) GetRoot() (data *Data, err error) {

	// Check that heap is not empty
	if len(*h.dataHeap) == 0 {
		return nil, errors.New("heap is empty")
	}
	d := (*h.dataHeap)[0]
	return d, nil
}

// Size return the current size of the Heap
func (h *Heap) Size() int {
	return h.size
}

// dataHeap is a type which implements a minimum heap interface from the
// container/heap package which means that data with minimum values will be
// stored at the top of the heap.
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

func (dh *dataHeap) Push(v interface{}) {
	item, ok := v.(*Data)
	if !ok {
		return
	}

	item.Index = len(*dh)
	*dh = append(*dh, item)
}

func (dh *dataHeap) Pop() interface{} {
	item := (*dh)[len(*dh)-1]
	*dh = (*dh)[0 : len(*dh)-1]
	return item
}
