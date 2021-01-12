// Package main simulate precision of a medical test.
// For more details see:
// https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
package main

import (
	"fmt"
	"math/rand"
)

// oneChanceIn returns true one in n times.
func oneChanceIn(n int) bool {
	return rand.Intn(n) == 1
}

// isSick returns true if a randomly sampled person is sick.
func isSick() bool {
	// The disease strikes 1/1000 of the population.
	return oneChanceIn(1000)
}

// diagnosed returns true if a person is sick or misdiagnosed as sick.
func diagnosed(sick bool) bool {
	if sick {
		return true // We're 100% correct in sick people.
	}

	// The test of a disease presents a rate of 5% (1 in 20) false positives.
	// (false positive = healthy diagnosed as sick)
	return oneChanceIn(20)
}

// simulate run selects sampleSize random people and return the fraction of people
// actually sick from the total number of people diagnosed as sick.
func simulate(sampleSize int) float64 {
	var numSick, numDiagnosed int

	for i := 0; i < sampleSize; i++ {
		sick := isSick()
		if sick {
			numSick++
		}

		if diagnosed(sick) {
			numDiagnosed++
		}
	}

	return float64(numSick) / float64(numDiagnosed)
}

func main() {
	fmt.Println(simulate(1_000_000))
}
