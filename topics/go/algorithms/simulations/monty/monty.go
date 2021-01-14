// Simulation of "Monty Hall problem"
// https://en.wikipedia.org/wiki/Monty_Hall_problem
package main

import (
	"fmt"
	"math/rand"
)

// stayWinsGame is a single simulation of game. It return true if "stay" strategy wins.
func stayWinsGame() bool {
	carDoor := rand.Intn(3)
	playerDoor := rand.Intn(3)

	return carDoor == playerDoor
}

// fraction return a/b as float
func fraction(a, b int) float64 {
	return float64(a) / float64(b)
}

// simulation runs n games and return the fraction of games where "stay"
// strategy won and fraction where "switch" strategy won
func simulation(n int) (float64, float64) {
	stayWin, switchWin := 0, 0
	for i := 0; i < n; i++ {
		if stayWinsGame() {
			stayWin++
		} else {
			switchWin++
		}
	}

	return fraction(stayWin, n), fraction(switchWin, n)
}

func main() {
	stayFrac, switchFrac := simulation(1_000_000)
	fmt.Printf("stay: %%%.2f\n", stayFrac*100)
	fmt.Printf("switch: %%%.2f\n", switchFrac*100)
}
