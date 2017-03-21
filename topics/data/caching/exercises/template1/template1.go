// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to show how to cache data from an API, and then
// use that data in analyzing a dataset.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	// statusURL provides an explanation of Citibike station statuses.
	statusURL = "https://feeds.citibikenyc.com/stations/status.json"

	// currentStatusURL provides the current Citibike station statuses.
	currentStatusURL = "https://feeds.citibikenyc.com/stations/stations.json"
)

// stationData is used to unmarshal the JSON document returned form citiBikeURL.
type stationData struct {
	ExecutionTime   string    `json:"executionTime"`
	StationBeanList []station `json:"stationBeanList"`
}

// station is used to unmarshal each of the station documents in stationData.
type station struct {
	ID          int    `json:"id"`
	StationName string `json:"stationName"`
	StatusKey   int    `json:"statusKey"`
}

func main() {

	// Get the JSON status codes from the statusURL.
	response, err := http.Get(statusURL)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the response body.
	defer response.Body.Close()

	// Read the body of the response into []byte.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data into a map. Here the keys are the
	// status codes and the values are the meanings of the codes.
	var statuses map[string]string
	if err := json.Unmarshal(body, &statuses); err != nil {
		log.Fatal(err)
		return
	}

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds
	c := cache.New(5*time.Minute, 30*time.Second)

	// Put the status codes and meanings into the cache.
	for k, v := range statuses {
		c.Set(k, v, cache.DefaultExpiration)
	}

	// Get the current CitiBike station statuses from currentStatusURL.

	// Unmarshal the JSON data into the variable.

	// Determine which stations are not in service.

	// Get the respective code from the Cache.
}
