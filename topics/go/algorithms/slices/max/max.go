package slices

import (
	"fmt"
)

// Max returns the maximum integer in the slice.
func Max(n []int) (int, error) {
	if len(n) == 0 {
		return 0, fmt.Errorf("slice %#v has no elements", n)
	}

	max := n[0]

	// Loop over the slice of integers.
	for _, num := range n[1:] {

		// If num is greater than max, assign max to num.
		if num > max {
			max = num
		}
	}

	return max, nil
}
