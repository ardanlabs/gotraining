// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// go test -v example1_table_test.go

// Table tests for example 1.
package main

import (
	"testing"

	"github.com/ArdanStudios/gotraining/06-testing/example1/buoy"
)

// stationIDs maintain a list of ids to check.
var stationIDs = []string{"42002", "42001", "42012", "42019"}

// TestTableStation checks the station service call is working
func TestTableStation(t *testing.T) {
	// Iterate through the slice of station ids.
	for index := 0; index < len(stationIDs); index++ {
		// Station id to test.
		stationID := stationIDs[index]

		// Perform call to find a station.
		station, err := buoy.FindStation(stationID)

		// We should not get an error.
		t.Log("There should be no error after the call to FindStation.")
		if err != nil {
			t.Fatalf("ERROR : %s", err)
		}

		// We should not have a nil pointer for the station.
		t.Log("We should not have a nil pointer for the station.")
		if station == nil {
			t.Fatalf("ERROR : %s", err)
		}

		// We should get back the station document.
		t.Logf("StationID \"%s\" should match \"%s\" in the station document.", stationID, station.StationID)
		if station.StationID != stationID {
			t.Errorf("ERROR : Expecting[%s] Received[%s]", stationID, station.StationID)
		}
	}
}
