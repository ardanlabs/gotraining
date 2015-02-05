// Package models contains data structures and behavior for the data.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"gopkg.in/mgo.v2/bson"
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
	DateModified *time.Time `bson:"date_modified" json:"-"`
	DateCreated  *time.Time `bson:"date_created,omitempty" json:"-"`
}

// Validate checks the fields to verify the value is in a proper state.
func (ua *UserAddress) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	if ua.Type == 0 {
		inv = append(inv, app.Invalid{"Type", "The value of Type cannot be 0."})
	}

	if ua.LineOne == "" {
		inv = append(inv, app.Invalid{"LineOne", "A value of LineOne cannot be empty."})
	}

	if ua.City == "" {
		inv = append(inv, app.Invalid{"City", "A value of City cannot be empty."})
	}

	if ua.State == "" {
		inv = append(inv, app.Invalid{"State", "A value of State cannot be empty."})
	}

	if ua.Zipcode == "" {
		inv = append(inv, app.Invalid{"Zipcode", "A value of Zipcode cannot be empty."})
	}

	if ua.Phone == "" {
		inv = append(inv, app.Invalid{"Phone", "A value of Phone cannot be empty."})
	}

	if len(inv) > 0 {
		return inv, errors.New("Validation failures identified.")
	}

	return nil, nil
}

// Compare checks the fields against another UserAddress value.
func (ua *UserAddress) Compare(uat *UserAddress) ([]app.Invalid, error) {
	var inv []app.Invalid

	if ua.Type != uat.Type {
		inv = append(inv, app.Invalid{"Type", fmt.Sprintf("The value of Type is not the same. %s != %s", ua.Type, uat.Type)})
	}

	if ua.LineOne != uat.LineOne {
		inv = append(inv, app.Invalid{"LineOne", fmt.Sprintf("The value of LineOne is not the same. %s != %s", ua.LineOne, uat.LineOne)})
	}

	if ua.City != uat.City {
		inv = append(inv, app.Invalid{"City", fmt.Sprintf("The value of City is not the same. %s != %s", ua.City, uat.City)})
	}

	if ua.State != uat.State {
		inv = append(inv, app.Invalid{"State", fmt.Sprintf("The value of State is not the same. %s != %s", ua.State, uat.State)})
	}

	if ua.Zipcode != uat.Zipcode {
		inv = append(inv, app.Invalid{"Zipcode", fmt.Sprintf("The value of Zipcode is not the same. %s != %s", ua.Zipcode, uat.Zipcode)})
	}

	if ua.Phone != uat.Phone {
		inv = append(inv, app.Invalid{"Phone", fmt.Sprintf("The value of Phone is not the same. %s != %s", ua.Phone, uat.Phone)})
	}

	if len(inv) > 0 {
		return inv, errors.New("Compare failures identified.")
	}

	return nil, nil
}

// User contains information about a user.
type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	UserType     int           `bson:"type" json:"type"`
	FirstName    string        `bson:"first_name" json:"first_name"`
	LastName     string        `bson:"last_name" json:"last_name"`
	Email        string        `bson:"email" json:"email"`
	Company      string        `bson:"company" json:"company"`
	Addresses    []UserAddress `bson:"addresses" json:"addresses"`
	DateModified *time.Time    `bson:"date_modified" json:"-"`
	DateCreated  *time.Time    `bson:"date_created,omitempty" json:"-"`
}

// Validate checks the fields to verify the value is in a proper state.
func (u *User) Validate() ([]app.Invalid, error) {
	var inv []app.Invalid

	if u.UserType == 0 {
		inv = append(inv, app.Invalid{"UserType", "The value of UserType cannot be 0."})
	}

	if u.FirstName == "" {
		inv = append(inv, app.Invalid{"FirstName", "A value of FirstName cannot be empty."})
	}

	if u.LastName == "" {
		inv = append(inv, app.Invalid{"LastName", "A value of LastName cannot be empty."})
	}

	if u.Email == "" {
		inv = append(inv, app.Invalid{"Email", "A value of Email cannot be empty."})
	}

	if u.Company == "" {
		inv = append(inv, app.Invalid{"Company", "A value of Company cannot be empty."})
	}

	for _, ua := range u.Addresses {
		if va, err := ua.Validate(); err != nil {
			inv = append(inv, va...)
		}
	}

	if len(inv) > 0 {
		return inv, errors.New("Validation failures identified.")
	}

	return nil, nil
}

// Compare checks the fields against another User value.
func (u *User) Compare(ut *User) ([]app.Invalid, error) {
	var inv []app.Invalid

	if u.UserType != ut.UserType {
		inv = append(inv, app.Invalid{"UserType", fmt.Sprintf("The value of UserType is not the same. %s != %s", u.UserType, ut.UserType)})
	}

	if u.FirstName != ut.FirstName {
		inv = append(inv, app.Invalid{"FirstName", fmt.Sprintf("The value of FirstName is not the same. %s != %s", u.FirstName, ut.FirstName)})
	}

	if u.LastName != ut.LastName {
		inv = append(inv, app.Invalid{"LastName", fmt.Sprintf("The value of LastName is not the same. %s != %s", u.LastName, ut.LastName)})
	}

	if u.Email != ut.Email {
		inv = append(inv, app.Invalid{"Email", fmt.Sprintf("The value of Email is not the same. %s != %s", u.Email, ut.Email)})
	}

	if u.Company != ut.Company {
		inv = append(inv, app.Invalid{"Company", fmt.Sprintf("The value of Company is not the same. %s != %s", u.Company, ut.Company)})
	}

	uLen := len(u.Addresses)
	utLen := len(ut.Addresses)

	if uLen != utLen {
		inv = append(inv, app.Invalid{"Addresses", fmt.Sprintf("The set of Addresses is not the same. %s != %s", uLen, utLen)})
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
		return inv, errors.New("Compare failures identified.")
	}

	return nil, nil
}
