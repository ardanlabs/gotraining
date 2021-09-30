// Package selectionsort implement of Selection Sort algorithm in Go.
package selectionsort

// selectionSortIterative takes a random list of numbers and uses the
// `iterative` process to sort it and return the sorted list.
func selectionSortIterative(randomList []int) []int {

	// Loop until the random list will be sorted.
	for i := range randomList {
		index := i

		// Continue the loop until to found the smallest number in the list.
		for j := i; j < len(randomList); j++ {
			if randomList[j] < randomList[i] {
				index = j
			}
		}

		// We will swap values based on indexes.
		randomList[i], randomList[index] = randomList[index], randomList[i]
	}

	return randomList
}
