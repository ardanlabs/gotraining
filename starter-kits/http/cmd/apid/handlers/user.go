// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"context"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
	"github.com/pkg/errors"
)

// UserList returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func UserList(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	u, err := user.List(ctx, reqDB)
	if err != nil {
		return errors.Wrap(err, "")
	}

	web.Respond(ctx, w, u, http.StatusOK)
	return nil
}

// UserRetrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func UserRetrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	u, err := user.Retrieve(ctx, reqDB, params["id"])
	if err != nil {
		return errors.Wrapf(err, "Id: %s", params["id"])
	}

	web.Respond(ctx, w, u, http.StatusOK)
	return nil
}

// UserCreate inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func UserCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	var cu user.CreateUser
	if err := web.Unmarshal(r.Body, &cu); err != nil {
		return errors.Wrap(err, "")
	}

	u, err := user.Create(ctx, reqDB, &cu)
	if err != nil {
		return errors.Wrapf(err, "User: %+v", &cu)
	}

	web.Respond(ctx, w, u, http.StatusCreated)
	return nil
}

// UserUpdate updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := MasterDB.MGOCopy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer reqDB.MGOClose()

	var cu user.CreateUser
	if err := web.Unmarshal(r.Body, &cu); err != nil {
		return errors.Wrap(err, "")
	}

	if err := user.Update(ctx, reqDB, params["id"], &cu); err != nil {
		return errors.Wrapf(err, "Id: %s  User: %+v", params["id"], &cu)
	}

	web.Respond(ctx, w, nil, http.StatusNoContent)
	return nil
}

// UserDelete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	reqDB, err := MasterDB.MGOCopy()
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
