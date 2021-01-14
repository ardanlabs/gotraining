// Package main simulate two-sice rolls.
package main

import (
	"fmt"
	"math/rand"
)

// diceRoll simulate a dice roll.
func diceRoll() int {
	// Intn(6) returns values in the range 0-5 (inclusive), we want 1-6.
	return rand.Intn(6) + 1
}

// simulate runs n simulation of two game cube rolls.
// It returns the percentage for each total of first and second roll.
func simulate(runs int) map[int]float64 {
	counts := make(map[int]int)
	for i := 0; i < runs; i++ {
		val := diceRoll() + diceRoll()
		counts[val]++
	}

	// Convert from counts to fractions.
	fracs := make(map[int]float64)
	for val, count := range counts {
		frac := float64(count) / float64(runs)
		fracs[val] = frac
	}

	return fracs
}

func main() {
	fracs := simulate(1_000_000)
	for i := 2; i <= 12; i++ {
		fmt.Printf("%2d -> %.2f\n", i, fracs[i])
	}
}
