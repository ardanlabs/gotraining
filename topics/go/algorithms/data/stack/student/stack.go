// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package stack asks the student to implement a stack in Go.
package stack

// Data represents what is being stored on the stack.
type Data struct {
	Name string
}

// Stack represents a stack of data.
type Stack struct {
	data []*Data
}

// Make allows the creation of a stack with an initial
// capacity for efficiency. Otherwise a stack can be
// used in its zero value state.
func Make(cap int) *Stack {
	return nil
}

// Count returns the number of items in the stack.
func (s *Stack) Count() int {
	return 0
}

// Push adds data into the stack.
func (s *Stack) Push(data *Data) {
}

// Pop removes data from the stack.
func (s *Stack) Pop() (*Data, error) {
	return nil, nil
}

// Peek provides the data stored on the stack based
// on the level from the bottom. A value of 0 would
// return the top piece of data.
func (s *Stack) Peek(level int) (*Data, error) {
	return nil, nil
}

// Operate accepts a function that takes data and calls
// the specified function for every piece of data found.
// It traverses from the top down through the stack.
func (s *Stack) Operate(f func(data *Data) error) error {
	return nil
}
