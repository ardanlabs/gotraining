// Package utilities implements a host of helpful miscellaneous functions to the library.
package utilities

import (
	"sort"

	"gonum.org/v1/gonum/mat"
)

type sortedIntMap struct {
	m map[int]float64
	s []int
}

func (sm *sortedIntMap) Len() int {
	return len(sm.m)
}

func (sm *sortedIntMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] < sm.m[sm.s[j]]
}

func (sm *sortedIntMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func SortIntMap(m map[int]float64) []int {
	sm := new(sortedIntMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func FloatsToMatrix(floats []float64) *mat.Dense {
	return mat.NewDense(1, len(floats), floats)
}

func VectorToMatrix(vector mat.Vector) *mat.Dense {
	denseCopy := mat.VecDenseCopyOf(vector)
	vec := denseCopy.RawVector()
	return mat.NewDense(1, len(vec.Data), vec.Data)
}
