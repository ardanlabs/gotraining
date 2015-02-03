// Package services provides business and data processing.
package services

import (
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const usersCollection = "users"

// UsersList retrieves a list of existing users from the database.
func UsersList(c *app.Context) ([]models.User, error) {
	log.Println(c.SessionID, ": services : UsersList : Started")

	var users []models.User
	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : UsersList: MGO :\n\ndb.users.find()\n\n", c.SessionID)
		return collection.Find(nil).All(&users)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : UsersList : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : UsersList : Completed")
	return users, nil
}

// UsersRetrieve gets the specified user from the database.
func UsersRetrieve(c *app.Context, id bson.ObjectId) (*models.User, error) {
	log.Println(c.SessionID, ": services : UsersRetrieve : Started")

	var u *models.User
	f := func(collection *mgo.Collection) error {
		q := bson.M{"_id": id}
		log.Printf("%s : services : UsersRetrieve: MGO :\n\ndb.users.find(%s)\n\n", c.SessionID, app.Query(q))
		return collection.Find(q).One(&u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : UsersRetrieve : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : UsersRetrieve : Completed")
	return u, nil
}

// UsersCreate inserts a new user into the database.
func UsersCreate(c *app.Context, u *models.User) error {
	log.Println(c.SessionID, ": services : UsersCreate : Started")

	u.ID = bson.NewObjectId()

	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : UsersCreate : MGO :\n\ndb.users.insert(%s)\n\n", c.SessionID, app.Query(u))
		return collection.Insert(u)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : UsersCreate : Completed : ERROR :", err)
		return err
	}

	log.Println(c.SessionID, ": services : UsersCreate : Completed")
	return nil
}

// UsersDelete inserts a new user into the database.
func UsersDelete(c *app.Context, id bson.ObjectId) error {
	log.Println(c.SessionID, ": services : UsersDelete : Started")

	f := func(collection *mgo.Collection) error {
		q := bson.M{"_id": id}
		log.Printf("%s : services : UsersRetrieve: MGO :\n\ndb.users.remove(%s)\n\n", c.SessionID, app.Query(q))
		return collection.Remove(q)
	}

	if err := app.ExecuteDB(c.Session, usersCollection, f); err != nil {
		log.Println(c.SessionID, ": services : UsersDelete : Completed : ERROR :", err)
		return err
	}

	log.Println(c.SessionID, ": services : UsersDelete : Completed")
	return nil
}
