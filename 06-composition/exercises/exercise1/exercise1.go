// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/UvFEZQHDu0

// Follow the guided comments to:
//
// Declare a sysadmin type that implements the administrator interface.
//
// Declare a programmer type that implements the developer interface.
//
// Declare a company type that embeds both an administrator and a developer.
//
// Create a sysadmin, programmers, and a company which are available for hire,
// and use them to complete some predefined tasks.
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

// Declare a struct type named sysadmin: it should have a name field.
type sysadmin struct {
	name string
}

// Define an administrate method on the sysadmin type, implementing the
// administrator interface.  administrate should print out the name of the
// sysadmin, as well as the system they are administering.
func (s *sysadmin) administrate(system string) {
	fmt.Println(s.name, "is administering", system)
}

// Declare a struct type named programmer: it should have a name field.
type programmer struct {
	name string
}

// Define a develop method on the programmer type, implementing the developer
// interface.  develop should print out the name of the programmer, as well as
// the system they are developing.
func (p *programmer) develop(system string) {
	fmt.Println(p.name, "is developing", system)
}

// Declare a struct type named company: it should embed administrator and developer.
type company struct {
	administrator
	developer
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

// tasks contains a set of systems we must administer or develop.
var tasks = []struct {
	system     string
	needsDev   bool
	needsAdmin bool
}{
	{system: "exercise1", needsDev: true},
	{system: "server5", needsAdmin: true},
	{system: "project-omega", needsDev: true},
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

	// Create a variable named techfirm of type company, and initialize it by
	// hiring (popping) an administrator from admins and a developer from devs.
	techfirm := company{
		administrator: admins.popAdmin(),
		developer:     devs.popDev(),
	}

	// Push techfirm onto both devs and admins (we can now transparently
	// outsource to techfirm for development and administrative needs).
	admins.pushAdmin(techfirm)
	devs.pushDev(techfirm)

	// Iterate over tasks.
	for _, task := range tasks {
		// Check if the task needs a developer. If so, pop a developer from devs,
		// print its type information, and have it develop the system.
		if task.needsDev {
			dev := devs.popDev()
			fmt.Printf("Developer Type: %T\n", dev)
			dev.develop(task.system)
		}

		// Check if the task needs an administrator. If so, pop an administrator from
		// admins, print its type information, and have it administrate the system.
		if task.needsAdmin {
			admin := admins.popAdmin()
			fmt.Printf("Administrator Type: %T\n", admin)
			admin.administrate(task.system)
		}
	}
}
