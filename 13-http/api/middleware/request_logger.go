package middleware

import (
	"log"
	"time"

	"github.com/ardanlabs/gotraining/13-http/api/app"
)

// RequestLogger writes some information about the request to the logs in
// the format: SESSIONID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(h app.Handler) app.Handler {
	return func(c *app.Context) error {
		log.Printf("%v : middleware : RequestLogger : Started", c.SessionID)

		start := time.Now()
		err := h(c)

		log.Printf("%v : (%d) %s %s -> %s (%s)",
			c.SessionID,
			c.Status, c.Request.Method, c.Request.URL.Path,
			c.Request.RemoteAddr, time.Since(start),
		)

		log.Printf("%v : middleware : RequestLogger : Completed", c.SessionID)

		return err
	}
}
