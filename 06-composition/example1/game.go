package main

import "github.com/ArdanStudios/gotraining/06-composition/example1/objects"

func main() {
	b := objects.NewBuilding(10, 0, "rect", "silver")
	c := objects.NewCloud(10, 1000)
	p := objects.NewPlayer(20, 0, "skinny", "blue")

	b.Display()
	c.Display()
	p.Display()

	c.Move(20, 1100)
	p.Move(45, 10)

	b.Display()
	c.Display()
	p.Display()
}
