// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// go test -v example1_test.go

// Standard tests for example 1.
package main

import (
	"testing"

	"github.com/ArdanStudios/gotraining/06-testing/example1/buoy"
)

// TestStation checks the station service call is working
func TestStation(t *testing.T) {
	// Perform call to find a station.
	stationID := "42002"
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

	// The name in this document should match this known value.
	name := "Station 42002 - West Gulf"
	t.Logf("Name \"%s\" should match \"%s\" in the station document.", name, station.Name)
	if station.Name != name {
		t.Errorf("ERROR : Expecting[%s] Received[%s]", name, station.Name)
	}
}

// Test_InvalidStation checks that an error occurs with an invalid station id.
func Test_InvalidStation(t *testing.T) {
	// Perform call to find a station.
	stationID := "00000"
	station, err := buoy.FindStation(stationID)

	// We should not get an error.
	t.Log("There should be no error after the call to FindStation.")
	if err != nil {
		t.Fatalf("ERROR : %s", err)
	}

	// We should have a nil pointer for the station.
	t.Log("We should get back a nil pointer for the station.")
	if station != nil {
		t.Errorf("ERROR : %s", err)
	}
}
