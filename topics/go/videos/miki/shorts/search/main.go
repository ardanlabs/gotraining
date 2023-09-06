package main

import (
	"fmt"
	"slices"
)

var (
	grades  = "FDCBA"
	cutoffs = []int{60, 70, 80, 90}
)

func score(n int) byte {
	// See also sort.SearchInts.
	i, ok := slices.BinarySearch(cutoffs, n)
	if ok {
		i++
	}
	return grades[i]
}

func main() {
	scores := []int{33, 99, 77, 70, 89, 90, 100}
	for _, n := range scores {
		fmt.Printf("%d -> %c\n", n, score(n))
	}
}
