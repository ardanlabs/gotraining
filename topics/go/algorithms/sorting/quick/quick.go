// Package quicksort implementation of Quick sort algorithm in Go.
package quicksort

// quickSort is an in-place sorting algorithm. It takes a random list of numbers,
// and uses the `recursive` process to divides it into partitions then sorts those.
// - Time complexity O(nlog n)
// - Space complexity O(log n)
func quickSort(randomList []int, leftIdx, rightIdx int) []int {
	switch {
	case leftIdx > rightIdx:
		return randomList

	// Divides array into two partitions.
	case leftIdx < rightIdx:
		randomList, pivotIdx := partition(randomList, leftIdx, rightIdx)

		quickSort(randomList, leftIdx, pivotIdx-1)
		quickSort(randomList, pivotIdx+1, rightIdx)
	}

	return randomList
}

// partition it takes a portion of an array then sort it.
func partition(randomList []int, leftIdx, rightIdx int) ([]int, int) {
	pivot := randomList[rightIdx]

	for smallest := leftIdx; smallest < rightIdx; smallest++ {
		if randomList[smallest] < pivot {
			randomList[smallest], randomList[leftIdx] = randomList[leftIdx], randomList[smallest]
			leftIdx++
		}
	}

	randomList[leftIdx], randomList[rightIdx] = randomList[rightIdx], randomList[leftIdx]

	return randomList, leftIdx
}
