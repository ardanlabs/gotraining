package stats

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](values []T) T {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m
}
