package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example7/toy"
)

// main is the entry point for the application.
func main() {
	bat := toy.New()
	bat.Height = 28
	bat.Weight = 16

	fmt.Println(bat)
}
