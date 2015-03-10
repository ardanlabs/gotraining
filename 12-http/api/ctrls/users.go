// Package ctrls contains the controller logic for processing requests.
package ctrls

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/models"
	"github.com/ArdanStudios/gotraining/12-http/api/services"
)

// usersCtrl maintains the set of controllers for the users api.
type usersCtrl struct{}

// Users fronts the access to the users controller functionality.
var Users usersCtrl

// List returns all the existing users in the system.
// 200 Success, 204 No Content, 500 Internal
func (uc usersCtrl) List(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : List : Started")

	u, err := services.Users.List(c)
	if err != nil {
		switch err {
		case services.ErrNotFound:
			c.Respond(nil, http.StatusNoContent)
			log.Println(c.SessionID, ": ctrls : Users : List : Completed : 204 :", err)

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : List : Completed : 500 :", err)
		}

		return
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : List : Completed : 200")
}

// Retrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func (uc usersCtrl) Retrieve(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Retrieve : Started")

	u, err := services.Users.Retrieve(c, c.Params["id"])
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			c.RespondError(err.Error(), http.StatusBadRequest)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 400 :", err)

		case services.ErrNotFound:
			c.RespondError(err.Error(), http.StatusNotFound)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 404 : Not Found")

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 500 :", err)
		}

		return
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 200")
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func (uc usersCtrl) Create(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Create : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondError(err.Error(), http.StatusBadRequest)
		log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 400 :", err)
		return
	}

	if v, err := services.Users.Create(c, &u); err != nil {
		switch err {
		case services.ErrValidation:
			c.RespondInvalid(v)
			log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 400 :", err)

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 500 :", err)
		}

		return
	}

	c.Params = map[string]string{"id": u.UserID}
	uc.Retrieve(c)

	log.Println(c.SessionID, ": ctrls : Users : Create : Completed")
}

// Update updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func (uc usersCtrl) Update(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Update : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondError(err.Error(), http.StatusBadRequest)
		log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 400 :", err)
		return
	}

	if v, err := services.Users.Update(c, c.Params["id"], &u); err != nil {
		switch err {
		case services.ErrValidation:
			c.RespondInvalid(v)
			log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 400 :", err)

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 500 :", err)
		}

		return
	}

	uc.Retrieve(c)

	log.Println(c.SessionID, ": ctrls : Users : Update : Completed")
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func (uc usersCtrl) Delete(c *app.Context) {
	log.Println(c.SessionID, ": ctrls : Users : Delete : Started")

	u, err := services.Users.Retrieve(c, c.Params["id"])
	if err != nil {
		switch err {
		case services.ErrInvalidID:
			c.RespondError(err.Error(), http.StatusBadRequest)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 400 :", err)

		case services.ErrNotFound:
			c.RespondError(err.Error(), http.StatusNotFound)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 404 : Not Found")

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 500 :", err)
		}

		return
	}

	if err := services.Users.Delete(c, c.Params["id"]); err != nil {
		switch err {
		case services.ErrInvalidID:
			c.RespondError(err.Error(), http.StatusBadRequest)
			log.Println(c.SessionID, ": ctrls : Users : Delete : Completed : 400 :", err)

		default:
			c.RespondError(err.Error(), http.StatusInternalServerError)
			log.Println(c.SessionID, ": ctrls : Users : Delete : Completed : 500 :", err)
		}

		return
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : Delete : Completed")
}
