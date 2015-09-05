// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/cjZ-iIHIsK

// Sample program to show how polymorphic behavior with interfaces.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// admin defines a admin in the program.
type admin struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending admin Email To %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the application.
func main() {
	// Create two values one of type user and one of type admin.
	bill := user{"Bill", "bill@email.com"}
	jill := admin{"Jill", "jill@email.com"}

	// Pass a pointer of the values to support the interface.
	sendNotification(&bill)
	sendNotification(&jill)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
