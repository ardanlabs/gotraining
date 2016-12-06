// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/app"
	mgo "gopkg.in/mgo.v2"
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

// Mongo handles session management.
func Mongo(h app.Handler) app.Handler {

	// Init the MongoDB master session.
	session := NewMGOSession()

	// Wrap the handlers inside a session copy/close.
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(app.KeyValues).(*app.Values)

		// Get a MongoDB session connection.
		log.Printf("%s : Mongo : *****> Capture Mongo Session\n", v.TraceID)
		v.DB = session.Copy()

		// Defer releasing the db session connection.
		defer func() {
			log.Printf("%s : Mongo : *****> Release Mongo Session\n", v.TraceID)
			v.DB.Close()
		}()

		return h(ctx, w, r, params)
	}
}

// NewMGOSession sets up the MongoDB environment.
func NewMGOSession() *mgo.Session {
	log.Printf("api : mongodb : init : Started : Host[%s] Database[%s]\n", mongoDBHosts, database)

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
		log.Fatalln("MongoDB Dial", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	log.Printf("api : mongodb : init : Completed : Host[%s] Database[%s]\n", mongoDBHosts, database)
	return session
}
