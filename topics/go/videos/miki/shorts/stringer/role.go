// Use golang.org/x/tools/cmd/stringer to generate fmt.Stringer for your types.

package main

import "fmt"

type Role byte

const (
	Reader Role = iota + 1
	Writer
	Admin
)

func main() {
	// Will print: "Writer" after running
	// "stringer -type Role role.go"
	fmt.Println(Writer)
}
