// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/VxnL7AEZSl

// go build -gcflags -m

// Package prediction provides code to show how branch prediction can affect performance.
package prediction

// benchmark is the boilerpalte code to view branch prediction
// misses by the hardware and how it affects performance.
func benchmark(n int, f func(j int, d int) int) {
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var j int
	k := -1

	// n is provided by the benchmark framework.
	for i := 0; i < n; i++ {
		// We want to iterate over the array for
		// consistency with each test.
		if k == 9 {
			k = -1
		}
		k++

		// Call the specified test function. This function
		// will be inlined.
		j = f(j, data[k])
	}
	j++
}

// ifMostlyTrue is the most idiomatic way to write this code. If tests for a
// condition that occurrs 90% of the time.
func ifMostlyTrue(j int, d int) int {
	if d != 3 {
		return j + 1
	}
	return j + 2
}

// ifNotMostlyTrue is an idiomatic way to write this code. If tests for a
// condition that occurrs only 10% of the time.
func ifNotMostlyTrue(j int, d int) int {
	if d == 3 {
		return j + 2
	}
	return j + 1
}
