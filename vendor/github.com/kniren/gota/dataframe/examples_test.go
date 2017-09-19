package dataframe_test

import (
	"fmt"
	"strings"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

func ExampleNew() {
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)
	fmt.Println(df)
}

func ExampleLoadRecords() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	fmt.Println(df)
}

func ExampleLoadRecords_options() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
		dataframe.DetectTypes(false),
		dataframe.DefaultType(series.Float),
		dataframe.WithTypes(map[string]series.Type{
			"A": series.String,
			"D": series.Bool,
		}),
	)
	fmt.Println(df)
}

func ExampleLoadMaps() {
	df := dataframe.LoadMaps(
		[]map[string]interface{}{
			map[string]interface{}{
				"A": "a",
				"B": 1,
				"C": true,
				"D": 0,
			},
			map[string]interface{}{
				"A": "b",
				"B": 2,
				"C": true,
				"D": 0.5,
			},
		},
	)
	fmt.Println(df)
}

func ExampleReadCSV() {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := dataframe.ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
}

func ExampleReadJSON() {
	jsonStr := `[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`
	df := dataframe.ReadJSON(strings.NewReader(jsonStr))
	fmt.Println(df)
}

func ExampleDataFrame_Subset() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	sub := df.Subset([]int{0, 2})
	fmt.Println(sub)
}

func ExampleDataFrame_Select() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	sel1 := df.Select([]int{0, 2})
	sel2 := df.Select([]string{"A", "C"})
	fmt.Println(sel1)
	fmt.Println(sel2)
}

func ExampleDataFrame_Filter() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	fil := df.Filter(
		dataframe.F{
			Colname:    "A",
			Comparator: series.Eq,
			Comparando: "a",
		},
		dataframe.F{
			Colname:    "B",
			Comparator: series.Greater,
			Comparando: 4,
		},
	)
	fil2 := fil.Filter(
		dataframe.F{
			Colname:    "D",
			Comparator: series.Eq,
			Comparando: true,
		},
	)
	fmt.Println(fil)
	fmt.Println(fil2)
}

func ExampleDataFrame_Mutate() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	// Change column C with a new one
	mut := df.Mutate(
		series.New([]string{"a", "b", "c", "d"}, series.String, "C"),
	)
	// Add a new column E
	mut2 := df.Mutate(
		series.New([]string{"a", "b", "c", "d"}, series.String, "E"),
	)
	fmt.Println(mut)
	fmt.Println(mut2)
}

func ExampleDataFrame_InnerJoin() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	df2 := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "true"},
			[]string{"4", "2", "false"},
			[]string{"2", "8", "false"},
			[]string{"5", "9", "false"},
		},
	)
	join := df.InnerJoin(df2, "D")
	fmt.Println(join)
}

func ExampleDataFrame_Set() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	df2 := df.Set(
		series.Ints([]int{0, 2}),
		dataframe.LoadRecords(
			[][]string{
				[]string{"A", "B", "C", "D"},
				[]string{"b", "4", "6.0", "true"},
				[]string{"c", "3", "6.0", "false"},
			},
		),
	)
	fmt.Println(df2)
}

func ExampleDataFrame_Arrange() {
	df := dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"b", "4", "6.0", "true"},
			[]string{"c", "3", "6.0", "false"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	sorted := df.Arrange(
		dataframe.Sort("A"),
		dataframe.RevSort("B"),
	)
	fmt.Println(sorted)
}
