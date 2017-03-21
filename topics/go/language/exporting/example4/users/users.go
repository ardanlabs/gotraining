// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package users provides support for user management.
package users

// User represents information about a user.
type User struct {
	Name string
	ID   int

	password string
}
