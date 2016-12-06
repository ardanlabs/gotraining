// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/app"
)

// RequestLogger writes some information about the request to the logs in
// the format: SESSIONID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(h app.Handler) app.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(app.KeyValues).(*app.Values)

		start := time.Now()
		err := h(ctx, w, r, params)

		log.Printf("%s : RL : *****> %s %s -> %s (%s)",
			v.TraceID,
			r.Method, r.URL.Path,
			r.RemoteAddr, time.Since(start),
		)

		return err
	}
}
