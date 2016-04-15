// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to marshal a user defined
// struct type into a string.
package main

import (
	"encoding/json"
	"fmt"
)

type (
	// buoyCondition contains information for an individual station.
	buoyCondition struct {
		WindSpeed     float64 `json:"wind_speed_milehour"`
		WindDirection int     `json:"wind_direction_degnorth"`
		WindGust      float64 `json:"gust_wind_speed_milehour"`
	}

	// buoyLocation contains the buoys location.
	buoyLocation struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	// BuoyStation contains information for an individual station.
	buoyStation struct {
		StationID string        `json:"station_id"`
		Name      string        `json:"name"`
		LocDesc   string        `json:"location_desc"`
		Condition buoyCondition `json:"condition"`
		Location  buoyLocation  `json:"location"`
	}
)

func main() {

	// Declare and initialize a value of type buoyStation.
	station := buoyStation{
		StationID: "42036",
		Name:      "Station 42036 - West Tampa",
		LocDesc:   "112 NM WNW of Tampa, FL",
		Condition: buoyCondition{
			WindSpeed:     17.895520000000001204,
			WindDirection: 190,
			WindGust:      22.36939999999999884,
		},
		Location: buoyLocation{
			Type:        "Point",
			Coordinates: []float64{-84.516999999999995907, 28.5},
		},
	}

	// Marshal the buoyStation value into a JSON
	// string representation.
	r1, err := json.Marshal(&station)
	if err != nil {
		fmt.Println("Marshal", err)
		return
	}

	// Convert the byte slice to a string and display.
	fmt.Printf("%s\n\n", r1)

	// Marshal the buoyStation value into a pretty-print
	// JSON string representation.
	r2, err := json.MarshalIndent(&station, "", "    ")
	if err != nil {
		fmt.Println("MarshalIndent", err)
		return
	}

	// Convert the byte slice to a string and display.
	fmt.Println(string(r2))
}
