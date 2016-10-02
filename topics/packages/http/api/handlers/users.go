// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/models"
	"github.com/ardanlabs/gotraining/topics/packages/http/api/services"
)

// usersHandle maintains the set of handlers for the users api.
type usersHandle struct{}

// Users fronts the access to the users service functionality.
var Users usersHandle

// List returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func (usersHandle) List(c *app.Context) error {
	u, err := services.Users.List(c)
	if err != nil {
		return err
	}

	c.Respond(u, http.StatusOK)
	return nil
}

// Retrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func (usersHandle) Retrieve(c *app.Context) error {
	u, err := services.Users.Retrieve(c, c.Params["id"])
	if err != nil {
		return err
	}

	c.Respond(u, http.StatusOK)
	return nil
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func (usersHandle) Create(c *app.Context) error {
	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		return err
	}

	if v, err := services.Users.Create(c, &u); err != nil {
		switch err {
		case app.ErrValidation:
			c.RespondInvalid(v)
			return nil

		default:
			return err
		}
	}

	c.Params = map[string]string{"id": u.UserID}
	return Users.Retrieve(c)
}

// Update updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func (usersHandle) Update(c *app.Context) error {
	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		return err
	}

	if v, err := services.Users.Update(c, c.Params["id"], &u); err != nil {
		switch err {
		case app.ErrValidation:
			c.RespondInvalid(v)
			return nil

		default:
			return err
		}
	}

	return Users.Retrieve(c)
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func (usersHandle) Delete(c *app.Context) error {
	u, err := services.Users.Retrieve(c, c.Params["id"])
	if err != nil {
		return err
	}

	if err := services.Users.Delete(c, c.Params["id"]); err != nil {
		return err
	}

	c.Respond(u, http.StatusOK)
	return nil
}
