// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Using the template, declare a set of concrete types that implement the set
// of predefined interface types. Then create values of these types and use
// them to complete a set of predefined tasks.
package main

// Add import(s).
import "fmt"

// administrator represents a person or other entity capable of administering
// hardware and software infrastructure.
type administrator interface {
	administrate(system string)
}

// developer represents a person or other entity capable of writing software.
type developer interface {
	develop(system string)
}

// =============================================================================

// adminlist represents a group of administrators.
type adminlist struct {
	list []administrator
}

// pushAdmin adds an administrator to the adminlist.
func (l *adminlist) pushAdmin(a administrator) {
	l.list = append(l.list, a)
}

// popAdmin removes an administrator from the adminlist.
func (l *adminlist) popAdmin() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// =============================================================================

// devlist represents a group of developers.
type devlist struct {
	list []developer
}

// pushDev adds a developer to the devlist.
func (l *devlist) pushDev(d developer) {
	l.list = append(l.list, d)
}

// popDev removes a developer from the devlist.
func (l *devlist) popDev() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// =============================================================================

// Declare a concrete type named sysadmin with a name field of type string.
type sysadmin struct {
	name string
}

// Declare a method named administrate for the sysadmin type, implementing the
// administrator interface. administrate should print out the name of the
// sysadmin, as well as the system they are administering.
func (s *sysadmin) administrate(system string) {
	fmt.Println(s.name, "is administering", system)
}

// Declare a concrete type named programmer with a name field of type string.
type programmer struct {
	name string
}

// Declare a method named develop for the programmer type, implementing the
// developer interface. develop should print out the name of the
// programmer, as well as the system they are coding.
func (p *programmer) develop(system string) {
	fmt.Println(p.name, "is developing", system)
}

// Declare a concrete type named company. Declare it as the composition of
// the administrator and developer interface types.
type company struct {
	administrator
	developer
}

// =============================================================================

func main() {

	// Create a variable named admins of type adminlist.
	var admins adminlist

	// Create a variable named devs of type devlist.
	var devs devlist

	// Push a new sysadmin onto admins.
	admins.pushAdmin(&sysadmin{"John"})

	// Push two new programmers onto devs.
	devs.pushDev(&programmer{"Mary"})
	devs.pushDev(&programmer{"Steve"})

	// Create a variable named cmp of type company, and initialize it by
	// hiring (popping) an administrator from admins and a developer from devs.
	cmp := company{
		administrator: admins.popAdmin(),
		developer:     devs.popDev(),
	}

	// Push the company value on both lists since the company implements
	// each interface.
	admins.pushAdmin(&cmp)
	devs.pushDev(&cmp)

	// A set of tasks for administrators and developers to perform.
	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	// Iterate over tasks.
	for _, task := range tasks {

		// Check if the task needs an administrator else use a developer.
		if task.needsAdmin {

			// Pop an administrator value from the admins list and
			// call the administrate method.
			adm := admins.popAdmin()
			adm.administrate(task.system)

			continue
		}

		// Pop a developer value from the devs list and
		// call the develop method.
		dev := devs.popDev()
		dev.develop(task.system)
	}
}
