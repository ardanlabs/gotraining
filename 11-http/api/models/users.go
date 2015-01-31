// Package models contains data structures and behavior for the data.
package models

import (
	"log"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserAddress contains information about a user's address.
type UserAddress struct {
	Type         int       `bson:"type" json:"type"`
	LineOne      string    `bson:"line_one" json:"line_one"`
	LineTwo      string    `bson:"line_two" json:"line_two,omitempty"`
	City         string    `bson:"city" json:"city"`
	State        string    `bson:"state" json:"state"`
	Zipcode      string    `bson:"zipcode" json:"zipcode"`
	Phone        string    `bson:"phone" json:"phone"`
	DateModified time.Time `bson:"date_modified" json:"-"`
	DateCreated  time.Time `bson:"date_created" json:"-"`
}

// Validate checks the fields for verify the value is in a proper state.
func (ua *UserAddress) Validate() ([]app.Invalid, error) {
	// Check fields
	return nil, nil
}

// User contains information about a user.
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"-"`
	UserType     int           `bson:"type" json:"type"`
	FirstName    string        `bson:"first_name" json:"first_name"`
	LastName     string        `bson:"last_name" json:"last_name"`
	Email        string        `bson:"email" json:"email"`
	PostalCode   string        `bson:"postal_code,omitempty" json:"postal_code"`
	CompanyID    string        `bson:"company_id" json:"company_id"`
	Addresses    []UserAddress `bson:"addresses" json:"addresses"`
	DateModified time.Time     `bson:"date_modified" json:"-"`
	DateCreated  time.Time     `bson:"date_created" json:"-"`
}

// Validate checks the fields for verify the value is in a proper state.
func (u *User) Validate() ([]app.Invalid, error) {
	var v []app.Invalid

	// Check fields

	for _, ua := range u.Addresses {
		if va, err := ua.Validate(); err != nil {
			v = append(v, va...)
		}
	}

	return nil, nil
}

// UsersList retrieves a list of existing users from the database.
func UsersList(c *app.Context) ([]User, error) {
	log.Println(c.SessionID, ": services : UsersList : Started")

	var users []User
	f := func(collection *mgo.Collection) error {
		log.Println(c.SessionID, ": services : UsersList: MGO :\n\ndb.users.find()")
		return collection.Find(nil).All(&users)
	}

	if err := app.ExecuteDB(c.Session, app.DB, f); err != nil {
		log.Println(c.SessionID, ": services : UsersList : Completed : ERROR :", err)
		return nil, err
	}

	log.Println(c.SessionID, ": services : UsersList : Completed")
	return users, nil
}

// CreateUser inserts a new user into the database.
func (u *User) Create(c *app.Context) error {
	log.Println(c.SessionID, ": services : CreateUser : Started")

	u.ID = bson.NewObjectId()

	f := func(collection *mgo.Collection) error {
		log.Printf("%s : services : CreateUser: MGO :\n\ndb.users.insert(%s)\n", c.SessionID, app.Query(u))
		return collection.Insert(u)
	}

	if err := app.ExecuteDB(c.Session, app.DB, f); err != nil {
		log.Println(c.SessionID, ": services : CreateUser : Completed : ERROR :", err)
		return err
	}

	log.Println(c.SessionID, ": services : CreateUser : Completed")
	return nil
}
