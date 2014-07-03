package main

import "fmt"

type animal struct {
	name    string
	breed   string
	hasTail bool
}

type dog struct {
	animal animal
	bark   int
}

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
	Inc(&bart.bark)

	// Pass the "address of" the bart value.
	Display(&bart)
}

// Declaring bark as a pointer variable whose value is
// always an address and points to values of type int.
func Inc(bark *int) {
	*bark++
	fmt.Printf("&bark[%p] bark[%p] *bark[%d]\n", &bark, bark, *bark)
}

// Declaring d as a pointer variable whose value is always an address
// and points to values of type dog.
func Display(d *dog) {
	fmt.Printf("%+v\n", d)
}
