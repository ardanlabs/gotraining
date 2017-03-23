// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package db

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

// ErrInvalidDBProvided is returned in the event that an uninitialized db is
// used to perform actions against.
var ErrInvalidDBProvided = errors.New("invalid DB provided")

// master manages a set of different MongoDB master sessions.
var master = struct {
	sync.RWMutex
	ses map[string]*mgo.Session
}{
	ses: make(map[string]*mgo.Session),
}

// RegMasterSession adds a new master session to the set. If no url is provided,
// it will default to localhost:27017.
func RegMasterSession(name string, url string, timeout time.Duration) error {
	master.Lock()
	defer master.Unlock()

	if _, exists := master.ses[name]; exists {
		return errors.New("master session already exists")
	}

	ses, err := newMasterSession(url, timeout)
	if err != nil {
		return errors.Wrapf(err, "mongo.New: %s,%v", url, timeout)
	}

	master.ses[name] = ses

	return nil
}

// Query provides a string version of the value
func Query(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(json)
}

// new creates a new master session. If no url is provided, it will defaul to
// localhost:27017. If a zero value timeout is specified, a timeout of 60sec
// will be used instead.
func newMasterSession(url string, timeout time.Duration) (*mgo.Session, error) {

	// Set the default timeout for the session.
	if timeout == 0 {
		timeout = 60 * time.Second
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	ses, err := mgo.DialWithTimeout(url, timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "mgo.DialWithTimeout: %s,%v", url, timeout)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	ses.SetMode(mgo.Monotonic, true)

	return ses, nil
}

// DB represents our mongo database session.
type DB struct {
	database *mgo.Database
	session  *mgo.Session
}

// New returns a new DB value for use with MongoDB based on a registered
// master session.
func New(name string) (*DB, error) {
	var mSes *mgo.Session
	var exists bool
	master.RLock()
	{
		mSes, exists = master.ses[name]
	}
	master.RUnlock()

	if !exists {
		return nil, errors.Errorf("master sesssion %q does not exist", name)
	}

	ses := mSes.Copy()

	// As per the mgo documentation, https://godoc.org/gopkg.in/mgo.v2#Session.DB
	// if no database name is specified, then use the default one, or the one that
	// the connection was dialed with.
	db := mSes.DB("")

	dbOut := DB{
		database: db,
		session:  ses,
	}

	return &dbOut, nil
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() {
	db.session.Close()
}

// Execute is used to execute MongoDB commands.
func (db *DB) Execute(collName string, f func(*mgo.Collection) error) error {
	if db == nil || db.session == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil || db.session == nil")
	}

	return f(db.database.C(collName))
}

// ExecuteTimeout is used to execute MongoDB commands with a timeout.
func (db *DB) ExecuteTimeout(collName string, f func(*mgo.Collection) error, timeout time.Duration) error {
	if db == nil || db.session == nil {
		return errors.Wrap(ErrInvalidDBProvided, "db == nil || db.session == nil")
	}

	db.session.SetSocketTimeout(timeout)

	return f(db.database.C(collName))
}
