// Package models contains data structures and associated behavior
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
)

// UserAddress contains information about a user's address.
type UserAddress struct {
	Type         int        `bson:"type" json:"type"`
	LineOne      string     `bson:"line_one" json:"line_one"`
	LineTwo      string     `bson:"line_two" json:"line_two,omitempty"`
	City         string     `bson:"city" json:"city"`
	State        string     `bson:"state" json:"state"`
	Zipcode      string     `bson:"zipcode" json:"zipcode"`
	Phone        string     `bson:"phone" json:"phone"`
	DateModified *time.Time `bson:"date_modified" json:"date_modified"`
	DateCreated  *time.Time `bson:"date_created,omitempty" json:"date_created"`
}

// Validate checks the fields to verify the value is in a proper state.
func (ua *UserAddress) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	if ua.Type == 0 {
		inv = append(inv, app.Invalid{Fld: "Type", Err: "The value of Type cannot be 0."})
	}

	if ua.LineOne == "" {
		inv = append(inv, app.Invalid{Fld: "LineOne", Err: "A value of LineOne cannot be empty."})
	}

	if ua.City == "" {
		inv = append(inv, app.Invalid{Fld: "City", Err: "A value of City cannot be empty."})
	}

	if ua.State == "" {
		inv = append(inv, app.Invalid{Fld: "State", Err: "A value of State cannot be empty."})
	}

	if ua.Zipcode == "" {
		inv = append(inv, app.Invalid{Fld: "Zipcode", Err: "A value of Zipcode cannot be empty."})
	}

	if ua.Phone == "" {
		inv = append(inv, app.Invalid{Fld: "Phone", Err: "A value of Phone cannot be empty."})
	}

	if len(inv) > 0 {
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
	UserType     int           `bson:"type" json:"type"`
	FirstName    string        `bson:"first_name" json:"first_name"`
	LastName     string        `bson:"last_name" json:"last_name"`
	Email        string        `bson:"email" json:"email"`
	Company      string        `bson:"company" json:"company"`
	Addresses    []UserAddress `bson:"addresses" json:"addresses"`
	DateModified *time.Time    `bson:"date_modified" json:"date_modified"`
	DateCreated  *time.Time    `bson:"date_created,omitempty" json:"date_created"`
}

// Validate checks the fields to verify the value is in a proper state.
func (u *User) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	if u.UserType == 0 {
		inv = append(inv, app.Invalid{Fld: "UserType", Err: "The value of UserType cannot be 0."})
	}

	if u.FirstName == "" {
		inv = append(inv, app.Invalid{Fld: "FirstName", Err: "A value of FirstName cannot be empty."})
	}

	if u.LastName == "" {
		inv = append(inv, app.Invalid{Fld: "LastName", Err: "A value of LastName cannot be empty."})
	}

	if u.Email == "" {
		inv = append(inv, app.Invalid{Fld: "Email", Err: "A value of Email cannot be empty."})
	}

	if u.Company == "" {
		inv = append(inv, app.Invalid{Fld: "Company", Err: "A value of Company cannot be empty."})
	}

	if len(u.Addresses) == 0 {
		inv = append(inv, app.Invalid{Fld: "Addresses", Err: "There must be at least one address."})
	} else {
		for _, ua := range u.Addresses {
			if va, err := ua.Validate(); err != nil {
				inv = append(inv, va...)
			}
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
