// Package ctrls contains the controller logic for processing requests.
package ctrls

import (
	"encoding/json"
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
)

// ListUsers returns all the existing users in the system.
// 200 Success, 500 Internal
func ListUsers(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : ListUsers : Started")

	u, err := models.ListUsers(c)
	if err != nil {
		c.RespondInternal500(err)
		log.Println(c.SessionID, ": ctrls : ListUsers : Completed : 500 :", err)
		return
	}

	c.RespondSuccess200(u)

	log.Println(c.SessionID, ": ctrls : ListUsers : Completed : 200")
}

// CreateUser inserts a new user into the system.
// 200 Success, 409 Validation, 500 Internal
func CreateUser(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : CreateUser : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondBadRequest400(err)
		log.Println(c.SessionID, ": ctrls : CreateUser : Completed : 500 :", err)
		return
	}

	if v, err := u.Validate(); err != nil {
		c.RespondValidation409(v)
		log.Println(c.SessionID, ": ctrls : CreateUser : Completed : 409 :", err)
		return
	}

	if err := u.Create(c); err != nil {
		c.RespondInternal500(err)
		log.Println(c.SessionID, ": ctrls : CreateUser : Completed : 500 :", err)
		return
	}

	r := struct {
		ID string
	}{
		u.ID.Hex(),
	}

	c.RespondSuccess200(&r)

	log.Println(c.SessionID, ": ctrls : CreateUser : Completed : 200")
}

// ShowUser returns the specified user from the system.
// 200 Success, 404 Not Found, 500 Internal
func ShowUser(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : ShowUser : Started")

	log.Println(c.SessionID, ": ctrls : ShowUser : Completed : 200")
}

// UpdateUser updates the specified user in the system.
// 200 Success, 409 Validation, 500 Internal
func UpdateUser(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UpdateUser : Started")

	log.Println(c.SessionID, ": ctrls : UpdateUser : Completed : 200")
}

// DeleteUser removed the specified user from the system.
// 200 Success, 404 Not Found, 500 Internal
func DeleteUser(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : DeleteUser : Started")

	log.Println(c.SessionID, ": ctrls : DeleteUser : Completed : 200")
}
