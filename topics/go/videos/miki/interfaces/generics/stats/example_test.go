package stats

import "fmt"

func ExampleMaxInts() {
	vals := []int{23, 8, 4, 42, 16, 15}
	fmt.Println(MaxInts(vals))
	fmt.Println(MaxInts(nil))

	// Output:
	// 42 <nil>
	// 0 max of empty slice
}

func ExampleMaxFloat64s() {
	vals := []float64{23, 8, 4, 42, 16, 15}
	fmt.Println(MaxFloat64s(vals))
	fmt.Println(MaxFloat64s(nil))

	// Output:
	// 42 <nil>
	// 0 max of empty slice
}
