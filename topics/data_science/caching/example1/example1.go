// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// go build
// ./example1
//
// Sample program to show how to cache data from an API.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

// statusURL provides an explanation of Citibike station statuses.
const statusURL = "https://feeds.citibikenyc.com/stations/status.json"

func main() {

	// Get the JSON response from the URL.
	response, err := http.Get(statusURL)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing response body.
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

	// For a sanity check. Output the keys and values in the cache
	// to standard out.
	for k := range statuses {
		v, found := c.Get(k)
		if found {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
	}
}
