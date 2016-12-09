// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package db

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

// MongoDB connection information.
// This information is hardcoded into this app to allow you to run
// this code without requiring configuration. Please don't do this
// in production code; I recommend using environmental variables and
// a package like envconfig [https://github.com/kelseyhightower/envconfig].
const (
	mongoDBHosts = "ds039441.mongolab.com:39441"
	authDatabase = "gotraining"
	authUserName = "got"
	authPassword = "got2015"
	database     = "gotraining"
)

//==============================================================================

// Init sets up the MongoDB environment and provides a master session.
func Init() *mgo.Session {
	log.Printf("Init : Host[%s] Database[%s]\n", mongoDBHosts, database)

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: authDatabase,
		Username: authUserName,
		Password: authPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(&mongoDBDialInfo)
	if err != nil {
		log.Fatalln("Init : MongoDB Dial", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	return session
}

// Query provides a string version of the value
func Query(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(json)
}

// Execute the MongoDB function using session and collection information.
func Execute(session *mgo.Session, collectionName string, f func(*mgo.Collection) error) error {
	log.Printf("Execute : Started : Collection[%s]\n", collectionName)

	// Capture the specified collection from our connection.
	collection := session.DB("").C(collectionName)
	if collection == nil {
		err := fmt.Errorf("Collection %s does not exist", collectionName)
		log.Println("Execute : Completed : ERROR :", err)
		return err
	}

	// Execute the MongoDB call.
	if err := f(collection); err != nil {
		log.Println("Execute : Completed : ERROR :", err)
		return err
	}

	log.Println("Execute : Completed")
	return nil
}
