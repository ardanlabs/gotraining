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

	// Work the front half of the list, moving the largest value we find
	// to the front of the list with each call to move.
	//
	// start: [1 7 7 3     | 1 6 1 4]
	//  move: [1 7 7 <4>   | 1 6 1 <3>]  index: 3  swap:   [3]=3 < [7]=4
	//  move: [1 7 7 4     | 1 6 1 3]    index: 2  noswap: [2]=7 > [5]=6 && [6]=1
	//  move: [1 7 7 4     | 1 6 1 3]    index: 1  noswap: [1]=7 > [3]=4 && [4]=1
	//  move: [<7> <1> 7 4 | 1 6 1 3]    index: 0  swap:   [0]=1 < [1]=7
	//        [7 <4> 7 <1> | 1 6 1 3]    index: 1  swap:   [1]=1 < [3]=4
	//        [7 4 7 <3>   | 1 6 1 <1>]  index: 3  swap:   [3]=1 < [7]=3
	// end:   [7 4 7 3     | 1 6 1 1]
	for index := (len(list) / 2) - 1; index >= 0; index-- {
		list = moveLargest(list, len(list), index)
	}

	// We move the number from the left most index to the right and
	// then cut the size of the list. After we move a number from left
	// to right, we must move the largest number we find once again to the
	// front of the list.
	//
	// start: [7 4 7 3 1 6 1 1]
	//  move: [7 4 <6> 3 1 <1> 1]           [7]
	//  move: [<6> 4 <1> 3 1 1]           [7 7]
	//  move: [<4> <3> 1 1 1]           [6 7 7]
	//  move: [<3> 1 1 1]             [4 6 7 7]
	//  move: [1 1 1]               [3 4 6 7 7]
	//  move: [1 1]               [1 3 4 6 7 7]
	//  done:                 [1 1 1 3 4 6 7 7]
	size := len(list)
	for index := size - 1; index >= 1; index-- {
		list[0], list[index] = list[index], list[0]
		size--
		list = moveLargest(list, size, 0)
	}

	return list
}

// moveLargest starts at the index positions specified in the list and attempts
// to move the largest number it can find to that position in the list.
func moveLargest(list []int, size int, index int) []int {

	// Calculate the index deviation so numbers in the list can be
	// compared and swapped if needed.
	// index 0: cmpIdx1: 1 cmpIdx2:  2   index 5: cmpIdx1: 11 cmpIdx2: 12
	// index 1: cmpIdx1: 3 cmpIdx2:  4   index 6: cmpIdx1: 13 cmpIdx2: 14
	// index 2: cmpIdx1: 5 cmpIdx2:  6   index 7: cmpIdx1: 15 cmpIdx2: 16
	// index 3: cmpIdx1: 7 cmpIdx2:  8   index 8: cmpIdx1: 17 cmpIdx2: 19
	// index 4: cmpIdx1: 9 cmpIdx2: 10   index 9: cmpIdx1: 19 cmpIdx2: 20
	cmpIdx1, cmpIdx2 := 2*index+1, 2*index+2

	// Save the specified index as the index with the current largest value.
	largestValueIdx := index

	// Check if the value at the first deviation index is greater than
	// the value at the current largest index. If so, save that
	// index position.
	if cmpIdx1 < size && list[cmpIdx1] > list[largestValueIdx] {
		largestValueIdx = cmpIdx1
	}

	// Check the second deviation index is within bounds and is greater
	// than the value at the current largest index. If so, save that
	// index position.
	if cmpIdx2 < size && list[cmpIdx2] > list[largestValueIdx] {
		largestValueIdx = cmpIdx2
	}

	// If we found a larger value than the value at the specified index, swap
	// those numbers and then recurse to find more numbers to swap from that
	// point in the list.
	if largestValueIdx != index {
		list[index], list[largestValueIdx] = list[largestValueIdx], list[index]
		list = moveLargest(list, size, largestValueIdx)
	}

	return list
}
