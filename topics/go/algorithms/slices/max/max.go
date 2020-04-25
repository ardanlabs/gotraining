package max

import "fmt"

// Max returns the maximum integer in the slice.
func Max(n []int) (int, error) {

	// First check there are numbers in the collection.
	if len(n) == 0 {
		return 0, fmt.Errorf("slice %#v has no elements", n)
	}

	// If the length of the slice is 1 then return the
	// integer at index 0.
	if len(n) == 1 {
		return n[0], nil
	}

	// Save the first value as current max and then loop over
	// the slice of integers looking for a larger number.
	max := n[0]
	for _, num := range n[1:] {

		// If num is greater than max, assign max to num.
		if num > max {
			max = num
		}
	}

	return max, nil
}
