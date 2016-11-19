// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to save data from an API in an embedded k/v store.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
)

// statusURL provides an explanation of Citibike station statuses.
const statusURL = "https://feeds.citibikenyc.com/stations/status.json"

func main() {

	// Get the JSON response from the URL.
	response, err := http.Get(statusURL)
	if err != nil {
		log.Fatal(err)
	}

	// Defer response body close.
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

	// Remove the embedded k/v store if it already exists.
	os.Remove("embedded.db")

	// Open the embedded.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a "bucket" in the boltdb file for our data.
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("CitibikeCache"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Put the status codes and meanings into the boltdb file.
	for k, v := range statuses {
		if err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("CitibikeCache"))
			err := b.Put([]byte(k), []byte(v))
			return err
		}); err != nil {
			log.Fatal(err)
		}
	}

	// For a sanity check. Output the keys and values in the embedded
	// boltdb file to standard out.
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("CitibikeCache"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
