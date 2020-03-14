package strings

func ReverseString(str string) string {

	// Here we create a slice of runes.
	runes := []rune(str)

	// Here we create int that will be a pointer to the front of the runes.
	var beg int

	// Here  we create int that will be a pointer to the end of the runes.
	end := len(runes) - 1

	// While there are still code points to check.
	for beg < end {

		// Swap the code point.
		r := runes[beg]

		runes[beg] = runes[end]

		runes[end] = r

		beg, end = beg+1, end-1

	}

	// Here we return a new string in reverse order.
	return string(runes)
}
