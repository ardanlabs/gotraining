// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Qn32CmIAIn

// Sample program to show what happens when the outer and inner
// type implement the same interface.
package main

import (
	"fmt"
)

type (
	// notifier is an interface that defined notification
	// type behavior.
	notifier interface {
		notify()
	}

	// user defines a user in the program.
	user struct {
		name  string
		email string
	}

	// admin represents an admin user with privileges.
	admin struct {
		user
		level string
	}
)

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("user: Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// notify implements a method that can be called via
// a value of type Admin.
func (a *admin) notify() {
	fmt.Printf("User: Sending Admin Email To %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the application.
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is NOT "promoted" to the outer type.
	sendNotification(&ad)

	// We can acces the inner type's method direectly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
