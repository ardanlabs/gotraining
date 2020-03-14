package strings

func ReverseString(str string) string {

	// Here we create a slice of r.
	r := []rune(str)

	// Here we create int that will be a pointer to the front of the r.
	var beg int

	// Here  we create int that will be a pointer to the end of the r.
	end := len(r) - 1

	// While there are still code points to check.
	for beg < end {

		// Swap the code point.
		c := r[beg]

		r[beg] = r[end]

		r[end] = c

		beg, end = beg+1, end-1

	}

	// Here we return a new string in reverse order.
	return string(r)
}
