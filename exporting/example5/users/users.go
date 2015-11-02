// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package users provides support for user management.
package users

// user represents information about a user.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
type Manager struct {
	Title string

	user
}
