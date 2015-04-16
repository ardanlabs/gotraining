// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// go test -v example1_example_test.go

// Package tests provides testing examples.
package tests

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/09-testing/01-testing/example1/buoy"
)

// ExampleFindStation provides a basic example for using the ExampleFindStation API.
func ExampleFindStation() {
	station, err := buoy.FindStation("42002")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*station)

	// Output:
	// {ObjectIdHex("51e873dbb199b3cb9318d996") 42002 Station 42002 - West Gulf 207 NM East of Brownsville, TX {2.23694 150 4.47388} {Point [-93.666 25.79]}}
}
