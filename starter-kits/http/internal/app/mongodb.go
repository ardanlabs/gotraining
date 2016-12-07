// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package app

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// Query provides a string version of the value
func Query(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(json)
}

// ExecuteDB the MongoDB literal function.
func ExecuteDB(session *mgo.Session, collectionName string, f func(*mgo.Collection) error) error {
	log.Printf("ExecuteDB : Started : Collection[%s]\n", collectionName)

	// Capture the specified collection.
	collection := session.DB("").C(collectionName)
	if collection == nil {
		err := fmt.Errorf("Collection %s does not exist", collectionName)
		log.Println("ExecuteDB : Completed : ERROR :", err)
		return err
	}

	// Execute the MongoDB call.
	if err := f(collection); err != nil {
		log.Println("ExecuteDB : Completed : ERROR :", err)
		return err
	}

	log.Println("ExecuteDB : Completed")
	return nil
}
