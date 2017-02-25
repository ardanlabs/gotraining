// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example that creates interface pollution
// by improperly using an interface when one is not needed.
package main

// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation.
type server struct {
	host string

	// PRETEND THERE ARE MORE FIELDS.
}

// NewServer returns an interface value of type Server
// with a server implementation.
func NewServer(host string) Server {

	// SMELL - Storing an unexported type pointer in the interface.
	return &server{host}
}

// Start allows the server to begin to accept requests.
func (s *server) Start() error {

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Stop shuts the server down.
func (s *server) Stop() error {

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {

	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func main() {

	// Create a new Server.
	srv := NewServer("localhost")

	// Use the API.
	srv.Start()
	srv.Stop()
	srv.Wait()
}

// =============================================================================

// NOTES:

// Smells:
//  * The package declares an interface that matches the entire API of its own concrete type.
//  * The interface is exported but the concrete type is unexported.
//  * The factory function returns the interface value with the unexported concrete type value inside.
//  * The interface can be removed and nothing changes for the user of the API.
//  * The interface is not decoupling the API from change.
