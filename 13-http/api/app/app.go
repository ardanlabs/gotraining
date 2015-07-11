package app

import (
	"errors"
	"net/http"

	"code.google.com/p/go-uuid/uuid"

	"github.com/dimfeld/httptreemux"
)

var (
	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("No user(s) found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

// A Handler is a type that handles an http request within our own little mini
// framework. The fun part is that our context is fully controlled and
// configured by us so we can extend the functionality of the Context whenever
// we want.
type Handler func(*Context) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
}

// New create an App value that handle a set of routes for the application.
func New() *App {
	return &App{
		TreeMux: httptreemux.New(),
	}
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler) {
	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		c := Context{
			Session:        GetSession(),
			ResponseWriter: w,
			Request:        r,
			Params:         p,
			SessionID:      uuid.New(),
		}
		defer c.Session.Close()

		// Call the handler and handle any possible error.
		if err := handler(&c); err != nil {
			c.Error(err)
		}
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}
