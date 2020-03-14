package strings

func ReverseString(str string) string {

	// Here we create a slice of codePoints.
	codePoints := []rune(str)

	// Here we create int that will be a pointer to the front of the codePoints.
	var beg int

	// Here  we create int that will be a pointer to the end of the codePoints.
	end := len(codePoints) - 1

	// While there are still code points to check.
	for beg < end {

		// Swap the code points by:

		// 1. Create a code point with the value at index beg.
		c := codePoints[beg]

		// 2. Swap the code point at index beg with the code point at index end.
		codePoints[beg] = codePoints[end]

		// 3. Swap the code point at index end with c.
		codePoints[end] = c

		beg = beg + 1

		end = end - 1
	}

	// Here we return a new string in reverse order.
	return string(codePoints)
}
