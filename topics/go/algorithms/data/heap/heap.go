// Package heap implements a heap interface, and add the ability to store values
// of type Data in the minimum Heap.
package heap

import (
	"errors"
)

// Data represents what is being stored on the heap.
type Data struct {
	Value int
	Index int
}

// Heap represents a list of data.
type Heap struct {
	capacity int
	data []Data
}

// New returns a new Heap with in which we will store our data.
func New(cap int) (*Heap, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}
	h := Heap{
		capacity: cap,
	}

	return &h, nil
}

// Store is storing the element of type Data in the heap.
func (h *Heap) Store(d Data) error {
	if len(h.data) > h.capacity {
		return errors.New("out of heap capacity")
	}
	// The element which we will be store in the heap will be added closely to
	// the end of the heap and after what it will be sift up, until it not get
	// in the rightLeaf position in the heap according it's value.
	d.Index = len(h.data)
	h.data = append(h.data, d)
	h.up(len(h.data)-1)
	return nil
}

// Extract returns the minimum element of type Data from the heap or error then
// heap is empty. Since we have minimum heap implementation.
// The size of the heap will be decreased.
func (h *Heap) Extract() (*Data, error) {

	// Check that heap is not empty
	if len(h.data) == 0 {
		return nil, errors.New("heap is empty")
	}

	// get element from the root of the heap, save it to result variable,
	// when change root element of the heap with heap leaf and sift it down to
	// fix the heap
	result := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.down(0)
	h.data = h.data[:len(h.data)-1]

	return &result, nil
}

// Remove removes given element of type Data from the heap or return error when
// heap is empty. The size value of the Heap will be decreased.
func (h *Heap) Remove(d Data) error {

	// Check that heap is not empty
	if len(h.data) == 0 {
		return errors.New("heap is empty")
	}

	h.data[d.Index].Value = h.data[0].Value - 1
	h.up(d.Index)
	h.Extract()
	return nil
}

// GetRoot returns the element which is stored in the root of the heap and does
// not change the heap. Returns error if heap is empty.
func (h *Heap) GetRoot() (data *Data, err error) {

	// Check that heap is not empty
	if len(h.data) == 0 {
		return nil, errors.New("heap is empty")
	}

	return &h.data[0], nil
}

// Size return the current size of the Heap
func (h *Heap) Size() int {
	return len(h.data)
}

// down is sift element with given index down to the heap until the element
// has not beign set to appropriate place
func (h *Heap)down(i int) {
	maxIndex := i

	ll := leftLeaf(i)
	if ll < 0 || ll >= len(h.data) - 1{
		return
	}
	if h.data[ll].Value < h.data[maxIndex].Value {
		maxIndex = ll
	}

	rl := rightLeaf(i)
	if rl < 0 {
		return
	}

	if h.data[rl].Value < h.data[maxIndex].Value {
		maxIndex = rl
	}

	if h.data[i] != h.data[maxIndex] {

		h.data[i].Index = h.data[maxIndex].Index
		h.data[i], h.data[maxIndex] = h.data[maxIndex], h.data[i]
		h.data[i].Index--

		h.down(maxIndex)
	}
}

// up is sift up element with given index up on the heap unitl the element
// has not beign set to appropriate place
func (h *Heap) up(j int) {
	for {
		i := parent(j)
		if h.data[i] == h.data[j] || h.data[j].Value > h.data[i].Value {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

// parent Returns the index of the parent element
func parent(i int) int {
	return (i - 1) / 2
}

// leftLeaf Returns the index of the element on the leftLeaf
func leftLeaf(i int) int {
	return 2*i
}

// rightLeaf Returns the index of the element on the rightLeaf
func rightLeaf(i int) int {
	return 2*i+1
}