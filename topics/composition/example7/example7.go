// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/Z-Wdj9Mfmg

// Sample program to show how you can personally mock concrete types when
// you need to for your own packages or tests.
package main

import (
	"github.com/ardanlabs/gotraining/topics/interfaces/example4/pubsub"
)

// publish is an interface to allow this package to mock the Publish
// behavior from the pubsub package.
type publish interface {
	Publish(key string, v interface{}) error
}

// mock is a concrete type to help support the mock.
type mock struct{}

// Publish implements the publish interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// main is the entry point for the application.
func main() {
	var pub publish

	// Use the pubsub package.
	pub = pubsub.New("localhost")
	pub.Publish("key", "value")

	// Use the mock type value.
	pub = &mock{}
	pub.Publish("key", "value")
}
