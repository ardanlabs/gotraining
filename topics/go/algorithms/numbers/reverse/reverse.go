package numbers

// Reverse takes in a integer and returns a the integer in reverse order.
func Reverse(num int) int {

	// Set result to zero.
	var result int

	// Loop until num is not zero.
	for num != 0 {

		// Get the last digit from num.
		// Example: 125 % 10 = 5.
		last := num % 10

		// Move result one place to the left and add last.
		result = result*10 + last

		// Remove the right most digit from num.
		// Example: 125 / 10 = 12.
		num = num / 10
	}
	return result
}
