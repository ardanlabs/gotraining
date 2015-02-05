// Package services provides business and data processing.
package services

import (
	"errors"
	"log"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

// usersService maintains the set of services for the users api.
type usersService struct{}

// Users fronts the access to the users service functionality.
var Users usersService

// List retrieves a list of existing users from the database.
func (us usersService) List(c *app.Context) ([]models.User, error) {
	log.Println(c.SessionID, ": services : Users : List : Started")

	var users []models.User
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : List: MGO :\n\ndb.users.find()\n\n", c.SessionID)
		return collection.Find(nil).All(&users)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : List : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : Users : List : Completed")
	return users, nil
}

// Retrieve gets the specified user from the database.
func (us usersService) Retrieve(c *app.Context, id string) (*models.User, error) {
	log.Println(c.SessionID, ": services : Users : Retrieve : Started")

	if ok := bson.IsObjectIdHex(id); !ok {
		err := errors.New("Invalid user id.")
		log.Println(c.SessionID, ": services : Users : Retrieve : Completed : ERROR :", err)
		return nil, err
	}

	var u *models.User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"_id": bson.ObjectIdHex(id)}
		log.Printf("%s : services : Users : Retrieve: MGO :\n\ndb.users.find(%s)\n\n", c.SessionID, app.Query(q))
		return collection.Find(q).One(&u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : Retrieve : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : Users : Retrieve : Completed")
	return u, nil
}

// Create inserts a new user into the database.
func (us usersService) Create(c *app.Context, u *models.User) error {
	log.Println(c.SessionID, ": services : Users : Create : Started")

	now := time.Now()
	u.ID = bson.NewObjectId()
	u.DateModified = &now
	u.DateCreated = u.DateModified

	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : Users : Create : MGO :\n\ndb.users.insert(%s)\n\n", c.SessionID, app.Query(u))
		return collection.Insert(u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : Users : Create : Completed : ERROR :", err)
		return err
	}

	log.Println(c.SessionID, ": services : Users : Create : Completed")
	return nil
}

// Delete inserts a new user into the database.
func (us usersService) Delete(c *app.Context, id string) error {
	log.Println(c.SessionID, ": services : Users : Delete : Started")

	if ok := bson.IsObjectIdHex(id); !ok {
		err := errors.New("Invalid user id.")
		log.Println(c.SessionID, ": services : Users : Delete : Completed : ERROR :", err)
		return err
	}

	f := func(collection *mgo.Collection) error {
		q := bson.M{"_id": bson.ObjectIdHex(id)}
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
