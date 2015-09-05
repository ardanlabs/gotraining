// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package animals provides support for animals.
package animals

// animal represents information about all animals.
type animal struct {
	Name string
	Age  int
}

// Dog represents information about dogs.
type Dog struct {
	animal       // The embedded type is unexported.
	BarkStrength int
}
