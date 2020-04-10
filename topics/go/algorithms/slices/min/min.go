package slices

// Min returns the minimum integer in the slice.
func Min(n []int) int {

	// If the length of the slice is 1 then return the integer at index 0.
	if len(n) == 1 {
		return n[0]
	}

	// Create an integer and assign it to the first index of the slice.
	min := n[0]

	// Loop over the slice of integers.
	for _, num := range n {

		// If num is less than min. Assign min to num.
		if num < min {
			min = num
		}
	}

	return min
}
