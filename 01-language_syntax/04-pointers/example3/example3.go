// http://play.golang.org/p/45qPXTLif_

// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// dog is a type that describes any dog.
type dog struct {
	name    string
	hasTail bool
	bark    int
}

// main is the entry point for the application.
func main() {
	// Declare and initialize a variable named bart of type dog.
	bart := dog{
		name:    "Bart",
		hasTail: true,
		bark:    10,
	}

	// Pass the "address of" the bart value.
	display(&bart)

	// Pass the "address of" the Bark field from within the bart value.
	increment(&bart.bark)

	// Pass the "address of" the bart value.
	display(&bart)
}

// increment declares bark as a pointer variable whose value is
// always an address and points to values of type int.
func increment(bark *int) {
	*bark++
	fmt.Printf("&bark[%p] bark[%p] *bark[%d]\n", &bark, bark, *bark)
}

// display declares a as a pointer variable whose value is always an address
// and points to values of type dog.
func display(a *dog) {
	fmt.Printf("%p\t%+v\n", a, *a)
}
