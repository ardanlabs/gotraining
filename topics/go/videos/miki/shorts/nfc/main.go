package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	city1 := "Kraków"
	city2 := "Kraków"

	fmt.Println(city1 == city2) // false

	city3 := norm.NFKC.String(city2)
	fmt.Println(city1 == city3) // true
}
