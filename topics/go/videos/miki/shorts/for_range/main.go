package main

import (
	"fmt"
)

const (
	maxX = 1000
	maxY = 600
)

type Point struct {
	X int
	Y int
}

func inBounds(p Point) bool {
	return p.X >= 0 && p.X <= maxX && p.Y >= 0 && p.Y <= maxY
}

func main() {
	ch := make(chan bool)
	points := []Point{
		{10, 20},
		{900, 700},
		{450, 353},
	}

	// fan out
	for _, p := range points {
		p := p // avoid closure capture
		go func() {
			ch <- inBounds(p)
		}()
	}

	// collect
	numOK := 0
	for range points {
		ok := <-ch
		if ok {
			numOK++
		}
	}
	fmt.Println(numOK, "points in range")
}
