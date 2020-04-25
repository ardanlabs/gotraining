package min

import "fmt"

// Min returns the minimum integer in the slice.
func Min(n []int) (int, error) {

	// First check there are numbers in the collection.
	if len(n) == 0 {
		return 0, fmt.Errorf("slice %#v has no elements", n)
	}

	// If the length of the slice is 1 then return the
	// integer at index 0.
	if len(n) == 1 {
		return n[0], nil
	}

	// Save the first value as current min and then loop over
	// the slice of integers looking for a smaller number.
	min := n[0]
	for _, num := range n[1:] {

		// If num is less than min. Assign min to num.
		if num < min {
			min = num
		}
	}

	return min, nil
}
