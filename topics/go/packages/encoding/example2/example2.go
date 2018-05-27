// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to unmarshal a JSON document into
// a user defined struct type from a file.
package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// Open the file.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Open File", err)
		return
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of buoy stations.
	var stations []buoyStation
	err = json.NewDecoder(file).Decode(&stations)
	if err != nil {
		fmt.Println("Decode File", err)
		return
	}

	// Iterate over the slice and display
	// each station.
	for _, station := range stations {
		fmt.Printf("%+v\n\n", station)
	}
}
