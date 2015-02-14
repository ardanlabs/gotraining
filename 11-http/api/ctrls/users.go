// Package ctrls contains the controller logic for processing requests.
package ctrls

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/app"
	"github.com/ArdanStudios/gotraining/11-http/api/models"
	"github.com/ArdanStudios/gotraining/11-http/api/services"
)

// usersCtrl maintains the set of controllers for the users api.
type usersCtrl struct{}

// Users fronts the access to the users controller functionality.
var Users usersCtrl

// UsersList returns all the existing users in the system.
// 200 Success, 204 No Content, 500 Internal
func (uc usersCtrl) List(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : List : Started")

	u, err := services.Users.List(c)
	if err != nil {
		switch err {
		case services.ErrNotFound:
			c.RespondNoContent204()
			log.Println(c.SessionID, ": ctrls : Users : List : Completed : 204 :", err)

		default:
			c.RespondInternal500(err)
			log.Println(c.SessionID, ": ctrls : Users : List : Completed : 500 :", err)
		}

		return
	}

	c.RespondSuccess200(u)

	log.Println(c.SessionID, ": ctrls : Users : List : Completed : 200")
}

// UsersCreate inserts a new user into the system.
// 200 Success, 409 validation, 500 Internal
func (uc usersCtrl) Create(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Create : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondBadRequest400(err)
		log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 400 :", err)
		return
	}

	if v, err := services.Users.Create(c, &u); err != nil {
		switch err {
		case services.ErrValidation:
			c.RespondValidation400(v)
			log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 400 :", err)

		default:
			c.RespondInternal500(err)
			log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 500 :", err)
		}

		return
	}

	r := struct {
		UserID string `json:"user_id"`
	}{
		u.UserID,
	}

	c.RespondSuccess200(&r)

	log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 200")
}

// UsersRetrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func (uc usersCtrl) Retrieve(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Retrieve : Started")

	u, err := services.Users.Retrieve(c, c.Params["id"])
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			c.RespondValidation400([]app.Invalid{{Fld: "id", Err: err.Error()}})
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 400 :", err)

		case services.ErrNotFound:
			c.RespondNotFound404()
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 404 : Not Found")

		default:
			c.RespondInternal500(err)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 500 :", err)
		}

		return
	}

	c.RespondSuccess200(&u)

	log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 200")
}

// UsersUpdate updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func (uc usersCtrl) Update(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Update : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondBadRequest400(err)
		log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 400 :", err)
		return
	}

	if v, err := services.Users.Update(c, c.Params["id"], &u); err != nil {
		switch err {
		case services.ErrValidation:
			c.RespondValidation400(v)
			log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 400 :", err)

		default:
			c.RespondInternal500(err)
			log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 500 :", err)
		}

		return
	}

	r := struct {
		Message string `json:"message"`
	}{
		fmt.Sprintf("User with ID %s has been updated.", u.UserID),
	}

	c.RespondSuccess200(&r)

	log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 200")
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func (uc usersCtrl) Delete(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Delete : Started")

	if err := services.Users.Delete(c, c.Params["id"]); err != nil {
		switch err {
		case services.ErrInvalidID:
			c.RespondValidation400([]app.Invalid{{Fld: "id", Err: err.Error()}})
			log.Println(c.SessionID, ": ctrls : Users : Delete : Completed : 400 :", err)

		default:
			c.RespondInternal500(err)
			log.Println(c.SessionID, ": ctrls : Users : Delete : Completed : 500 :", err)
		}
	}

	r := struct {
		Message string `json:"message"`
	}{
		fmt.Sprintf("User with ID %s has been removed.", c.Params["id"]),
	}

	c.RespondSuccess200(&r)

	log.Println(c.SessionID, ": ctrls : Users : Delete : Completed : 200")
}
