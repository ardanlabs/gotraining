// Sample program to show how to iterate over a slice.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dog represents an animal of type dog.
type Dog struct {
	Name string
	Age  int
}

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for the application.
func main() {
	// Get the slice of Dog values.
	dogs := getDogs()

	// Iterate through the slice and displays the values.
	for _, dog := range dogs {
		fmt.Printf("Addr[%p] Name[%s]\tAge[%d]\n", &dog, dog.Name, dog.Age)
	}
}

// getDogs returns a slice of dog values.
func getDogs() []Dog {
	// Create a slice of names.
	names := []string{"bill", "jack", "sammy", "jill", "choley", "harley", "jamie", "Ed", "Lisa", "Missy"}

	// Create a nil slice of Dogs and randomly append Dog values
	// of different names and ages.
	var dogs []Dog
	for dog := 0; dog < 25; dog++ {
		name := rand.Intn(9)
		age := rand.Intn(20)
		dogs = append(dogs, Dog{names[name], age})
	}

	// Return the slice of Dogs.
	return dogs
}
