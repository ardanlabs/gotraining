// Package interpolationsearch provides an example of an interpolation search implementation.
package interpolationsearch

// interpolationSearchIterative this algorithm is improved the binarysearch with
// the `iterative` method, for finding the middle index, it has another position.
// It means we will not looking for the middle of the index anymore we calculate
// the index in another way.
func interpolationSearchIterative(sortedList []int, target int) int {
	var leftIdx int
	rightIdx := len(sortedList) - 1

	// Loop until we find the target or searched the list.
	for leftIdx <= rightIdx &&
		target >= sortedList[leftIdx] &&
		target <= sortedList[rightIdx] &&
		len(sortedList) > 0 {

		// Calculate the position index of the list.
		a := int(float64(rightIdx-leftIdx) / float64(sortedList[rightIdx]-sortedList[leftIdx]))
		b := target - sortedList[leftIdx]
		positionIdx := leftIdx + a*b

		// Capture the value to check.
		value := sortedList[positionIdx]

		switch {

		// Check if we found the target.
		case value == target:
			return positionIdx

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case value > target:
			rightIdx = positionIdx - 1

		// If the value is less than the target, cut the list
		// by moving the leftIdx into the list.
		case value < target:
			leftIdx = positionIdx - 1
		}
	}

	return -1
}

// interpolationSearchRecursive this algorithm is improved the binarysearch algorithm with the `recursive` method,
// for finding the middle index, it has another position.
// It means we will not looking for the middle of the index anymore. We calculate the index in another way.
func interpolationSearchRecursive(sortedList []int, target int, leftIdx int, rightIdx int) int {

	// Check until we find the target or searched the list.
	if leftIdx <= rightIdx &&
		target >= sortedList[leftIdx] &&
		target <= sortedList[rightIdx] &&
		len(sortedList) > 0 {

		// Calculate the position index of the list.
		a := int(float64(rightIdx-leftIdx) / float64(sortedList[rightIdx]-sortedList[leftIdx]))
		b := target - sortedList[leftIdx]
		positionIdx := leftIdx + a*b

		// Capture the value to check.
		value := sortedList[positionIdx]

		switch {

		// Check if we found the target.
		case value == target:
			return positionIdx

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case value > target:
			return interpolationSearchRecursive(sortedList, target, leftIdx, positionIdx-1)

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case value < target:
			return interpolationSearchRecursive(sortedList, target, positionIdx+1, rightIdx)
		}

	}

	return -1
}
