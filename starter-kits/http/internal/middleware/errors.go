package middleware

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/web"
	"github.com/pkg/errors"
)

// ErrorHandler for catching and responding errors.
func ErrorHandler(next web.Handler) web.Handler {

	// Create the handler that will be attached in the middleware chain.
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(web.KeyValues).(*web.Values)

		// In the event of a panic, we want to capture it here so we can send an
		// error down the stack.
		defer func() {
			if r := recover(); r != nil {

				// Log the panic.
				log.Printf("%s : ERROR : Panic Caught : %s\n", v.TraceID, r)

				// Respond with the error.
				web.RespondError(ctx, w, errors.New("unhandled"), http.StatusInternalServerError)

				// Print out the stack.
				log.Printf("%s : ERROR : Stacktrace\n%s\n", v.TraceID, debug.Stack())
			}
		}()

		if err := next(ctx, w, r, params); err != nil {

			if errors.Cause(err) != web.ErrNotFound {

				// Log the error.
				log.Printf("%s : ERROR : %+v\n", v.TraceID, err)
			}

			// Respond with the error.
			web.Error(ctx, w, errors.Cause(err))

			// The error has been handled so we can stop propigating it.
			return nil
		}

		return nil
	}

	return h
}
