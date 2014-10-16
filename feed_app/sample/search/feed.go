// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// feed.go provides support for reading the JSON file of feeds and
// returning a slice for application processing.

// Package search contains the framework for using matchers to retreive
// and search different types of content.
package search

import (
	"encoding/json"
	"os"
)

var dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds(path ...string) ([]Feed, error) {
	// Using a variadic argument to create an optional parameter.
	// Used by testing to manipulate the path to the file.
	if path != nil {
		dataFile = path[0] + dataFile
	}

	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
