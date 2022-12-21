package stats

import "fmt"

// MaxInts return the maximal value in a slice of int
func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

// MaxFloat64s return the maximal value in a slice of float64
func MaxFloat64s(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}
