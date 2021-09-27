// Package binarysearch provides an example of a binary search implementation.
package binarysearch

import "fmt"

// binarySearchIterative takes a sorted list of numbers and uses the
// `iterative` process to find the target. The function returns the
// index postion of where the target is found.
// - the worst case of this algorithm is O(logn)
// - the best case of this algorithm is O(1)
func binarySearchIterative(sortedList []int, target int) (int, error) {
	var leftIdx int
	rightIdx := len(sortedList) - 1

	// Loop until we find the target or searched the list.
	for leftIdx <= rightIdx {

		// Calculate the middle index of the list.
		mid := (leftIdx + rightIdx) / 2

		// Capture the value to check.
		value := sortedList[mid]

		switch {

		// Check if we found the target.
		case value == target:
			return mid, nil

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case value > target:
			rightIdx = mid - 1

		// If the value is less than the target, cut the list
		// by moving the leftIdx into the list.
		case value < target:
			leftIdx = mid + 1
		}
	}

	return -1, fmt.Errorf("target not found")
}

// binarySearchRecursive takes the list of the sorted numbers and check it
// with `recursive` process to find the value and return the index of array or
// return the error if the value not found.
func binarySearchRecursive(sortedList []int, target int, leftIdx int, rightIdx int) (int, error) {

	// Calculate the middle index of the list.
	midIdx := (leftIdx + rightIdx) / 2

	// Check until leftIdx is smaller or equal with rightIdx.
	if leftIdx <= rightIdx {

		switch {

		// Check if we found the target.
		case sortedList[midIdx] == target:
			return midIdx, nil

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case sortedList[midIdx] > target:
			return binarySearchRecursive(sortedList, target, leftIdx, midIdx-1)

		// If the value is less than the target, cut the list
		// by moving the leftIdx into the list.
		case sortedList[midIdx] < target:
			return binarySearchRecursive(sortedList, target, midIdx+1, rightIdx)
		}
	}

	return -1, fmt.Errorf("target not found")
}
