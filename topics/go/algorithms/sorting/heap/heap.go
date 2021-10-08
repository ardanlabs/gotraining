// Package heapsort implement of Heap Sort algorithm in Go.
package heapsort

type heapList struct {
	list []int
	size int
}

// heapSort takes a random list of numbers and returns the sorted list.
func HeapSort(list []int) []int {
	heap := initial(list)

	// Loop through the list until it is sorted, after initial the list.
	for index := len(heap.list) - 1; index >= 1; index-- {
		heap.size--
		heap.list[0], heap.list[index] = heap.list[index], heap.list[0]
		heap.heapify(0)
	}

	return heap.list
}

// initial is take a list of array, and it will add them to the heapList,
// and pass the index to the heapify function.
func initial(list []int) heapList {
	heap := heapList{
		list: list,
		size: len(list),
	}

	for index := len(list) / 2; index >= 0; index-- {
		heap.heapify(index)
	}

	return heap
}

// heapify take the index of array and base on it will sort the array.
func (heap heapList) heapify(index int) {

	// leftIdx is for the left child index of heap.
	// rightIdx is for the right child index of heap.
	leftIdx, rightIdx := 2*index+1, 2*index+2
	largeIdx := index

	if leftIdx < heap.length() && heap.list[leftIdx] > heap.list[largeIdx] {
		largeIdx = leftIdx
	}

	if rightIdx < heap.length() && heap.list[rightIdx] > heap.list[largeIdx] {
		largeIdx = rightIdx
	}

	if largeIdx != index {
		heap.list[index], heap.list[largeIdx] = heap.list[largeIdx], heap.list[index]
		heap.heapify(largeIdx)
	}
}

// length is return the length of the heap array.
func (heap heapList) length() int {
	return heap.size
}
