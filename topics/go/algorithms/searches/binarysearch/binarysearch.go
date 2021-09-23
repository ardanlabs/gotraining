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
func binarySearchRecursive(List []int, Target int, low int, high int) (int, error) {

	// 	find the middle index of array
	mid := (low + high) / 2

	// 	If low is greater than high it means the value not found,
	//	I returned -1 and an error to show user value not found
	for low > high {
		return -1, fmt.Errorf("sorry value not found")
	}

	// 	check is List[mid] value equal with Target value than return the middle value
	if List[mid] == Target {
		return mid, nil
	}

	// 	If List[mid] is bigger than Target
	//	I returned the binarySearchRecursive function with changing the high value to mid minus by one
	if List[mid] > Target {
		return binarySearchRecursive(List, Target, low, mid-1)
	}

	// 	If List[mid] is bigger than Target
	//	I returned the binarySearchRecursive function with changing the low value to mid plus by one
	if List[mid] < Target {
		return binarySearchRecursive(List, Target, mid+1, high)
	}

	// return -1 if user not found
	return -1, fmt.Errorf("sorry value not found")
}
