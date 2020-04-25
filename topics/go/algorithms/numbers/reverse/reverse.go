package numbers

// Reverse takes the specified integer and reverses it.
func Reverse(num int) int {

	// Construct result to its zero value.
	var result int

	// Loop until num is zero.
	for num != 0 {

		// Perform a modulus operation to get the last digit from the value set in num.
		// https://www.geeksforgeeks.org/find-first-last-digits-number/
		last := num % 10

		// Multiple the current result by 10 to shit the digits in
		// the current result to the left.
		result = result * 10

		// Add the digit we took from the end of num to the result.
		result += last

		// // Remove the digit we just reversed from num.
		num = num / 10
	}
	return result
}
