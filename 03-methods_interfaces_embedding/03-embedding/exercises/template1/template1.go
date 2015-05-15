// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/XOBHUvz5uz

// Declare a nail type using an empty struct. Then declare two interfaces, one
// named nailDriver and the other named nailPuller. nailDriver has one method named
// driveNail that accepts a slice of nails and returns a slice of nails. nailPuller
// has one method named pullNail that also accepts and returns a slice of nails,
// but also accepts the totalNails that exist.
//
// Declare a tool type named clawhammer using an empty struct. Implement the
// two interfaces using a value receiver. For the driveNail method, check that
// the number of nails in the slice is not zero and then remove the first
// element of the slice. For the pullNail method, check the number of nails in
// the slice with the total nails and then append a new nail to the slice.
//
// Declare a toolbox type with two fields named totalNails and nails. The
// totalNails field will be of type int and the nails field will be a slice
// of nails. Then embed the clawhammer tool inside the toolbox type. Declare a method
// named addNails using a pointer reciever that accepts an integer for the number
// of nails to add to the toolbox. Append that number of nail values to the slice
// and set the totalNails field. Declare a method named nailCount using a pointer
// receiver that returns the total number of nails in the slice and the value for
// the totalNails field.
//
// Write a main function that creates a toolbox, adds nails to the toolbox,
// displays the nail count, uses the clawhammer to fasten some nails, displays
// the nail count again, unfastens the nails and displays the nail count one
// more time.
package main

// Declare a type named nail as an empty struct.

// Declare an interface type named nailDriver.
// Declare a method name driveNail that accepts a slice of nails and returns
// a slice of nails.

// Declare an interface type named nailPuller.
// Declare a method name pullNail that accepts a slice of nails and returns
// a slice of nails. It also accepts the totalNails that exist.

// Declare a type named clawhammer as an empty struct.

// Declare a method that implements the nailDriver interface using a value
// reciever.
func ( /* receiver */ ) driveNail(nails []nail) []nail {
	// Check that the number of nails in the slice is not
	// zero and then remove the first element of the slice.
}

// Declare a method that implements the nailPuller interface using a value
// reciever.
func ( /* receiver */ ) pullNail(nails []nail, totalNails int) []nail {
	// Check the number of nails in the slice with the total
	// nails and then append a new nail to the slice.
}

// Declare a type named toolbox with two fields named totalNails and nails. The
// totalNails field will be of type int and the nails field will be a slice
// of nails. Then embed the clawhammer tool inside the toolbox type.

// Declare a method named addNails using a value receiver that accepts an integer
// for the number of nails to add to the toolbox.
func ( /* receiver */ ) addNails(n int) {
	// Append that number of nail values to the slice and
	// set the totalNails field.
}

// Declare a method named nailCount using a pointer receiver that returns the
// total number of nails in the slice and the value for the totalNails field.
// nailCount returns the number of total nails and left in the box.
func ( /* receiver */ ) nailCount() (total int, left int) {
}

// main is the entry point for the application.
func main() {
	// Create a toolbox.

	// Add 4 nails to the toolbox.

	// Display the nails we have in the toolbox.

	// Fasten 2 of the nails using the clawhammer.

	// Display the nails we have in the toolbox.

	// Unfasten the 2 nails using the clawhammer.

	// Display the nails we have in the toolbox.
}
