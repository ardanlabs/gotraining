// Package binary_tree implements a binary_tree interface, and add the ability to store values
// of type Data in the minimum BinaryTree.
package binary_tree

import (
	"errors"
)

// Data represents what is being stored on the binary_tree.
type Data struct {
	Value int
	Index int
}

// BinaryTree represents a list of data.
type BinaryTree struct {
	capacity int
	data []Data
}

// New returns a new BinaryTree with in which we will store our data.
func New(cap int) (*BinaryTree, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}
	h := BinaryTree{
		capacity: cap,
	}

	return &h, nil
}

// Store is storing the element of type Data in the binary_tree.
func (h *BinaryTree) Store(d Data) error {
	if len(h.data) > h.capacity {
		return errors.New("out of binary_tree capacity")
	}
	// The element which we will be store in the binary_tree will be added closely to
	// the end of the binary_tree and after what it will be sift up, until it not get
	// in the rightLeaf position in the binary_tree according it's value.
	d.Index = len(h.data)
	h.data = append(h.data, d)
	h.up(len(h.data)-1)
	return nil
}

// Extract returns the minimum element of type Data from the binary_tree or error then
// binary_tree is empty. Since we have minimum binary_tree implementation.
// The size of the binary_tree will be decreased.
func (h *BinaryTree) Extract() (*Data, error) {

	// Check that binary_tree is not empty
	if len(h.data) == 0 {
		return nil, errors.New("binary_tree is empty")
	}

	// get element from the root of the binary_tree, save it to result variable,
	// when change root element of the binary_tree with binary_tree leaf and sift it down to
	// fix the binary_tree
	result := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.down(0)
	h.data = h.data[:len(h.data)-1]

	return &result, nil
}

// Remove removes given element of type Data from the binary_tree or return error when
// binary_tree is empty. The size value of the BinaryTree will be decreased.
func (h *BinaryTree) Remove(d Data) error {

	// Check that binary_tree is not empty
	if len(h.data) == 0 {
		return errors.New("binary_tree is empty")
	}

	h.data[d.Index].Value = h.data[0].Value - 1
	h.up(d.Index)
	h.Extract()
	return nil
}

// GetRoot returns the element which is stored in the root of the binary_tree and does
// not change the binary_tree. Returns error if binary_tree is empty.
func (h *BinaryTree) GetRoot() (data *Data, err error) {

	// Check that binary_tree is not empty
	if len(h.data) == 0 {
		return nil, errors.New("binary_tree is empty")
	}

	return &h.data[0], nil
}

// Size return the current size of the BinaryTree
func (h *BinaryTree) Size() int {
	return len(h.data)
}

// down is sift element with given index down to the binary_tree until the element
// has not beign set to appropriate place
func (h *BinaryTree)down(i int) {
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

// up is sift up element with given index up on the binary_tree until the element
// has not being set to appropriate place
func (h *BinaryTree) up(j int) {
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