// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// animal is a type that describes any animal.
type animal struct {
	name    string
	breed   string
	hasTail bool
}

// dog is a type that describes a dog.
type dog struct {
	animal animal
	bark   int
}

// main is the entry point for the application.
func main() {
	// Declare and initialize a variable named bart of type dog.
	bart := dog{
		animal: animal{
			name:    "Bart",
			breed:   "Lab",
			hasTail: true,
		},
		bark: 10,
	}

	// Display the "value of" bart and the "address of" each field member.
	fmt.Printf("%+v &Name[%p] &Breed[%p] &HasTail[%p] &Bark[%p]\n",
		&bart, &bart.animal.name, &bart.animal.breed, &bart.animal.hasTail, &bart.bark)

	// Pass the "address of" the Bark field from within the bart value.
	inc(&bart.bark)

	// Pass the "address of" the bart value.
	display(&bart)
}

// inc declares bark as a pointer variable whose value is
// always an address and points to values of type int.
func inc(bark *int) {
	*bark++
	fmt.Printf("&bark[%p] bark[%p] *bark[%d]\n", &bark, bark, *bark)
}

// display declares d as a pointer variable whose value is always an address
// and points to values of type dog.
func display(d *dog) {
	fmt.Printf("%+v\n", d)
}
