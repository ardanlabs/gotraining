package game

import "fmt"

func ExampleKey() {
	fmt.Println(Copper)
	fmt.Println(Key(9))

	fmt.Printf(" v: %v\n", Jade)
	fmt.Printf("+v: %+v\n", Jade)
	fmt.Printf("#v: %#v\n", Jade)

	p1 := Player{
		Name: "Parzival",
		Keys: []Key{Jade},
	}
	fmt.Printf("p1: %#v\n", p1)

	// Output:
	// copper
	// <Key 9>
	//  v: jade
	// +v: jade
	// #v: 0x2
	// p1: game.Player{Name:"Parzival", Keys:[]game.Key{0x2}}
}
