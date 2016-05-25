// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how you can personally mock concrete types when
// you need to for your own packages or tests.
package main

import (
	"github.com/ardanlabs/gotraining/topics/composition/example7/pubsub"
)

// publisher is an interface to allow this package to mock the pubsub
// package support.
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}

// Publish implements the publisher interface for the mock.
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {

	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}

func main() {
	var p publisher

	// Use the pubsub package.
	p = pubsub.New("localhost")
	p.Publish("key", "value")
	p.Subscribe("key")

	// Use the mock type value.
	p = &mock{}
	p.Publish("key", "value")
	p.Subscribe("key")
}
