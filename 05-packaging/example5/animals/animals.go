// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package animals provides support for animals.
package animals

// Animal represents information about all animals.
type Animal struct {
	Name string
	Age  int
}

// Dog represents information about dogs.
type Dog struct {
	Animal
	BarkStrength int
}
