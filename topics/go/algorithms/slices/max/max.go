package slices

// Max returns the maximum integer in the slice.
func Max(n []int) int {

	// If the length of the slice is 1 then return the integer at index 0.
	if len(n) == 1 {
		return n[0]
	}

	// Create an integer and assign it to the first index of the slice.
	max := n[0]

	// Loop over the slice of integers.
	for _, num := range n {

		// If num is greater than max. Assign max to num.
		if num > max {
			max = num
		}
	}
	return max
}
