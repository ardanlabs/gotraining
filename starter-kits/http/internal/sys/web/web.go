// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package web provides a thin layer of support for writing web services.
package web

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/braintree/manners"
	"github.com/dimfeld/httptreemux"
	"github.com/pborman/uuid"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/mgo.v2"
)

// TraceIDHeader is the header added to outgoing requests which adds the
// traceID to it.
const TraceIDHeader = "X-Trace-ID"

//==============================================================================

// validate provides a validator for checking models.
var validate = validator.New(&validator.Config{
	TagName:      "validate",
	FieldNameTag: "json",
})

// Unmarshal decodes the input to the struct type and checks the
// fields to verify the value is in a proper state.
func Unmarshal(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return err
	}

	var inv InvalidError
	if fve := validate.Struct(v); fve != nil {
		for _, fe := range fve.(validator.ValidationErrors) {
			inv = append(inv, Invalid{Fld: fe.Field, Err: fe.Tag})
		}
		return inv
	}

	return nil
}

//==============================================================================

// Key represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	DB         *mgo.Session
	TraceID    string
	Now        time.Time
	StatusCode int
}

//==============================================================================

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// A Middleware is a type that wraps a handler to remove boilerplate or other
// concerns not direct to any given Handler.
type Middleware func(Handler) Handler

//==============================================================================

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	*httptreemux.TreeMux
	Values map[string]interface{}

	mw []Middleware
}

// New create an App value that handle a set of routes for the application.
// You can provide any number of middleware and they'll be used to wrap every
// request handler.
func New(mw ...Middleware) *App {
	return &App{
		TreeMux: httptreemux.New(),
		Values:  make(map[string]interface{}),
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
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		// Create the context for the request.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: uuid.New(),
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValues, &v)

		// Set the trace id on the outgoing requests before any other header to
		// ensure that the trace id is ALWAYS added to the request regardless of
		// any error occuring or not.
		w.Header().Set(TraceIDHeader, v.TraceID)

		// Call the wrapped handler functions.
		handler(ctx, w, r, params)
	}

	// Add this handler for the specified verb and route.
	a.TreeMux.Handle(verb, path, h)
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
	log.Printf("Run : Start : Using Host[%s]\n", host)

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
		log.Printf("Run : Captured %v. Shutting Down...\n", sig)

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
