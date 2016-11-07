// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package app provides a thin layer of support for writing web services. It
// integrates with the ardanlabs kit repo to provide support for logging,
// configuration, database, routing and application context. The base things
// you need to write a web service is provided.
package app

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/braintree/manners"
	"github.com/dimfeld/httptreemux"
	"github.com/pborman/uuid"
)

// TraceIDHeader is the header added to outgoing requests which adds the
// traceID to it.
const TraceIDHeader = "X-Trace-ID"

//==============================================================================

var (
	// ErrNotAuthorized occurs when the call is not authorized.
	ErrNotAuthorized = errors.New("Not authorized")

	// ErrDBNotConfigured occurs when the DB is not initialized.
	ErrDBNotConfigured = errors.New("DB not initialized")

	// ErrNotFound is abstracting the mgo not found error.
	ErrNotFound = errors.New("Entity not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in it's proper form")

	// ErrValidation occurs when there are validation errors.
	ErrValidation = errors.New("Validation errors occurred")
)

// app maintains some framework state.
var app = struct {
	userHeaders map[string]string
}{
	userHeaders: make(map[string]string),
}

//==============================================================================

// A Handler is a type that handles an http request within our own little mini
// framework. The fun part is that our context is fully controlled and
// configured by us so we can extend the functionality of the Context whenever
// we want.
type Handler func(*Context) error

// A Middleware is a type that wraps a handler to remove boilerplate or other
// concerns not direct to any given Handler.
type Middleware func(Handler) Handler

//==============================================================================

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
	Ctx map[string]interface{}

	mw []Middleware
}

// New create an App value that handle a set of routes for the application.
// You can provide any number of middleware and they'll be used to wrap every
// request handler.
func New(mw ...Middleware) *App {
	return &App{
		TreeMux: httptreemux.New(),
		Ctx:     make(map[string]interface{}),
		mw:      mw,
	}
}

// Group creates a new App Group based on the current App and provided
// middleware.
func (a *App) Group(mw ...Middleware) *Group {
	return &Group{
		app: a,
		mw:  mw,
	}
}

// Use adds the set of provided middleware onto the Application middleware
// chain. Any route running off of this App will use all the middleware provided
// this way always regardless of the ordering of the Handle/Use functions.
func (a *App) Use(mw ...Middleware) {
	a.mw = append(a.mw, mw...)
}

// Handle is our mechanism for mounting Handlers for a given HTTP verb and path
// pair, this makes for really easy, convenient routing.
func (a *App) Handle(verb, path string, handler Handler, mw ...Middleware) {

	// Wrap up the application-wide first, this will call the first function
	// of each middleware which will return a function of type Handler. Each
	// Handler will then be wrapped up with the other handlers from the chain.
	handler = wrapMiddleware(wrapMiddleware(handler, mw), a.mw)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		c := Context{
			ResponseWriter: w,
			Request:        r,
			Now:            time.Now(),
			Params:         p,
			SessionID:      uuid.New(),
			Ctx:            make(map[string]interface{}),
			App:            a,
		}

		// Set the request id on the outgoing requests before any other header to
		// ensure that the trace id is ALWAYS added to the request regardless of
		// any error occuring or not.
		c.Header().Set(TraceIDHeader, c.SessionID)

		// Call the wrapped handler and handle any possible error.
		if err := handler(&c); err != nil {
			c.Error(err)
		}
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
}

// CORS providing support for Cross-Origin Resource Sharing.
// https://metajack.im/2010/01/19/crossdomain-ajax-for-xmpp-http-binding-made-easy/
func (a *App) CORS() {
	h := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
	}

	a.TreeMux.OptionsHandler = h

	app.userHeaders["Access-Control-Allow-Origin"] = "*"
}

//==============================================================================

// Group allows a segment of middleware to be shared amongst handlers.
type Group struct {
	app *App
	mw  []Middleware
}

// Use adds the set of provided middleware onto the Application middleware chain.
func (g *Group) Use(mw ...Middleware) {
	g.mw = append(g.mw, mw...)
}

// Handle proxies the Handle function of the underlying App.
func (g *Group) Handle(verb, path string, handler Handler, mw ...Middleware) {

	// Wrap up the route specific middleware last because rememeber, the
	// middleware is wrapped backwards.
	handler = wrapMiddleware(handler, mw)

	// Wrap it with the App wrapper and additionally the group level middleware.
	g.app.Handle(verb, path, handler, g.mw...)
}

//==============================================================================

// Run is called to start the web service.
func Run(host string, routes http.Handler, readTimeout, writeTimeout time.Duration) error {

	log.Printf("startup : Run : Start : Using Host[%s]\n", host)

	// Create a new server and set timeout values.
	server := manners.NewWithServer(&http.Server{
		Addr:           host,
		Handler:        routes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	})

	go func() {

		// Listen for an interrupt signal from the OS.
		osSignals := make(chan os.Signal)
		signal.Notify(osSignals, os.Interrupt)

		sig := <-osSignals
		log.Printf("shutdown : Run : Captured %v. Shutting Down...\n", sig)

		// Shut down the API server.
		server.Close()
	}()

	return server.ListenAndServe()
}

//==============================================================================

// wrapMiddleware wraps a handler with some middleware.
func wrapMiddleware(handler Handler, mw []Middleware) Handler {

	// Wrap with our group specific middleware.
	for i := len(mw) - 1; i >= 0; i-- {
		if mw[i] != nil {
			handler = mw[i](handler)
		}
	}

	return handler
}
