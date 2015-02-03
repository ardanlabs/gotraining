// Package models contains data structures and behavior for the data.
package models

import (
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
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
	Company      string        `bson:"company" json:"company"`
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
