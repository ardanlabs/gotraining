// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// go build
// ./exercise1
//
// Sample program to show how to unmarshal JSON data from an API.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-hep/csvutil"
)

// citiBikeURL provides the station statuses of CitiBike bike sharing stations.
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData is used to unmarshal the JSON document returned form citiBikeURL.
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in stationData.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

// outputData includes the fields that will be output to the CSV.
type outputData struct {
	ID             string
	BikesAvailable int
	BikesDisabled  int
	DocksAvailable int
	DocksDisabled  int
}

func main() {

	// Get the JSON response from the URL.
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the response Body.
	defer response.Body.Close()

	// Read the body of the response into []byte.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Declare a variable of type stationData.
	var sd stationData

	// Unmarshal the JSON data into the variable.
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	// Register the output CSV table.
	fname := "citibike.csv"
	tbl, err := csvutil.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer tbl.Close()
	tbl.Writer.Comma = ','

	// Write the header.
	header := "station_id,bikes_available,bikes_disabled,docks_available,docks_disabled\n"
	if err = tbl.WriteHeader(header); err != nil {
		log.Fatal(err)
	}

	// Write the station data from the JSON.
	for _, station := range sd.Data.Stations {
		data := outputData{
			ID:             station.ID,
			BikesAvailable: station.NumBikesAvailable,
			BikesDisabled:  station.NumBikesDisabled,
			DocksAvailable: station.NumDocksAvailable,
			DocksDisabled:  station.NumDocksDisabled,
		}
		err = tbl.WriteRow(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure the output file is properly saved.
	if err := tbl.Close(); err != nil {
		log.Fatal(err)
	}
}
