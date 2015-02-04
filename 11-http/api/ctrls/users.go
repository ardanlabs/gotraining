// Package ctrls contains the controller logic for processing requests.
package ctrls

import (
	"encoding/json"
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/services"
)

// UsersList returns all the existing users in the system.
// 200 Success, 500 Internal
func UsersList(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UsersList : Started")

	u, err := services.Users.List(c)
	if err != nil {
		c.RespondInternal500(err)
		log.Println(c.SessionID, ": ctrls : UsersList : Completed : 500 :", err)
		return
	}

	c.RespondSuccess200(u)

	log.Println(c.SessionID, ": ctrls : UsersList : Completed : 200")
}

// UsersCreate inserts a new user into the system.
// 200 Success, 409 Validation, 500 Internal
func UsersCreate(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UsersCreate : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondBadRequest400(err)
		log.Println(c.SessionID, ": ctrls : UsersCreate : Completed : 500 :", err)
		return
	}

	if v, err := u.Validate(); err != nil {
		c.RespondValidation409(v)
		log.Println(c.SessionID, ": ctrls : UsersCreate : Completed : 409 :", err)
		return
	}

	if err := services.Users.Create(c, &u); err != nil {
		c.RespondInternal500(err)
		log.Println(c.SessionID, ": ctrls : UsersCreate : Completed : 500 :", err)
		return
	}

	r := struct {
		ID string
	}{
		u.ID.Hex(),
	}

	c.RespondSuccess200(&r)

	log.Println(c.SessionID, ": ctrls : UsersCreate : Completed : 200")
}

// UsersRetrieve returns the specified user from the system.
// 200 Success, 404 Not Found, 500 Internal
func UsersRetrieve(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UsersRetrieve : Started")

	log.Println(c.SessionID, ": ctrls : UsersRetrieve : Completed : 200")
}

// UsersUpdate updates the specified user in the system.
// 200 Success, 409 Validation, 500 Internal
func UsersUpdate(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UsersUpdate : Started")

	log.Println(c.SessionID, ": ctrls : UsersUpdate : Completed : 200")
}

// UsersDelete removed the specified user from the system.
// 200 Success, 404 Not Found, 500 Internal
func UsersDelete(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : UsersDelete : Started")

	log.Println(c.SessionID, ": ctrls : UsersDelete : Completed : 200")
}
