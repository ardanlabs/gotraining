package dataframe_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

func generateSeries(n, rep int) (data []series.Series) {
	rand.Seed(100)
	for j := 0; j < rep; j++ {
		var is []int
		var bs []bool
		var fs []float64
		var ss []string
		for i := 0; i < n; i++ {
			is = append(is, rand.Int())
		}
		for i := 0; i < n; i++ {
			fs = append(fs, rand.Float64())
		}
		for i := 0; i < n; i++ {
			ss = append(ss, strconv.Itoa(rand.Int()))
		}
		for i := 0; i < n; i++ {
			r := rand.Intn(2)
			b := false
			if r == 1 {
				b = true
			}
			bs = append(bs, b)
		}
		data = append(data, series.Ints(is))
		data = append(data, series.Bools(bs))
		data = append(data, series.Floats(fs))
		data = append(data, series.Strings(ss))
	}
	return
}

func BenchmarkNew(b *testing.B) {
	table := []struct {
		name string
		data []series.Series
	}{
		{
			"100000x4",
			generateSeries(100000, 1),
		},
		{
			"100000x40",
			generateSeries(100000, 10),
		},
		{
			"100000x400",
			generateSeries(100000, 100),
		},
		{
			"1000x40",
			generateSeries(1000, 10),
		},
		{
			"1000x4000",
			generateSeries(1000, 1000),
		},
		{
			"1000x40000",
			generateSeries(1000, 10000),
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dataframe.New(test.data...)
			}
		})
	}
}

func BenchmarkDataFrame_Arrange(b *testing.B) {
	data := dataframe.New(generateSeries(100000, 5)...)
	table := []struct {
		name string
		data dataframe.DataFrame
		key  []dataframe.Order
	}{
		{
			"100000x20_1",
			data,
			[]dataframe.Order{dataframe.Sort("X0")},
		},
		{
			"100000x20_2",
			data,
			[]dataframe.Order{
				dataframe.Sort("X0"),
				dataframe.Sort("X1"),
			},
		},
		{
			"100000x20_3",
			data,
			[]dataframe.Order{
				dataframe.Sort("X0"),
				dataframe.Sort("X1"),
				dataframe.Sort("X2"),
			},
		},
	}
	for _, test := range table {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.data.Arrange(test.key...)
			}
		})
	}
}
