// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"context"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
	"github.com/pkg/errors"
)

// User represents the User API method handler set.
type User struct {
	MasterDB *db.DB

	// ADD OTHER STATE LIKE THE LOGGER AND CONFIG HERE.
}

// List returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func (u *User) List(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := u.MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	usrs, err := user.List(ctx, reqDB)
	if err != nil {
		return errors.Wrap(err, "")
	}

	web.Respond(ctx, w, usrs, http.StatusOK)
	return nil
}

// Retrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func (u *User) Retrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := u.MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	usr, err := user.Retrieve(ctx, reqDB, params["id"])
	if err != nil {
		return errors.Wrapf(err, "Id: %s", params["id"])
	}

	web.Respond(ctx, w, usr, http.StatusOK)
	return nil
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func (u *User) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := u.MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	var usr user.CreateUser
	if err := web.Unmarshal(r.Body, &usr); err != nil {
		return errors.Wrap(err, "")
	}

	nUsr, err := user.Create(ctx, reqDB, &usr)
	if err != nil {
		return errors.Wrapf(err, "User: %+v", &usr)
	}

	web.Respond(ctx, w, nUsr, http.StatusCreated)
	return nil
}

// Update updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func (u *User) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := u.MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	var usr user.CreateUser
	if err := web.Unmarshal(r.Body, &usr); err != nil {
		return errors.Wrap(err, "")
	}

	if err := user.Update(ctx, reqDB, params["id"], &usr); err != nil {
		return errors.Wrapf(err, "Id: %s  User: %+v", params["id"], &usr)
	}

	web.Respond(ctx, w, nil, http.StatusNoContent)
	return nil
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func (u *User) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := u.MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	if err := user.Delete(ctx, reqDB, params["id"]); err != nil {
		return errors.Wrapf(err, "Id: %s", params["id"])
	}

	web.Respond(ctx, w, nil, http.StatusNoContent)
	return nil
}
