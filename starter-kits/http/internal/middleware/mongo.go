// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/kit/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/kit/web"
)

// Mongo initializes the master session and wires in the connection middleware.
func Mongo() web.Middleware {

	// session contains the master session for accessing MongoDB.
	session, err := db.Init()
	if err != nil {
		log.Fatalf("startup : Mongo : Initialize Mongo : %+v\n", err)
	}

	// Return this middleware to be chained together.
	return func(next web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			v := ctx.Value(web.KeyValues).(*web.Values)

			// Get a MongoDB session connection.
			v.DB = session.Copy()
			defer v.DB.Close()

			if err := next(ctx, w, r, params); err != nil {
				return err
			}
			return nil
		}
	}
}
