package series_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/kniren/gota/series"
)

func generateInts(n int) (data []int) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Int())
	}
	return
}

func generateFloats(n int) (data []float64) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Float64())
	}
	return
}

func generateStrings(n int) (data []string) {
	for i := 0; i < n; i++ {
		data = append(data, strconv.Itoa(rand.Int()))
	}
	return
}

func generateBools(n int) (data []bool) {
	for i := 0; i < n; i++ {
		r := rand.Intn(2)
		b := false
		if r == 1 {
			b = true
		}
		data = append(data, b)
	}
	return
}

func generateIntsN(n, k int) (data []int) {
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(k))
	}
	return
}

func BenchmarkSeries_New(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name       string
		data       interface{}
		seriesType series.Type
	}{
		{
			"[]bool(100000)_Int",
			generateBools(100000),
			series.Int,
		},
		{
			"[]bool(100000)_String",
			generateBools(100000),
			series.String,
		},
		{
			"[]bool(100000)_Bool",
			generateBools(100000),
			series.Bool,
		},
		{
			"[]bool(100000)_Float",
			generateBools(100000),
			series.Float,
		},
		{
			"[]string(100000)_Int",
			generateStrings(100000),
			series.Int,
		},
		{
			"[]string(100000)_String",
			generateStrings(100000),
			series.String,
		},
		{
			"[]string(100000)_Bool",
			generateStrings(100000),
			series.Bool,
		},
		{
			"[]string(100000)_Float",
			generateStrings(100000),
			series.Float,
		},
		{
			"[]float64(100000)_Int",
			generateFloats(100000),
			series.Int,
		},
		{
			"[]float64(100000)_String",
			generateFloats(100000),
			series.String,
		},
		{
			"[]float64(100000)_Bool",
			generateFloats(100000),
			series.Bool,
		},
		{
			"[]float64(100000)_Float",
			generateFloats(100000),
			series.Float,
		},
		{
			"[]int(100000)_Int",
			generateInts(100000),
			series.Int,
		},
		{
			"[]int(100000)_String",
			generateInts(100000),
			series.String,
		},
		{
			"[]int(100000)_Bool",
			generateInts(100000),
			series.Bool,
		},
		{
			"[]int(100000)_Float",
			generateInts(100000),
			series.Float,
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				series.New(test.data, test.seriesType, test.name)
			}
		})
	}
}

func BenchmarkSeries_Copy(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name   string
		series series.Series
	}{
		{
			"[]int(100000)_Int",
			series.Ints(generateInts(100000)),
		},
		{
			"[]int(100000)_String",
			series.Strings(generateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			series.Bools(generateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			series.Floats(generateInts(100000)),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.series.Copy()
			}
		})
	}
}

func BenchmarkSeries_Subset(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name    string
		indexes interface{}
		series  series.Series
	}{
		{
			"[]int(100000)_Int",
			generateIntsN(10000, 2),
			series.Ints(generateInts(100000)),
		},
		{
			"[]int(100000)_String",
			generateIntsN(10000, 2),
			series.Strings(generateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			generateIntsN(10000, 2),
			series.Bools(generateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			generateIntsN(10000, 2),
			series.Floats(generateInts(100000)),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.series.Subset(test.indexes)
			}
		})
	}
}

func BenchmarkSeries_Set(b *testing.B) {
	rand.Seed(100)
	table := []struct {
		name      string
		indexes   interface{}
		newValues series.Series
		series    series.Series
	}{
		{
			"[]int(100000)_Int",
			generateIntsN(10000, 2),
			series.Ints(generateIntsN(10000, 2)),
			series.Ints(generateInts(100000)),
		},
		{
			"[]int(100000)_String",
			generateIntsN(10000, 2),
			series.Strings(generateIntsN(10000, 2)),
			series.Strings(generateInts(100000)),
		},
		{
			"[]int(100000)_Bool",
			generateIntsN(10000, 2),
			series.Bools(generateIntsN(10000, 2)),
			series.Bools(generateInts(100000)),
		},
		{
			"[]int(100000)_Float",
			generateIntsN(10000, 2),
			series.Floats(generateIntsN(10000, 2)),
			series.Floats(generateInts(100000)),
		},
	}
	for _, test := range table {
		s := test.series.Copy()
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.Set(test.indexes, test.newValues)
			}
		})
	}
}
