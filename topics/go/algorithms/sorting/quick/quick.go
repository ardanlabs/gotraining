// Package quicksort implementation of Quick sort algorithm in Go.
package quicksort

// quickSort is an in-place sorting algorithm. It takes a random list of numbers,
// and uses the `recursive` process to divides it into partitions then sorts those.
// - Time complexity O(nlog n)
// - Space complexity O(log n)
func quickSort(randomList []int, leftIdx int, rightIdx int) {
	switch {
	case len(randomList) < 1:
		return

	// Divides array into two partitions.
	case leftIdx < rightIdx:
		pivotIdx := partition(randomList, leftIdx, rightIdx)

		quickSort(randomList, leftIdx, pivotIdx-1)
		quickSort(randomList, pivotIdx+1, leftIdx)
	}
}

// partition it takes a portion of an array then sort it.
func partition(randomList []int, leftIdx int, rightIdx int) int {
	pivot := randomList[rightIdx]
	index := leftIdx - 1

	for smallest := leftIdx; smallest < rightIdx; smallest++ {
		if randomList[smallest] < pivot {
			index++
			randomList[index], randomList[rightIdx] = randomList[rightIdx], randomList[index]
		}
	}

	randomList[index+1], randomList[rightIdx] = randomList[rightIdx], randomList[index+1]

	return index + 1
}
