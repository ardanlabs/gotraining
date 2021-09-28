// Package linearsearch provides an example of a linear search implementation.
package linearsearch

// linearSearchIterative takes a sorted/random list of numbers
// and uses the `iterative` process to check index by index to
// find the target.
// - the worst case of this algorithm is O(n).
// - the best case of this algorithm is O(1).
func linearSearchIterative(list []int, target int) int {
	if len(list) <= 0 {
		return -1
	}

	for i := range list {
		if list[i] == target {
			return i
		}
	}

	return -1
}

// linearSearchRecursive takes a sorted/random list of numbers
// and uses the `recursive` process to check index by index to
// find the target.
func linearSearchRecursive(list []int, target int, index int) int {
	switch {
	case len(list) <= index:
		return -1

	case list[index] == target:
		return index

	case len(list) >= index:
		index++
		return linearSearchRecursive(list, target, index)
	}

	return -1
}

// doubleLinearSearchIterative takes a sorted/random list of numbers
// and uses the `iterative` process to check index by index to find
// the target in both left and right index.
func doubleLinearSearchIterative(list []int, target int) int {
	var leftIdx int
	rightIdx := len(list) - 1

	if len(list) <= 0 {
		return -1
	}

	// Continue loop until leftIdx is smaller or equal to rightIdx value.
	for leftIdx <= rightIdx {
		switch {
		case list[leftIdx] == target:
			return leftIdx

		case list[rightIdx] == target:
			return rightIdx

		default:
			leftIdx++
			rightIdx--
		}
	}

	return -1
}

// doubleLinearSearchRecursive takes a sorted/random list of numbers
// and uses the `recursive` process to check index by index to find
// the target in both left and right index.
func doubleLinearSearchRecursive(list []int, target int, leftIdx int, rightIdx int) int {
	if len(list) > 0 {
		switch {
		case list[leftIdx] == target:
			return leftIdx

		case list[rightIdx] == target:
			return rightIdx

		case leftIdx <= rightIdx:
			leftIdx++
			rightIdx--
			return doubleLinearSearchRecursive(list, target, leftIdx, rightIdx)
		}
	}

	return -1
}
