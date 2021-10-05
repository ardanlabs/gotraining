// Package selectionsort implement of Selection Sort algorithm in Go.
package selectionsort

// selectionSortIterative takes a random list of numbers and uses the
// `iterative` process to sort it and return the sorted list.
func selectionSortIterative(randomList []int) []int {

	// Loop through the list until it is sorted.
	for leftIdx := range randomList {
		index := leftIdx

		// Look for the smallest number in the list starting from leftIdx. If a
		// number of located, capture the index position of that number.
		for smallestIdx := leftIdx; smallestIdx < len(randomList); smallestIdx++ {
			if randomList[smallestIdx] < randomList[index] {
				index = smallestIdx
			}
		}

		// Swap the number from the leftIdx with the smallest number found.
		randomList[leftIdx], randomList[index] = randomList[index], randomList[leftIdx]
	}

	return randomList
}
