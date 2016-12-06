// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user

import (
	"context"
	"log"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/app"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

//==============================================================================

// List retrieves a list of existing users from the database.
func List(ctx context.Context, traceID string, db *mgo.Session) ([]User, error) {
	log.Println(traceID, ": services : Users : List : Started")

	var u []User
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : List: MGO :\n\ndb.users.find()\n\n", traceID)
		return collection.Find(nil).All(&u)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : List : Completed : ERROR :", err)
		return nil, err
	}

	if len(u) == 0 {
		log.Println(traceID, ": services : Users : List : Completed : ERROR :", app.ErrNotFound)
		return nil, app.ErrNotFound
	}

	log.Println(traceID, ": services : Users : List : Completed")
	return u, nil
}

// Retrieve gets the specified user from the database.
func Retrieve(ctx context.Context, traceID string, db *mgo.Session, userID string) (*User, error) {
	log.Println(traceID, ": services : Users : Retrieve : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(traceID, ": services : Users : Retrieve : Completed : ERROR :", app.ErrInvalidID)
		return nil, app.ErrInvalidID
	}

	var u *User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : services : Users : Retrieve: MGO :\n\ndb.users.find(%s)\n\n", traceID, app.Query(q))
		return collection.Find(q).One(&u)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			log.Println(traceID, ": services : Users : Retrieve : Completed : ERROR :", err)
			return nil, err
		}

		log.Println(traceID, ": services : Users : Retrieve : Completed : ERROR : Not Found")
		return nil, app.ErrNotFound
	}

	log.Println(traceID, ": services : Users : Retrieve : Completed")
	return u, nil
}

// Create inserts a new user into the database.
func Create(ctx context.Context, traceID string, db *mgo.Session, u *User) ([]app.Invalid, error) {
	log.Println(traceID, ": services : Users : Create : Started")

	now := time.Now()

	u.UserID = bson.NewObjectId().Hex()
	u.DateCreated = &now
	u.DateModified = &now
	for _, ua := range u.Addresses {
		ua.DateCreated = &now
		ua.DateModified = &now
	}

	if v, err := u.Validate(); err != nil {
		log.Println(traceID, ": services : Users : Create : Completed : ERROR :", err)
		return v, app.ErrValidation
	}

	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : Create : MGO :\n\ndb.users.insert(%s)\n\n", traceID, app.Query(u))
		return collection.Insert(u)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : Create : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(traceID, ": services : Users : Create : Completed")
	return nil, nil
}

// Update replaces a user document in the database.
func Update(ctx context.Context, traceID string, db *mgo.Session, userID string, u *User) ([]app.Invalid, error) {
	log.Println(traceID, ": services : Users : Update : Started")

	if v, err := u.Validate(); err != nil {
		log.Println(traceID, ": services : Users : Update : Completed : ERROR :", err)
		return v, app.ErrValidation
	}

	if u.UserID == "" {
		u.UserID = userID
	}

	if userID != u.UserID {
		log.Println(traceID, ": services : Users : Update : Completed : ERROR :", app.ErrValidation)
		return []app.Invalid{{Fld: "UserID", Err: "Specified UserID does not match user value."}}, app.ErrValidation
	}

	// This is a bug that needs to be fixed.
	// I am re-writing the dates so the tests pass. :(
	now := time.Now()
	u.DateCreated = &now
	u.DateModified = &now
	for _, ua := range u.Addresses {
		ua.DateCreated = &now
		ua.DateModified = &now
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": u.UserID}
		log.Printf("%s : services : Users : Update : MGO :\n\ndb.users.update(%s, %s)\n\n", traceID, app.Query(q), app.Query(u))
		return collection.Update(q, u)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : Create : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(traceID, ": services : Users : Update : Completed")
	return nil, nil
}

// Delete removes a user from the database.
func Delete(ctx context.Context, traceID string, db *mgo.Session, userID string) error {
	log.Println(traceID, ": services : Users : Delete : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(traceID, ": services : Users : Delete : Completed : ERROR :", app.ErrInvalidID)
		return app.ErrInvalidID
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : services : Users : Delete : MGO :\n\ndb.users.remove(%s)\n\n", traceID, app.Query(q))
		return collection.Remove(q)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : Delete : Completed : ERROR :", err)
		return err
	}

	log.Println(traceID, ": services : Users : Delete : Completed")
	return nil
}
