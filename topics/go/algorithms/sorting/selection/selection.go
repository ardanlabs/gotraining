// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implementation of Selection Sort in Go.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(1e2)
	fmt.Println("Before:", numbers)
	selectionSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}

// Find smallest value from an unsorted slice.
// Put the smallest value in the leftmost position by swapping value.
// Move the unsorted slice starting position to the right by one.
// Iterate process.
func selectionSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		idxMin := min(numbers, i)
		if idxMin != i {
			numbers[i], numbers[idxMin] = numbers[idxMin], numbers[i]
		}
	}
}

// Find the lowest value index
func min(numbers []int, i int) int {
	idxMin := i
	for i++; i < len(numbers); i++ {
		if numbers[i] < numbers[idxMin] {
			idxMin = i
		}
	}
	return idxMin
}
