package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/sys/web"
)

// ErrorHandler for catching and responding errors.
func ErrorHandler(next web.Handler) web.Handler {

	// Create the handler that will be attached in the middleware chain.
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) (err error) {
		v := ctx.Value(web.KeyValues).(*web.Values)

		// In the event of a panic, we want to capture it here so we can send an
		// error down the stack.
		defer func() {
			if r := recover(); r != nil {

				// Respond with the error.
				web.RespondError(ctx, w, v.TraceID, errors.New("unhandled"), http.StatusInternalServerError)

				// Log out that we caught the error.
				log.Printf("%s : ERROR Midware : *****> Panic Caught : %s\n", v.TraceID, r)

				// Print out the stack.
				log.Printf("%s : ERROR Midware : *****> Stacktrace\n%s\n", v.TraceID, debug.Stack())

				// Capture the error for logging.
				err = fmt.Errorf("%v", r)
			}
		}()

		if err = next(ctx, w, r, params); err != nil {

			// Log out that we caught the error.
			log.Printf("%s : ERROR Midware : *****> %s\n", v.TraceID, err)

			// Respond with the error.
			web.Error(ctx, w, v.TraceID, err)
			return err
		}

		return nil
	}

	return h
}
