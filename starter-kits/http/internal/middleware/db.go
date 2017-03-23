// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
)

// Mongo initializes the master session and wires in the connection middleware.
func Mongo() web.Middleware {

	// Return this middleware to be chained together.
	return func(next web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			v := ctx.Value(web.KeyValues).(*web.Values)

			var err error
			if v.DB, err = db.New("got"); err != nil {
				return err
			}
			defer v.DB.Close()

			if err := next(ctx, w, r, params); err != nil {
				return err
			}
			return nil
		}
	}
}
