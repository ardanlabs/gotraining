package ride_test

import (
	"fmt"
	
	"shorts/ride"
)

func ExampleRideFare() {
	fare := ride.RideFare(3.7, true)
	fmt.Println(fare)

	// Output:
	// 765
}
