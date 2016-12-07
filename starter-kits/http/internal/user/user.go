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

	u := []User{}
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : List: MGO :\n\ndb.users.find()\n\n", traceID)
		return collection.Find(nil).All(&u)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : List : Completed : ERROR :", err)
		return nil, err
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
func Create(ctx context.Context, traceID string, db *mgo.Session, cu *CreateUser) (*User, error) {
	log.Println(traceID, ": services : Users : Create : Started")

	now := time.Now()

	u := User{
		UserID:       bson.NewObjectId().Hex(),
		UserType:     cu.UserType,
		FirstName:    cu.FirstName,
		LastName:     cu.LastName,
		Email:        cu.Email,
		Company:      cu.Company,
		DateCreated:  &now,
		DateModified: &now,
		Addresses:    make([]Address, len(cu.Addresses)),
	}

	for i, cua := range cu.Addresses {
		u.Addresses[i] = Address{
			Type:         cua.Type,
			LineOne:      cua.LineOne,
			LineTwo:      cua.LineTwo,
			City:         cua.City,
			State:        cua.State,
			Zipcode:      cua.State,
			Phone:        cua.Phone,
			DateCreated:  &now,
			DateModified: &now,
		}
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
	return &u, nil
}

// Update replaces a user document in the database.
func Update(ctx context.Context, traceID string, db *mgo.Session, userID string, cu *CreateUser) error {
	log.Println(traceID, ": services : Users : Update : Started")

	// TODO: Check the userID is a valid bson id.

	now := time.Now()
	cu.DateModified = &now
	for _, cua := range cu.Addresses {
		cua.DateModified = &now
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		m := bson.M{"$set": cu}
		log.Printf("%s : services : Users : Update : MGO :\n\ndb.users.update(%s, %s)\n\n", traceID, app.Query(q), app.Query(m))
		return collection.Update(q, m)
	}

	if err := app.ExecuteDB(db, usersCollection, f); err != nil {
		log.Println(traceID, ": services : Users : Create : Completed : ERROR :", err)
		return err
	}

	log.Println(traceID, ": services : Users : Update : Completed")
	return nil
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
