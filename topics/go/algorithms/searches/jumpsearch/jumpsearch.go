// Package jumpsearch provides an example of a jump search implementation.
package jumpsearch

import (
	"math"
)

// jumpSearch takes a sorted list of numbers and uses the
// `binarysearch` and `linearsearch` algorithms to find the target.
// the worst cast | O(√n)
// the best cast  | O(1)
func jumpSearch(sortedList []int, target int) int {
	var index int

	// Calculate jump value of the list length.
	jump := int(math.Sqrt(float64(len(sortedList))))

	// If list is empty it will return -1.
	if len(sortedList) <= 0 {
		return -1
	}

	// Loop until we find the target or break the loop if target is smaller
	// than the sortedList value.
loop:
	for index <= len(sortedList)-1 {
		switch {

		// Check if we found the target.
		case sortedList[index] == target:
			return index

		// Break the loop if target is smaller than the sortedList value.
		case sortedList[index] > target:
			break loop

		// Continue adding the jump value to the index until the target found
		// or the loop breaking.
		default:
			index = index + jump
		}
	}

	// Add previous jump to the index value.
	previous := index - jump

	// Check the index is greater than the length of the list.
	if index > len(sortedList)-1 {
		index = len(sortedList) - 1
	}

	// Loop until we find the target or searched the list.
	for sortedList[index] >= target {
		switch {

		// Check if we found the target.
		case sortedList[index] == target:
			return index

		// Check the index with the previous index value.
		case index == previous:
			return -1

		default:
			index--
		}
	}

	return -1
}
