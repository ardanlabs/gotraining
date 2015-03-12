// Package handlers contains the handler logic for processing requests.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArdanStudios/gotraining/12-http/api/app"
	"github.com/ArdanStudios/gotraining/12-http/api/models"
	"github.com/ArdanStudios/gotraining/12-http/api/services"
)

// List returns all the existing users in the system.
// 200 Success, 204 No Content, 500 Internal
func UsersList(c *app.Context) error {
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

		return nil
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : List : Completed : 200")
	return nil
}

// Retrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func UsersRetrieve(c *app.Context) error {
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

		return nil
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : Retrieve : Completed : 200")
	return nil
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func UsersCreate(c *app.Context) error {
	log.Println(c.SessionID, ": ctrls : Users : Create : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondError(err.Error(), http.StatusBadRequest)
		log.Println(c.SessionID, ": ctrls : Users : Create : Completed : 400 :", err)
		return nil
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

		return nil
	}

	c.Params = map[string]string{"id": u.UserID}
	UsersRetrieve(c)

	log.Println(c.SessionID, ": ctrls : Users : Create : Completed")
	return nil
}

// Update updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func UsersUpdate(c *app.Context) error {
	log.Println(c.SessionID, ": ctrls : Users : Update : Started")

	var u models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&u); err != nil {
		c.RespondError(err.Error(), http.StatusBadRequest)
		log.Println(c.SessionID, ": ctrls : Users : Update : Completed : 400 :", err)
		return nil
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

		return nil
	}

	UsersRetrieve(c)

	log.Println(c.SessionID, ": ctrls : Users : Update : Completed")
	return nil
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func UsersDelete(c *app.Context) error {
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

		return nil
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

		return nil
	}

	c.Respond(u, http.StatusOK)

	log.Println(c.SessionID, ": ctrls : Users : Delete : Completed")
	return nil
}
