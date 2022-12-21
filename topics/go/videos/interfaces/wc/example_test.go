package wc

import "fmt"

func ExampleLineCount() {
	fmt.Println(LineCount("testdata/road.txt"))
	fmt.Println(LineCount("testdata/aow.txt.gz"))

	// Output:
	// 8 <nil>
	// 6982 <nil>
}
