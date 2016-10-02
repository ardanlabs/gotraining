// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
)

// UserAddress contains information about a user's address.
type UserAddress struct {
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

// Validate checks the fields to verify the value is in a proper state.
func (ua *UserAddress) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	errs := validate.Struct(ua)
	if errs != nil {
		for _, err := range errs {
			inv = append(inv, app.Invalid{Fld: err.Field, Err: err.Tag})
		}

		return inv, errors.New("Validation failures identified")
	}

	return nil, nil
}

// Compare checks the fields against another UserAddress value.
func (ua *UserAddress) Compare(uat *UserAddress) ([]app.Invalid, error) {
	var inv []app.Invalid

	if ua.Type != uat.Type {
		inv = append(inv, app.Invalid{Fld: "Type", Err: fmt.Sprintf("The value of Type is not the same. %d != %d", ua.Type, uat.Type)})
	}

	if ua.LineOne != uat.LineOne {
		inv = append(inv, app.Invalid{Fld: "LineOne", Err: fmt.Sprintf("The value of LineOne is not the same. %s != %s", ua.LineOne, uat.LineOne)})
	}

	if ua.City != uat.City {
		inv = append(inv, app.Invalid{Fld: "City", Err: fmt.Sprintf("The value of City is not the same. %s != %s", ua.City, uat.City)})
	}

	if ua.State != uat.State {
		inv = append(inv, app.Invalid{Fld: "State", Err: fmt.Sprintf("The value of State is not the same. %s != %s", ua.State, uat.State)})
	}

	if ua.Zipcode != uat.Zipcode {
		inv = append(inv, app.Invalid{Fld: "Zipcode", Err: fmt.Sprintf("The value of Zipcode is not the same. %s != %s", ua.Zipcode, uat.Zipcode)})
	}

	if ua.Phone != uat.Phone {
		inv = append(inv, app.Invalid{Fld: "Phone", Err: fmt.Sprintf("The value of Phone is not the same. %s != %s", ua.Phone, uat.Phone)})
	}

	if len(inv) > 0 {
		return inv, errors.New("Compare failures identified")
	}

	return nil, nil
}

// User contains information about a user.
type User struct {
	UserID       string        `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserType     int           `bson:"type" json:"type" validate:"required"`
	FirstName    string        `bson:"first_name" json:"first_name" validate:"required"`
	LastName     string        `bson:"last_name" json:"last_name" validate:"required"`
	Email        string        `bson:"email" json:"email" validate:"required"`
	Company      string        `bson:"company" json:"company" validate:"required"`
	Addresses    []UserAddress `bson:"addresses" json:"addresses" validate:"required"`
	DateModified *time.Time    `bson:"date_modified" json:"date_modified"`
	DateCreated  *time.Time    `bson:"date_created,omitempty" json:"date_created"`
}

// Validate checks the fields to verify the value is in a proper state.
func (u *User) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	errs := validate.Struct(u)
	if errs != nil {
		for _, err := range errs {
			inv = append(inv, app.Invalid{Fld: err.Field, Err: err.Tag})
		}

		return inv, errors.New("Validation failures identified")
	}

	for _, ua := range u.Addresses {
		if va, err := ua.Validate(); err != nil {
			inv = append(inv, va...)
		}
	}

	if len(inv) > 0 {
		return inv, errors.New("Validation failures identified")
	}

	return nil, nil
}

// Compare checks the fields against another User value.
func (u *User) Compare(ut *User) ([]app.Invalid, error) {
	var inv []app.Invalid

	if u.UserType != ut.UserType {
		inv = append(inv, app.Invalid{Fld: "UserType", Err: fmt.Sprintf("The value of UserType is not the same. %d != %d", u.UserType, ut.UserType)})
	}

	if u.FirstName != ut.FirstName {
		inv = append(inv, app.Invalid{Fld: "FirstName", Err: fmt.Sprintf("The value of FirstName is not the same. %s != %s", u.FirstName, ut.FirstName)})
	}

	if u.LastName != ut.LastName {
		inv = append(inv, app.Invalid{Fld: "LastName", Err: fmt.Sprintf("The value of LastName is not the same. %s != %s", u.LastName, ut.LastName)})
	}

	if u.Email != ut.Email {
		inv = append(inv, app.Invalid{Fld: "Email", Err: fmt.Sprintf("The value of Email is not the same. %s != %s", u.Email, ut.Email)})
	}

	if u.Company != ut.Company {
		inv = append(inv, app.Invalid{Fld: "Company", Err: fmt.Sprintf("The value of Company is not the same. %s != %s", u.Company, ut.Company)})
	}

	uLen := len(u.Addresses)
	utLen := len(ut.Addresses)

	if uLen != utLen {
		inv = append(inv, app.Invalid{Fld: "Addresses", Err: fmt.Sprintf("The set of Addresses is not the same. %d != %d", uLen, utLen)})
	}

	for idx, ua := range u.Addresses {
		if idx >= utLen {
			break
		}

		if va, err := ua.Compare(&ut.Addresses[idx]); err != nil {
			inv = append(inv, va...)
		}
	}

	if len(inv) > 0 {
		return inv, errors.New("Compare failures identified")
	}

	return nil, nil
}
