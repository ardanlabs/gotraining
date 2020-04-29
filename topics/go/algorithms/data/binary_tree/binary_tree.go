// Package binary_tree implements a binary_tree interface, and add the ability to store values
// of type Data in the minimum BinaryTree.
package binary_tree

import (
	"container/heap"
	"errors"
)

// Data represents what is being stored on the binary_tree.
type Data struct {
	Value int
	Index int
}

// BinaryTree represents a list of data.
type BinaryTree struct {
	size     int
	capacity int
	dataHeap *dataHeap
}

// New returns a new BinaryTree with a given capacity in which we will store our data.
func New(cap int) (*BinaryTree, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}

	b := BinaryTree{
		capacity: cap,
		dataHeap: &dataHeap{},
	}

	return &b, nil
}

// Store is storing the element of type Data in the binary_tree.
func (b *BinaryTree) Store(data *Data) error {

	// If we try to add data to full BinaryTree, we will return an error.
	if b.size >= b.capacity {
		return errors.New("out of binary_tree capacity")
	}

	// The element which we will be store in the binary_tree will be added closely to
	// the end of the binary_tree and after what it will be sift up, until it not get
	// in the right position in the binary_tree according it's value.
	heap.Push(b.dataHeap, data)
	b.size++
	return nil
}

// Extract returns the minimum element of type Data from the binary_tree or error then
// binary_tree is empty. Since we have minimum binary_tree implementation.
// The size of the binary_tree will be decreased.
func (b *BinaryTree) Extract() (*Data, error) {

	// Check that binary_tree is not empty
	if len(*b.dataHeap) == 0 {
		return nil, errors.New("binary_tree is empty")
	}

	// Pop Return the element from the binary_tree.
	element := heap.Pop(b.dataHeap)
	b.size--

	// Since the returned value we got has an empty interface type, let's
	// convert it to our Data type.
	data, ok := element.(*Data)
	if !ok {
		return nil, errors.New("wrong data type")
	}
	return data, nil
}

// Remove removes given element of type Data from the binary_tree or return error when
// binary_tree is empty. The size value of the BinaryTree will be decreased.
func (b *BinaryTree) Remove(data *Data) error {

	// Check that binary_tree is not empty
	if len(*b.dataHeap) == 0 {
		return errors.New("binary_tree is empty")
	}

	// Remove given data from the binary_tree
	heap.Remove(b.dataHeap, data.Index)
	b.size--

	return nil
}

// GetRoot returns the element which is stored in the root of the binary_tree and does
// not change the binary_tree. Returns error if binary_tree is empty.
func (b *BinaryTree) GetRoot() (data *Data, err error) {

	// Check that binary_tree is not empty
	if len(*b.dataHeap) == 0 {
		return nil, errors.New("binary_tree is empty")
	}
	d := (*b.dataHeap)[0]
	return d, nil
}

// Size return the current size of the BinaryTree
func (b *BinaryTree) Size() int {
	return b.size
}

// dataHeap is a type which implements a minimum binary_tree interface from the
// container/binary_tree package which means that data with minimum values will be
// stored at the top of the binary_tree.
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
