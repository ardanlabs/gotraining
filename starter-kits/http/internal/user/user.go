// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user

import (
	"context"
	"log"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

//==============================================================================

// List retrieves a list of existing users from the database.
func List(ctx context.Context, traceID string, dbSes *mgo.Session) ([]User, error) {
	log.Println(traceID, ": List : Started")

	u := []User{}
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : List : MGO :\n\ndb.users.find()\n\n", traceID)
		return collection.Find(nil).All(&u)
	}

	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		log.Println(traceID, ": List : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(traceID, ": List : Completed")
	return u, nil
}

// Retrieve gets the specified user from the database.
func Retrieve(ctx context.Context, traceID string, dbSes *mgo.Session, userID string) (*User, error) {
	log.Println(traceID, ": Retrieve : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(traceID, ": Retrieve : Completed : ERROR :", web.ErrInvalidID)
		return nil, web.ErrInvalidID
	}

	var u *User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : Retrieve : MGO :\n\ndb.users.find(%s)\n\n", traceID, db.Query(q))
		return collection.Find(q).One(&u)
	}

	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			log.Println(traceID, ": Retrieve : Completed : ERROR :", err)
			return nil, err
		}

		log.Println(traceID, ": Retrieve : Completed : ERROR : Not Found")
		return nil, web.ErrNotFound
	}

	log.Println(traceID, ": Retrieve : Completed")
	return u, nil
}

// Create inserts a new user into the database.
func Create(ctx context.Context, traceID string, dbSes *mgo.Session, cu *CreateUser) (*User, error) {
	log.Println(traceID, ": Create : Started")

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
		log.Printf("%s : Create : MGO :\n\ndb.users.insert(%s)\n\n", traceID, db.Query(u))
		return collection.Insert(u)
	}

	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		log.Println(traceID, ": Create : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(traceID, ": Create : Completed")
	return &u, nil
}

// Update replaces a user document in the database.
func Update(ctx context.Context, traceID string, dbSes *mgo.Session, userID string, cu *CreateUser) error {
	log.Println(traceID, ": Update : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(traceID, ": Update : Completed : ERROR :", web.ErrInvalidID)
		return web.ErrInvalidID
	}

	now := time.Now()
	cu.DateModified = &now
	for _, cua := range cu.Addresses {
		cua.DateModified = &now
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		m := bson.M{"$set": cu}
		log.Printf("%s : Update : MGO :\n\ndb.users.update(%s, %s)\n\n", traceID, db.Query(q), db.Query(m))
		return collection.Update(q, m)
	}

	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		log.Println(traceID, ": Update : Completed : ERROR :", err)
		if err == mgo.ErrNotFound {
			return web.ErrNotFound
		}
		return err
	}

	log.Println(traceID, ": Update : Completed")
	return nil
}

// Delete removes a user from the database.
func Delete(ctx context.Context, traceID string, dbSes *mgo.Session, userID string) error {
	log.Println(traceID, ": Delete : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(traceID, ": Delete : Completed : ERROR :", web.ErrInvalidID)
		return web.ErrInvalidID
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : Delete : MGO :\n\ndb.users.remove(%s)\n\n", traceID, db.Query(q))
		return collection.Remove(q)
	}

	if err := db.Execute(dbSes, usersCollection, f); err != nil {
		log.Println(traceID, ": Delete : Completed : ERROR :", err)
		if err == mgo.ErrNotFound {
			return web.ErrNotFound
		}
		return err
	}

	log.Println(traceID, ": Delete : Completed")
	return nil
}
