package middleware

import (
	"log"
	"time"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
)

// RequestLogger writes some information about the request to the logs in
// the format: SESSIONID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(h app.Handler) app.Handler {
	return func(c *app.Context) error {
		start := time.Now()
		err := h(c)

		log.Printf("%s : RL : *****> (%d) %s %s -> %s (%s)",
			c.SessionID,
			c.Status, c.Request.Method, c.Request.URL.Path,
			c.Request.RemoteAddr, time.Since(start),
		)

		return err
	}
}
