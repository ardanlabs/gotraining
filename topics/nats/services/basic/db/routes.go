// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"sync"

	"github.com/nats-io/nats"
)

func init() {
	AddRoute("/db/users", GetUsers)
}

//==============================================================================

// Handler declares the type for the message handler function.
type Handler func(*nats.EncodedConn, string, Request)

// routes maintains a set of uri to handler function routes.
var routes = struct {
	mu sync.Mutex
	h  map[string]Handler
}{
	h: make(map[string]Handler),
}

// =============================================================================

// AddRoute adds a new routes to the list of routes.
func AddRoute(uri string, h Handler) {
	routes.mu.Lock()
	{
		routes.h[uri] = h
	}
	routes.mu.Unlock()
}

// Route looks up the specified uri for the message handler function.
func Route(uri string) (Handler, bool) {
	var h Handler
	var exists bool

	routes.mu.Lock()
	{
		h, exists = routes.h[uri]
	}
	routes.mu.Unlock()

	return h, exists
}
