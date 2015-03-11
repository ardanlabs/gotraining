package app

import (
	"net/http"

	"code.google.com/p/go-uuid/uuid"

	"github.com/dimfeld/httptreemux"
)

// A Handler is a type that handles an http request within our own little mini
// framework. The fun part is that our context is fully controlled and
// configured by us so we can extend the functionality of the Context whenever
// we want.
type Handler func(*Context)

// This is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
}

func New() *App {
	return &App{
		TreeMux: httptreemux.New(),
	}
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler) {
	a.TreeMux.Handle(verb, path, func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		// create our context
		c := Context{
			Session:        GetSession(),
			ResponseWriter: w,
			Request:        r,
			Params:         p,
			SessionID:      uuid.New(),
		}
		defer c.Session.Close()

		// todo: defer here

		// authenticate our user
		if c.Authenticate() != nil {
			return
		}

		// call the handler
		handler(&c)
	})
}
