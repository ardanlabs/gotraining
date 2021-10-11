// Package heap implement the heapsort algorithm in Go. Heapsort can be thought
// of as an improved selection sort: like selection sort, heapsort divides its
// input into a sorted and an unsorted region, and it iteratively shrinks the
// unsorted region by extracting the largest element from it and inserting it
// into the sorted region. Unlike selection sort, heapsort does not waste time
// with a linear-time scan of the unsorted region; rather, heap sort maintains
// the unsorted region in a heap data structure to more quickly find the largest
// element in each step.
package heap

// HeapSort takes a random list of numbers and returns the sorted list.
func HeapSort(list []int) []int {

	// NEED COMMENT HERE FOR WHAT THIS IS DOING.
	for index := len(list) / 2; index >= 0; index-- {
		list = sort(list, len(list), index)
	}

	// NEED COMMENT HERE FOR WHAT THIS IS DOING.
	size := len(list)
	for index := len(list) - 1; index >= 1; index-- {
		size--
		list[0], list[index] = list[index], list[0]
		list = sort(list, size, 0)
	}

	return list
}

// sort take a list, size, and index position to sort from.
func sort(list []int, size int, index int) []int {

	// leftIdx is for the left child index of heap.
	// rightIdx is for the right child index of heap.
	leftIdx, rightIdx := 2*index+1, 2*index+2
	largeIdx := index

	if leftIdx < size && list[leftIdx] > list[largeIdx] {
		largeIdx = leftIdx
	}

	if rightIdx < size && list[rightIdx] > list[largeIdx] {
		largeIdx = rightIdx
	}

	if largeIdx != index {
		list[index], list[largeIdx] = list[largeIdx], list[index]
		list = sort(list, size, largeIdx)
	}

	return list
}
