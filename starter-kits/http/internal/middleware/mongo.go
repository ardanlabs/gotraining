// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/db"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
)

// Mongo initializes the master session and wires in the connection middleware.
func Mongo() web.Middleware {

	// session contains the master session for accessing MongoDB.
	session := db.Init()

	// Return this middleware to be chained together.
	return func(next web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			v := ctx.Value(web.KeyValues).(*web.Values)

			// Get a MongoDB session connection.
			log.Printf("%s : Mongo : *****> Capture Mongo Session\n", v.TraceID)
			v.DB = session.Copy()

			// Defer releasing the db session connection.
			defer func() {
				log.Printf("%s : Mongo : *****> Release Mongo Session\n", v.TraceID)
				v.DB.Close()
			}()

			return next(ctx, w, r, params)
		}
	}
}
