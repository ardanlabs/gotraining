// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// go test -v example1_example_test.go

// Package tests provides testing examples.
package tests

import (
	"encoding/json"
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

	d, err := json.MarshalIndent(station, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(d))

	// Output:
	// {
	//     "ID": "51e873dbb199b3cb9318d996",
	//     "station_id": "42002",
	//     "name": "Station 42002 - West Gulf",
	//     "location_desc": "207 NM East of Brownsville, TX",
	//     "condition": {
	//         "wind_speed_milehour": 2.23694,
	//         "wind_direction_degnorth": 150,
	//         "gust_wind_speed_milehour": 4.47388
	//     },
	//     "location": {
	//         "type": "Point",
	//         "coordinates": [
	//             -93.666,
	//             25.79
	//         ]
	//     }
	// }
}
