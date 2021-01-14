// Calculate the value of π using simulation.
// See https://medium.com/cantors-paradise/estimating-%CF%80-using-monte-carlo-simulations-3459a84b5ef9
package main

import (
	"fmt"
	"math"
	"math/rand"
)

// calculatePi calculates the value of π using n points.
func calculatePi(n int) float64 {
	const radius = 1
	var innerPoints int

	for i := 0; i < n; i++ {
		x, y := rand.Float64(), rand.Float64()
		if math.Sqrt(x*x+y*y) < radius {
			innerPoints++
		}
	}

	ratio := float64(innerPoints) / float64(n)
	// Since radius = 1, then circle area is π. We calculated points in range
	// (0, 0) <-> (1, 1) which is 1/4 of the circle.
	return 4 * ratio
}

func main() {
	fmt.Println("π =", calculatePi(100_000_000))
}
