// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
)

// CORS providing support for Cross-Origin Resource Sharing.
// https://metajack.im/2010/01/19/crossdomain-ajax-for-xmpp-http-binding-made-easy/
func CORS(a *web.App, origin string, methods string) web.Middleware {

	// Create the options request handler which will attach CORS options to it.
	a.TreeMux.OptionsHandler = func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
	}

	// Return this middleware to be chained together.
	return func(next web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			v := ctx.Value(web.KeyValues).(*web.Values)

			// Add the access control to the header.
			w.Header().Set("Access-Control-Allow-Origin", origin)
			log.Printf("%s : CORS : Access Control Allowed : Origin[%s] Methods[%s]", v.TraceID, origin, methods)

			return next(ctx, w, r, params)
		}
	}
}
