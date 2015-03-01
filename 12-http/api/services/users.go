// Package services provides business and data processing.
package services

import (
	"errors"
	"log"
	"time"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

var (
	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("No user(s) found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

// usersService maintains the set of services for the users api.
type usersService struct{}

// Users fronts the access to the users service functionality.
var Users usersService

// List retrieves a list of existing users from the database.
func (us usersService) List(c *app.Context) ([]models.User, error) {
	log.Println(c.SessionID, ": services : Users : List : Started")

	var u []models.User
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : List: MGO :\n\ndb.users.find()\n\n", c.SessionID)
		return collection.Find(nil).All(&u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : List : Completed : ERROR :", err)
		return nil, err
	}

	if len(u) == 0 {
		log.Println(c.SessionID, ": services : Users : List : Completed : ERROR :", ErrNotFound)
		return nil, ErrNotFound
	}

	log.Println(c.SessionID, ": services : Users : List : Completed")
	return u, nil
}

// Retrieve gets the specified user from the database.
func (us usersService) Retrieve(c *app.Context, userID string) (*models.User, error) {
	log.Println(c.SessionID, ": services : Users : Retrieve : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(c.SessionID, ": services : Users : Retrieve : Completed : ERROR :", ErrInvalidID)
		return nil, ErrInvalidID
	}

	var u *models.User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : services : Users : Retrieve: MGO :\n\ndb.users.find(%s)\n\n", c.SessionID, app.Query(q))
		return collection.Find(q).One(&u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		if err != mgo.ErrNotFound {
			log.Println(c.SessionID, ": services : Users : Retrieve : Completed : ERROR :", err)
			return nil, err
		}

		log.Println(c.SessionID, ": services : Users : Retrieve : Completed : ERROR : Not Found")
		return nil, ErrNotFound
	}

	log.Println(c.SessionID, ": services : Users : Retrieve : Completed")
	return u, nil
}

// Create inserts a new user into the database.
func (us usersService) Create(c *app.Context, u *models.User) ([]app.Invalid, error) {
	log.Println(c.SessionID, ": services : Users : Create : Started")

	now := time.Now()

	u.UserID = bson.NewObjectId().Hex()
	u.DateCreated = &now
	u.DateModified = &now
	for _, ua := range u.Addresses {
		ua.DateCreated = &now
		ua.DateModified = &now
	}

	if v, err := u.Validate(); err != nil {
		log.Println(c.SessionID, ": services : Users : Create : Completed : ERROR :", err)
		return v, ErrValidation
	}

	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : Create : MGO :\n\ndb.users.insert(%s)\n\n", c.SessionID, app.Query(u))
		return collection.Insert(u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : Create : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : Users : Create : Completed")
	return nil, nil
}

// Update replaces a user document in the database.
func (us usersService) Update(c *app.Context, userID string, u *models.User) ([]app.Invalid, error) {
	log.Println(c.SessionID, ": services : Users : Update : Started")

	if v, err := u.Validate(); err != nil {
		log.Println(c.SessionID, ": services : Users : Update : Completed : ERROR :", err)
		return v, ErrValidation
	}

	if u.UserID == "" {
		u.UserID = userID
	}

	if userID != u.UserID {
		log.Println(c.SessionID, ": services : Users : Update : Completed : ERROR :", ErrValidation)
		return []app.Invalid{{Fld: "UserID", Err: "Specified UserID does not match user value."}}, ErrValidation
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": u.UserID}
		log.Printf("%s : services : Users : Update : MGO :\n\ndb.users.update(%s, %s)\n\n", c.SessionID, app.Query(q), app.Query(u))
		return collection.Update(q, u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : Create : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : Users : Update : Completed")
	return nil, nil
}

// Delete inserts a new user into the database.
func (us usersService) Delete(c *app.Context, userID string) error {
	log.Println(c.SessionID, ": services : Users : Delete : Started")

	if !bson.IsObjectIdHex(userID) {
		log.Println(c.SessionID, ": services : Users : Delete : Completed : ERROR :", ErrInvalidID)
		return ErrInvalidID
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"user_id": userID}
		log.Printf("%s : services : Users : Delete : MGO :\n\ndb.users.remove(%s)\n\n", c.SessionID, app.Query(q))
		return collection.Remove(q)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : Delete : Completed : ERROR :", err)
		return err
	}

	log.Println(c.SessionID, ": services : Users : Delete : Completed")
	return nil
}
