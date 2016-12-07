// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user

import "time"

// CreateAddress contains information about a user's address.
type CreateAddress struct {
	Type         int        `bson:"type" json:"type" validate:"required"`
	LineOne      string     `bson:"line_one" json:"line_one" validate:"required"`
	LineTwo      string     `bson:"line_two" json:"line_two,omitempty"`
	City         string     `bson:"city" json:"city" validate:"required"`
	State        string     `bson:"state" json:"state" validate:"required"`
	Zipcode      string     `bson:"zipcode" json:"zipcode" validate:"required"`
	Phone        string     `bson:"phone" json:"phone" validate:"required"`
	DateModified *time.Time `bson:"date_modified" json:"-"`
}

// Address contains information about a user's address.
type Address struct {
	Type         int        `bson:"type" json:"type" validate:"required"`
	LineOne      string     `bson:"line_one" json:"line_one" validate:"required"`
	LineTwo      string     `bson:"line_two" json:"line_two,omitempty"`
	City         string     `bson:"city" json:"city" validate:"required"`
	State        string     `bson:"state" json:"state" validate:"required"`
	Zipcode      string     `bson:"zipcode" json:"zipcode" validate:"required"`
	Phone        string     `bson:"phone" json:"phone" validate:"required"`
	DateModified *time.Time `bson:"date_modified" json:"date_modified"`
	DateCreated  *time.Time `bson:"date_created,omitempty" json:"date_created"`
}

//==============================================================================

// CreateUser contains information about a user.
type CreateUser struct {
	UserType     int             `bson:"type" json:"type" validate:"required"`
	FirstName    string          `bson:"first_name" json:"first_name" validate:"required"`
	LastName     string          `bson:"last_name" json:"last_name" validate:"required"`
	Email        string          `bson:"email" json:"email" validate:"required"`
	Company      string          `bson:"company" json:"company" validate:"required"`
	Addresses    []CreateAddress `bson:"addresses" json:"addresses" validate:"required"`
	DateModified *time.Time      `bson:"date_modified" json:"-"`
}

// User contains information about a user.
type User struct {
	UserID       string     `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserType     int        `bson:"type" json:"type"`
	FirstName    string     `bson:"first_name" json:"first_name"`
	LastName     string     `bson:"last_name" json:"last_name"`
	Email        string     `bson:"email" json:"email"`
	Company      string     `bson:"company" json:"company"`
	Addresses    []Address  `bson:"addresses" json:"addresses"`
	DateModified *time.Time `bson:"date_modified" json:"date_modified"`
	DateCreated  *time.Time `bson:"date_created,omitempty" json:"date_created"`
}
