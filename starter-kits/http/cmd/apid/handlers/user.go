// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/app"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
)

// userHandle maintains the set of handlers for the users api.
type userHandle struct{}

// User fronts the access to the users service functionality.
var User userHandle

// List returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func (userHandle) List(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.List(ctx, v.TraceID, v.DB)
	if err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}

// Retrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func (userHandle) Retrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.Retrieve(ctx, v.TraceID, v.DB, params["id"])
	if err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func (userHandle) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}

	if invld, err := user.Create(ctx, v.TraceID, v.DB, &u); err != nil {
		switch err {
		case app.ErrValidation:
			app.RespondInvalid(w, v.TraceID, invld)
			return nil

		default:
			return err
		}
	}

	params = map[string]string{"id": u.UserID}

	return User.Retrieve(ctx, w, r, params)
}

// Update updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func (userHandle) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}

	if invld, err := user.Update(ctx, v.TraceID, v.DB, params["id"], &u); err != nil {
		switch err {
		case app.ErrValidation:
			app.RespondInvalid(w, v.TraceID, invld)
			return nil

		default:
			return err
		}
	}

	return User.Retrieve(ctx, w, r, params)
}

// Delete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func (userHandle) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.Retrieve(ctx, v.TraceID, v.DB, params["id"])
	if err != nil {
		return err
	}

	if err := user.Delete(ctx, v.TraceID, v.DB, params["id"]); err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}
