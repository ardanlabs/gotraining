// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
)

// RequestLogger writes some information about the request to the logs in
// the format: TraceID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(web.KeyValues).(*web.Values)

		start := time.Now()
		err := next(ctx, w, r, params)

		log.Printf("%s : (%d) : %s %s -> %s (%s) : Error[%v]",
			v.TraceID,
			v.StatusCode,
			r.Method, r.URL.Path,
			r.RemoteAddr, time.Since(start),
			err,
		)

		return err
	}
}
