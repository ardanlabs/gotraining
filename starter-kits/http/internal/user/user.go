// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user

import (
	"context"
	"log"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/kit/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/kit/web"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

// List retrieves a list of existing users from the database.
func List(ctx context.Context, traceID string, dbSes *mgo.Session) ([]User, error) {
	u := []User{}
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : List : MGO : db.users.find()\n", traceID)
		return collection.Find(nil).All(&u)
	}
	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		return nil, errors.Wrap(err, "db.Execute")
	}

	return u, nil
}

// Retrieve gets the specified user from the database.
func Retrieve(ctx context.Context, traceID string, dbSes *mgo.Session, userID string) (*User, error) {
	if !bson.IsObjectIdHex(userID) {
		return nil, errors.Wrapf(web.ErrInvalidID, "bson.IsObjectIdHex: %s", userID)
	}

	var u *User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : Retrieve : MGO : db.users.find(%s)\n", traceID, db.Query(q))
		return collection.Find(q).One(&u)
	}
	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			err = web.ErrNotFound
		}
		return nil, errors.Wrap(err, "db.Execute")
	}

	return u, nil
}

// Create inserts a new user into the database.
func Create(ctx context.Context, traceID string, dbSes *mgo.Session, cu *CreateUser) (*User, error) {
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
		log.Printf("%s : Create : MGO : db.users.insert(%s)\n", traceID, db.Query(u))
		return collection.Insert(u)
	}
	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		return nil, errors.Wrap(err, "db.Execute")
	}

	return &u, nil
}

// Update replaces a user document in the database.
func Update(ctx context.Context, traceID string, dbSes *mgo.Session, userID string, cu *CreateUser) error {
	if !bson.IsObjectIdHex(userID) {
		return errors.Wrap(web.ErrInvalidID, "check objectid")
	}

	now := time.Now()
	cu.DateModified = &now
	for _, cua := range cu.Addresses {
		cua.DateModified = &now
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		m := bson.M{"$set": cu}
		log.Printf("%s : Update : MGO : db.users.update(%s, %s)\n", traceID, db.Query(q), db.Query(m))
		return collection.Update(q, m)
	}
	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			err = web.ErrNotFound
		}
		return errors.Wrap(err, "db.Execute")
	}

	return nil
}

// Delete removes a user from the database.
func Delete(ctx context.Context, traceID string, dbSes *mgo.Session, userID string) error {
	if !bson.IsObjectIdHex(userID) {
		return errors.Wrapf(web.ErrInvalidID, "bson.IsObjectIdHex: %s", userID)
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : Delete : MGO : db.users.remove(%s)\n", traceID, db.Query(q))
		return collection.Remove(q)
	}
	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			err = web.ErrNotFound
		}
		return errors.Wrap(err, "db.Execute")
	}

	return nil
}
