// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/cZTG1NdSqC

// Sample program to show how to use an interface in Go.
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
)

// notify implements a method that can be called via
// a value of type user.
func (u user) notify() {
	fmt.Printf("user: Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create two values of type user.
	user1 := user{"Bill", "bill@email.com"}
	user2 := &user{"Jill", "jill@email.com"}

	// Values and pointers of type user implement the interface because value
	// receivers belong to the method sets of both values and pointers.

	// Pass a pointer of the values to support the interface.
	sendNotification(&user1)
	sendNotification(user2)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
