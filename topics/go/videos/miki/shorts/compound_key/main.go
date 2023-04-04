// You can use struct as a compound key in maps.

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	colors := make(map[Point]int)
	pt := Point{10, 20}
	colors[pt] = 0xFF0000
	fmt.Printf("%X\n", colors[pt])
}
