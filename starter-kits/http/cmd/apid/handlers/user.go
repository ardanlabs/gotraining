// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"context"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
)

// UserList returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func UserList(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(web.KeyValues).(*web.Values)

	u, err := user.List(ctx, v.TraceID, v.DB)
	if err != nil {
		return err
	}

	web.Respond(ctx, w, v.TraceID, u, http.StatusOK)
	return nil
}

// UserRetrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func UserRetrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(web.KeyValues).(*web.Values)

	u, err := user.Retrieve(ctx, v.TraceID, v.DB, params["id"])
	if err != nil {
		return err
	}

	web.Respond(ctx, w, v.TraceID, u, http.StatusOK)
	return nil
}

// UserCreate inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func UserCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(web.KeyValues).(*web.Values)

	var cu user.CreateUser
	if err := web.Unmarshal(r.Body, &cu); err != nil {
		return err
	}

	u, err := user.Create(ctx, v.TraceID, v.DB, &cu)
	if err != nil {
		return err
	}

	web.Respond(ctx, w, v.TraceID, u, http.StatusCreated)
	return nil
}

// UserUpdate updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(web.KeyValues).(*web.Values)

	var cu user.CreateUser
	if err := web.Unmarshal(r.Body, &cu); err != nil {
		return err
	}

	if err := user.Update(ctx, v.TraceID, v.DB, params["id"], &cu); err != nil {
		return err
	}

	web.Respond(ctx, w, v.TraceID, nil, http.StatusNoContent)
	return nil
}

// UserDelete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(web.KeyValues).(*web.Values)

	if err := user.Delete(ctx, v.TraceID, v.DB, params["id"]); err != nil {
		return err
	}

	web.Respond(ctx, w, v.TraceID, nil, http.StatusNoContent)
	return nil
}
