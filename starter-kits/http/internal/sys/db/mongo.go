// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package db

import (
	"encoding/json"
	"log"
	"time"

	"github.com/pkg/errors"
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

func init() {

	// Logging this here because it is currently hardcoded.
	log.Printf("Init : Host[%s] Database[%s]\n", mongoDBHosts, database)
}

// Init sets up the MongoDB environment and provides a master session.
func Init() (*mgo.Session, error) {

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
		return nil, errors.Wrap(err, "dial connection")
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	return session, nil
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

	// Capture the specified collection from our connection.
	collection := session.DB("").C(collectionName)
	if collection == nil {
		return errors.Errorf("Collection %s does not exist", collectionName)
	}

	// Execute the MongoDB call.
	if err := f(collection); err != nil {
		return errors.Wrap(err, "executing mgo")
	}

	return nil
}
