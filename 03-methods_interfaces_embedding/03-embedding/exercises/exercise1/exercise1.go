// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/aIES0zfHfg

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

import "fmt"

// nail represents a nail that is or is not fastened.
type nail struct{}

// nailDriver is implemented by tools to fasten nails.
type nailDriver interface {
	driveNail(nails []nail) []nail
}

// nailPuller is implemented by tools to unfasten nails.
type nailPuller interface {
	pullNail(nails []nail, totalNails int) []nail
}

// clawhammer is a tool that operates on nails.
type clawhammer struct{}

// driveNail allows the clawhammer to fasten a nail.
func (h clawhammer) driveNail(nails []nail) []nail {
	if len(nails) == 0 {
		return nails
	}

	fmt.Println("Clawhammer: Fastened nail into the board.")
	return nails[1:]
}

// pullNail allows the clawhammer to unfasten a nail.
func (h clawhammer) pullNail(nails []nail, totalNails int) []nail {
	if len(nails) == totalNails {
		return nails
	}

	fmt.Println("Clawhammer: Unfastened nail from the board.")
	return append(nails, nail{})
}

// toolbox can contains any number of tools.
type toolbox struct {
	totalNails int
	nails      []nail
	clawhammer
}

// addNails adds any number of nails to the toolbox.
func (tb *toolbox) addNails(n int) {
	for i := 0; i < n; i++ {
		tb.nails = append(tb.nails, nail{})
	}

	tb.totalNails = len(tb.nails)
}

// nailCount returns the number of total nails and left in the box.
func (tb *toolbox) nailCount() (total int, left int) {
	return tb.totalNails, tb.totalNails - len(tb.nails)
}

// main is the entry point for the application.
func main() {
	// Create a toolbox.
	var tb toolbox

	// Add 4 nails to the toolbox.
	tb.addNails(4)

	// Display the nails we have in the toolbox.
	total, left := tb.nailCount()
	fmt.Println("Total:", total, "Left:", left)

	// Fasten 2 of the nails using the clawhammer.
	tb.nails = tb.clawhammer.driveNail(tb.nails)
	tb.nails = tb.clawhammer.driveNail(tb.nails)

	// Display the nails we have in the toolbox.
	total, left = tb.nailCount()
	fmt.Println("Total:", total, "Left:", left)

	// Unfasten the 2 nails using the clawhammer.
	tb.nails = tb.clawhammer.pullNail(tb.nails, tb.totalNails)
	tb.nails = tb.clawhammer.pullNail(tb.nails, tb.totalNails)

	// Display the nails we have in the toolbox.
	total, left = tb.nailCount()
	fmt.Println("Total:", total, "Left:", left)
}
