// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/GF0gbY4SvN

// Sample program to show the properties of nil maps.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	show(map[int]string{1: "one", 2: "two", 3: "three"})

	fmt.Println()

	show(nil)
}

func show(m map[int]string) {
	defer func() {
		v := recover()
		if v != nil {
			fmt.Println("FAILED:", v)
		}
	}()

	fmt.Printf("Checking: %#v\n", m)

	fmt.Print("Iteration ... ")
	for range m {
	}
	fmt.Println("OK")

	fmt.Print("Key Read ... ")
	_ = m[23]
	fmt.Println("OK")

	fmt.Print("Deletion ... ")
	delete(m, 23)
	fmt.Println("OK")

	fmt.Print("Key Assignment ... ")
	m[4] = "four"
	fmt.Println("OK")
}
