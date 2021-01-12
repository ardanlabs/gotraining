// Package main simulate the "Birthday problem".
// See https://en.wikipedia.org/wiki/Birthday_problem for a description of the problem.
package main

import (
	"fmt"
	"math/rand"
)

// simulateBirthdayMatches returns true if the same number is selected
// twice by the random number generator selecting a number between
// 0 and 365 for a specified group of people.
func simulateBirthdayMatches(numOfPeople int) bool {
	const daysInYear = 365

	seen := make(map[int]bool)
	for i := 0; i < numOfPeople; i++ {
		day := rand.Intn(daysInYear)
		if seen[day] {
			return true
		}
		seen[day] = true
	}

	return false
}

// simulateBirthdays returns the fraction of groups
// that have two people with the same birthday.
func simulateBirthdays(numOfPeople, runs int) float64 {
	same := 0
	for i := 0; i < runs; i++ {
		if simulateBirthdayMatches(numOfPeople) {
			same++
		}
	}

	return float64(same) / float64(runs)
}

func main() {
	fmt.Println(simulateBirthdays(23, 1_000_00))
}
