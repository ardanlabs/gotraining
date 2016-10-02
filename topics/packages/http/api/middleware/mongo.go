package middleware

import (
	"log"

	"github.com/ardanlabs/gotraining/topics/packages/http/api/app"
)

// Mongo handles session management.
func Mongo(h app.Handler) app.Handler {

	// Wrap the handlers inside a session copy/close.
	f := func(c *app.Context) error {

		log.Printf("%s : Mongo : *****> Capture Mongo Session\n", c.SessionID)
		ses := app.GetSession()
		c.Ctx["DB"] = ses

		defer func() {
			log.Printf("%s : Mongo : *****> Release Mongo Session\n", c.SessionID)
			ses.Close()
		}()

		return h(c)
	}

	return f
}
