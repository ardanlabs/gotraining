// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/0NS3SbImQ3

// Sample program to show how you can personally mock concrete types when
// you need to for your own packages or tests.
package main

import (
	"github.com/ardanlabs/gotraining/topics/composition/example7/pubsub"
)

// publisher is an interface to allow this package to mock the Publish
// behavior from the pubsub package.
type publisher interface {
	Publish(key string, v interface{}) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// main is the entry point for the application.
func main() {
	var p publisher

	// Use the pubsub package.
	p = pubsub.New("localhost")
	p.Publish("key", "value")

	// Use the mock type value.
	p = &mock{}
	p.Publish("key", "value")
}
