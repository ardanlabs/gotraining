package dataframe

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/kniren/gota/series"
)

func checkAddr(addra, addrb []string) error {
	for i := 0; i < len(addra); i++ {
		for j := 0; j < len(addrb); j++ {
			if addra[i] == "<nil>" || addrb[j] == "<nil>" {
				continue
			}
			if addra[i] == addrb[j] {
				return fmt.Errorf("found same address on\nA:%v\nB:%v", i, j)
			}
		}
	}
	return nil
}

func checkAddrDf(a, b DataFrame) error {
	var addra []string
	for _, s := range a.columns {
		addra = append(addra, s.Addr()...)
	}
	var addrb []string
	for _, s := range b.columns {
		addrb = append(addrb, s.Addr()...)
	}
	if err := checkAddr(addra, addrb); err != nil {
		return fmt.Errorf("Error:%v\nA:%v\nB:%v", err, addra, addrb)
	}
	return nil
}

func TestDataFrame_New(t *testing.T) {
	series := []series.Series{
		series.Strings([]int{1, 2, 3, 4, 5}),
		series.New([]int{1, 2, 3, 4, 5}, series.String, "0"),
		series.Ints([]int{1, 2, 3, 4, 5}),
		series.New([]int{1, 2, 3, 4, 5}, series.String, "0"),
		series.New([]int{1, 2, 3, 4, 5}, series.Float, "1"),
		series.New([]int{1, 2, 3, 4, 5}, series.Bool, "1"),
	}
	d := New(series...)

	// Check that the names are renamed properly
	received := d.Names()
	expected := []string{"X0", "0_0", "X1", "0_1", "1_0", "1_1"}
	if !reflect.DeepEqual(received, expected) {
		t.Errorf(
			"Expected:\n%v\nReceived:\n%v",
			expected, received,
		)
	}

	// Check that the memory addresses of the original series and the series
	// inside the DataFrame are different
	var addra []string
	for _, s := range series {
		addra = append(addra, s.Addr()...)
	}
	var addrb []string
	for _, s := range d.columns {
		addrb = append(addrb, s.Addr()...)
	}
	if err := checkAddr(addra, addrb); err != nil {
		t.Errorf("Error:%v\nA:%v\nB:%v", err, addra, addrb)
	}
}

func TestDataFrame_Copy(t *testing.T) {
	a := New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)
	b := a.Copy()

	// Check that there are no shared memory addresses between DataFrames
	if err := checkAddrDf(a, b); err != nil {
		t.Error(err)
	}
	// Check that the types are the same between both DataFrames
	if !reflect.DeepEqual(a.Types(), b.Types()) {
		t.Errorf("Different types:\nA:%v\nB:%v", a.Types(), b.Types())
	}
	// Check that the values are the same between both DataFrames
	if !reflect.DeepEqual(a.Records(), b.Records()) {
		t.Errorf("Different values:\nA:%v\nB:%v", a.Records(), b.Records())
	}
}

func TestDataFrame_Subset(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		indexes interface{}
		expDf   DataFrame
	}{
		{
			[]int{1, 2},
			New(
				series.New([]string{"a", "b"}, series.String, "COL.1"),
				series.New([]int{2, 4}, series.Int, "COL.2"),
				series.New([]float64{4.0, 5.3}, series.Float, "COL.3"),
			),
		},
		{
			[]bool{false, true, true, false, false},
			New(
				series.New([]string{"a", "b"}, series.String, "COL.1"),
				series.New([]int{2, 4}, series.Int, "COL.2"),
				series.New([]float64{4.0, 5.3}, series.Float, "COL.3"),
			),
		},
		{
			series.Ints([]int{1, 2}),
			New(
				series.New([]string{"a", "b"}, series.String, "COL.1"),
				series.New([]int{2, 4}, series.Int, "COL.2"),
				series.New([]float64{4.0, 5.3}, series.Float, "COL.3"),
			),
		},
		{
			[]int{0, 0, 1, 1, 2, 2, 3, 4},
			New(
				series.New([]string{"b", "b", "a", "a", "b", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 1, 2, 2, 4, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 3.0, 4.0, 4.0, 5.3, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
	}
	for testnum, test := range table {
		b := a.Subset(test.indexes)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Select(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		indexes interface{}
		expDf   DataFrame
	}{
		{
			series.Bools([]bool{false, true, true}),
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]bool{false, true, true},
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			series.Ints([]int{1, 2}),
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]int{1, 2},
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]int{1},
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
			),
		},
		{
			1,
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
			),
		},
		{
			[]int{1, 2, 0},
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
			),
		},
		{
			[]int{0, 0},
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
			),
		},
		{
			"COL.3",
			New(
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]string{"COL.3"},
			New(
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]string{"COL.3", "COL.1"},
			New(
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
			),
		},
		{
			series.Strings([]string{"COL.3", "COL.1"}),
			New(
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
			),
		},
	}
	for testnum, test := range table {
		b := a.Select(test.indexes)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Rename(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		newname string
		oldname string
		expDf   DataFrame
	}{
		{
			"NEWCOL.1",
			"COL.1",
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "NEWCOL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			"NEWCOL.3",
			"COL.3",
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "NEWCOL.3"),
			),
		},
		{
			"NEWCOL.2",
			"COL.2",
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "NEWCOL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
	}
	for testnum, test := range table {
		b := a.Rename(test.newname, test.oldname)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_CBind(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		dfb   DataFrame
		expDf DataFrame
	}{
		{
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.5"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.6"),
			),
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.5"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.6"),
			),
		},
		{
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
			),
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
			),
		},
		{
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.6"),
			),
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.4"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.6"),
			),
		},
	}
	for testnum, test := range table {
		b := a.CBind(test.dfb)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_RBind(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		dfb   DataFrame
		expDf DataFrame
	}{
		{
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
			New(
				series.New([]string{"b", "a", "b", "c", "d", "b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4, 1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2, 3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			New(
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
			New(
				series.New([]string{"b", "a", "b", "c", "d", "1", "2", "4", "5", "4"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4, 1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2, 3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
	}
	for testnum, test := range table {
		b := a.RBind(test.dfb)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Records(t *testing.T) {
	a := New(
		series.New([]string{"a", "b", "c"}, series.String, "COL.1"),
		series.New([]int{1, 2, 3}, series.Int, "COL.2"),
		series.New([]float64{3, 2, 1}, series.Float, "COL.3"))
	expected := [][]string{
		[]string{"COL.1", "COL.2", "COL.3"},
		[]string{"a", "1", "3.000000"},
		[]string{"b", "2", "2.000000"},
		[]string{"c", "3", "1.000000"},
	}
	received := a.Records()
	if !reflect.DeepEqual(expected, received) {
		t.Error(
			"Error when saving records.\n",
			"Expected: ", expected, "\n",
			"Received: ", received,
		)
	}
}

func TestDataFrame_Mutate(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		s     series.Series
		expDf DataFrame
	}{
		{
			series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.1"),
			New(
				series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.2"),
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.4"),
			New(
				series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
				series.New([]string{"A", "B", "A", "A", "A"}, series.String, "COL.4"),
			),
		},
	}
	for testnum, test := range table {
		b := a.Mutate(test.s)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Filter(t *testing.T) {
	a := New(
		series.New([]string{"b", "a", "b", "c", "d"}, series.String, "COL.1"),
		series.New([]int{1, 2, 4, 5, 4}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 5.3, 3.2, 1.2}, series.Float, "COL.3"),
	)
	table := []struct {
		filters []F
		expDf   DataFrame
	}{
		{
			[]F{{"COL.2", series.GreaterEq, 4}},
			New(
				series.New([]string{"b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{4, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{5.3, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
		{
			[]F{
				{"COL.2", series.Greater, 4},
				{"COL.2", series.Eq, 1},
			},
			New(
				series.New([]string{"b", "c"}, series.String, "COL.1"),
				series.New([]int{1, 5}, series.Int, "COL.2"),
				series.New([]float64{3.0, 3.2}, series.Float, "COL.3"),
			),
		},
		{
			[]F{
				{"COL.2", series.Greater, 4},
				{"COL.2", series.Eq, 1},
				{"COL.1", series.Eq, "d"},
			},
			New(
				series.New([]string{"b", "c", "d"}, series.String, "COL.1"),
				series.New([]int{1, 5, 4}, series.Int, "COL.2"),
				series.New([]float64{3.0, 3.2, 1.2}, series.Float, "COL.3"),
			),
		},
	}
	for testnum, test := range table {
		b := a.Filter(test.filters...)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestLoadRecords(t *testing.T) {
	table := []struct {
		b     DataFrame
		expDf DataFrame
	}{
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"a", "1", "true", "0"},
					{"b", "2", "true", "0.5"},
				},
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]int{1, 2}, series.Int, "B"),
				series.New([]bool{true, true}, series.Bool, "C"),
				series.New([]float64{0, 0.5}, series.Float, "D"),
			),
		},
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"a", "1", "true", "0"},
					{"b", "2", "true", "0.5"},
				},
				HasHeader(true),
				DetectTypes(false),
				DefaultType(series.String),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]int{1, 2}, series.String, "B"),
				series.New([]bool{true, true}, series.String, "C"),
				series.New([]string{"0", "0.5"}, series.String, "D"),
			),
		},
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"a", "1", "true", "0"},
					{"b", "2", "true", "0.5"},
				},
				HasHeader(false),
				DetectTypes(false),
				DefaultType(series.String),
			),
			New(
				series.New([]string{"A", "a", "b"}, series.String, "X0"),
				series.New([]string{"B", "1", "2"}, series.String, "X1"),
				series.New([]string{"C", "true", "true"}, series.String, "X2"),
				series.New([]string{"D", "0", "0.5"}, series.String, "X3"),
			),
		},
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"a", "1", "true", "0"},
					{"b", "2", "true", "0.5"},
				},
				HasHeader(true),
				DetectTypes(false),
				DefaultType(series.String),
				WithTypes(map[string]series.Type{
					"B": series.Float,
					"C": series.String,
				}),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]float64{1, 2}, series.Float, "B"),
				series.New([]bool{true, true}, series.String, "C"),
				series.New([]string{"0", "0.5"}, series.String, "D"),
			),
		},
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"a", "1", "true", "0"},
					{"b", "2", "true", "0.5"},
				},
				HasHeader(true),
				DetectTypes(true),
				DefaultType(series.String),
				WithTypes(map[string]series.Type{
					"B": series.Float,
				}),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]float64{1, 2}, series.Float, "B"),
				series.New([]bool{true, true}, series.Bool, "C"),
				series.New([]string{"0", "0.5"}, series.Float, "D"),
			),
		},
	}
	for testnum, test := range table {
		b := test.b
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestLoadMaps(t *testing.T) {
	table := []struct {
		b     DataFrame
		expDf DataFrame
	}{
		{
			LoadMaps(
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
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]int{1, 2}, series.Int, "B"),
				series.New([]bool{true, true}, series.Bool, "C"),
				series.New([]float64{0, 0.5}, series.Float, "D"),
			),
		},
		{
			LoadMaps(
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
				HasHeader(true),
				DetectTypes(false),
				DefaultType(series.String),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]int{1, 2}, series.String, "B"),
				series.New([]bool{true, true}, series.String, "C"),
				series.New([]string{"0", "0.5"}, series.String, "D"),
			),
		},
		{
			LoadMaps(
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
				HasHeader(false),
				DetectTypes(false),
				DefaultType(series.String),
			),
			New(
				series.New([]string{"A", "a", "b"}, series.String, "X0"),
				series.New([]string{"B", "1", "2"}, series.String, "X1"),
				series.New([]string{"C", "true", "true"}, series.String, "X2"),
				series.New([]string{"D", "0", "0.5"}, series.String, "X3"),
			),
		},
		{
			LoadMaps(
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
				HasHeader(true),
				DetectTypes(false),
				DefaultType(series.String),
				WithTypes(map[string]series.Type{
					"B": series.Float,
					"C": series.String,
				}),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]float64{1, 2}, series.Float, "B"),
				series.New([]bool{true, true}, series.String, "C"),
				series.New([]string{"0", "0.5"}, series.String, "D"),
			),
		},
		{
			LoadMaps(
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
				HasHeader(true),
				DetectTypes(true),
				DefaultType(series.String),
				WithTypes(map[string]series.Type{
					"B": series.Float,
				}),
			),
			New(
				series.New([]string{"a", "b"}, series.String, "A"),
				series.New([]float64{1, 2}, series.Float, "B"),
				series.New([]bool{true, true}, series.Bool, "C"),
				series.New([]string{"0", "0.5"}, series.Float, "D"),
			),
		},
	}
	for testnum, test := range table {
		b := test.b
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestReadCSV(t *testing.T) {
	// Load the data from a CSV string and try to infer the type of the
	// columns
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
	a := ReadCSV(strings.NewReader(csvStr))
	if a.Err != nil {
		t.Errorf("Expected success, got error: %v", a.Err)
	}
}

func TestReadJSON(t *testing.T) {
	table := []struct {
		jsonStr string
		expDf   DataFrame
	}{
		{
			`[{"COL.1":null,"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`,
			LoadRecords(
				[][]string{
					[]string{"COL.1", "COL.2", "COL.3"},
					[]string{"NaN", "1", "3"},
					[]string{"5", "2", "2"},
					[]string{"6", "3", "1"},
				},
				DetectTypes(false),
				DefaultType(series.Int),
			),
		},
		{
			`[{"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]`,
			LoadRecords(
				[][]string{
					[]string{"COL.1", "COL.2", "COL.3"},
					[]string{"NaN", "1", "3"},
					[]string{"5", "2", "2"},
					[]string{"6", "3", "1"},
				},
				DetectTypes(false),
				DefaultType(series.Int),
			),
		},
	}
	for testnum, test := range table {
		c := ReadJSON(strings.NewReader(test.jsonStr))
		if err := c.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), c.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), c.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), c.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), c.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), c.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), c.Records())
		}
	}
}

func TestDataFrame_SetNames(t *testing.T) {
	a := New(
		series.New([]string{"a", "b", "c"}, series.String, "COL.1"),
		series.New([]int{1, 2, 3}, series.Int, "COL.2"),
		series.New([]float64{3, 2, 1}, series.Float, "COL.3"))
	n := []string{"wot", "tho", "tree"}
	err := a.SetNames(n)
	if err != nil {
		t.Error("Expected success, got error")
	}
	err = a.SetNames([]string{"yaaa"})
	if err == nil {
		t.Error("Expected error, got success")
	}
}

func TestDataFrame_InnerJoin(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "a", "5.1", "true"},
			[]string{"2", "b", "6.0", "true"},
			[]string{"3", "c", "6.0", "false"},
			[]string{"1", "d", "7.1", "false"},
		},
	)
	b := LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "true"},
			[]string{"4", "2", "false"},
			[]string{"2", "8", "false"},
			[]string{"5", "9", "false"},
		},
	)
	table := []struct {
		keys  []string
		expDf DataFrame
	}{
		{
			[]string{"A", "D"},
			LoadRecords(
				[][]string{
					[]string{"A", "D", "B", "C", "F"},
					[]string{"1", "true", "a", "5.1", "1"},
				},
			),
		},
		{
			[]string{"A"},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D_0", "F", "D_1"},
					[]string{"1", "a", "5.1", "true", "1", "true"},
					[]string{"2", "b", "6.0", "true", "8", "false"},
					[]string{"1", "d", "7.1", "false", "1", "true"},
				},
			),
		},
		{
			[]string{"D"},
			LoadRecords(
				[][]string{
					[]string{"D", "A_0", "B", "C", "A_1", "F"},
					[]string{"true", "1", "a", "5.1", "1", "1"},
					[]string{"true", "2", "b", "6.0", "1", "1"},
					[]string{"false", "3", "c", "6.0", "4", "2"},
					[]string{"false", "3", "c", "6.0", "2", "8"},
					[]string{"false", "3", "c", "6.0", "5", "9"},
					[]string{"false", "1", "d", "7.1", "4", "2"},
					[]string{"false", "1", "d", "7.1", "2", "8"},
					[]string{"false", "1", "d", "7.1", "5", "9"},
				},
			),
		},
	}
	for testnum, test := range table {
		c := a.InnerJoin(b, test.keys...)
		if err := c.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), c.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), c.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), c.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), c.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), c.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), c.Records())
		}
	}
}

func TestDataFrame_LeftJoin(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "1"},
			[]string{"2", "4", "6.0", "1"},
			[]string{"3", "3", "6.0", "0"},
			[]string{"1", "2", "7.1", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	b := LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "1"},
			[]string{"4", "2", "0"},
			[]string{"2", "8", "0"},
			[]string{"5", "9", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	table := []struct {
		keys  []string
		expDf DataFrame
	}{
		{
			[]string{"A", "D"},
			LoadRecords(
				[][]string{
					[]string{"A", "D", "B", "C", "F"},
					[]string{"1", "1", "4", "5.1", "1"},
					[]string{"2", "1", "4", "6.0", "NaN"},
					[]string{"3", "0", "3", "6.0", "NaN"},
					[]string{"1", "0", "2", "7.1", "NaN"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
		{
			[]string{"A"},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D_0", "F", "D_1"},
					[]string{"1", "4", "5.1", "1", "1", "1"},
					[]string{"2", "4", "6.0", "1", "8", "0"},
					[]string{"3", "3", "6.0", "0", "NaN", "NaN"},
					[]string{"1", "2", "7.1", "0", "1", "1"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
	}
	for testnum, test := range table {
		c := a.LeftJoin(b, test.keys...)
		if err := c.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), c.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), c.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), c.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), c.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), c.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), c.Records())
		}
	}
}

func TestDataFrame_RightJoin(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "1"},
			[]string{"4", "2", "0"},
			[]string{"2", "8", "0"},
			[]string{"5", "9", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	b := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "1"},
			[]string{"2", "4", "6.0", "1"},
			[]string{"3", "3", "6.0", "0"},
			[]string{"1", "2", "7.1", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	table := []struct {
		keys  []string
		expDf DataFrame
	}{
		{
			[]string{"A", "D"},
			LoadRecords(
				[][]string{
					[]string{"A", "D", "F", "B", "C"},
					[]string{"1", "1", "1", "4", "5.1"},
					[]string{"2", "1", "NaN", "4", "6.0"},
					[]string{"3", "0", "NaN", "3", "6.0"},
					[]string{"1", "0", "NaN", "2", "7.1"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
		{
			[]string{"A"},
			LoadRecords(
				[][]string{
					[]string{"A", "F", "D_0", "B", "C", "D_1"},
					[]string{"1", "1", "1", "4", "5.1", "1"},
					[]string{"2", "8", "0", "4", "6.0", "1"},
					[]string{"1", "1", "1", "2", "7.1", "0"},
					[]string{"3", "NaN", "NaN", "3", "6.0", "0"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
	}
	for testnum, test := range table {
		c := a.RightJoin(b, test.keys...)
		if err := c.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), c.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), c.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), c.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), c.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), c.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), c.Records())
		}
	}
}

func TestDataFrame_OuterJoin(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "1"},
			[]string{"2", "4", "6.0", "1"},
			[]string{"3", "3", "6.0", "0"},
			[]string{"1", "2", "7.1", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	b := LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "1"},
			[]string{"4", "2", "0"},
			[]string{"2", "8", "0"},
			[]string{"5", "9", "0"},
		},
		DetectTypes(false),
		DefaultType(series.Float),
	)
	table := []struct {
		keys  []string
		expDf DataFrame
	}{
		{
			[]string{"A", "D"},
			LoadRecords(
				[][]string{
					[]string{"A", "D", "B", "C", "F"},
					[]string{"1", "1", "4", "5.1", "1"},
					[]string{"2", "1", "4", "6.0", "NaN"},
					[]string{"3", "0", "3", "6.0", "NaN"},
					[]string{"1", "0", "2", "7.1", "NaN"},
					[]string{"4", "0", "NaN", "NaN", "2"},
					[]string{"2", "0", "NaN", "NaN", "8"},
					[]string{"5", "0", "NaN", "NaN", "9"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
		{
			[]string{"A"},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D_0", "F", "D_1"},
					[]string{"1", "4", "5.1", "1", "1", "1"},
					[]string{"2", "4", "6.0", "1", "8", "0"},
					[]string{"3", "3", "6.0", "0", "NaN", "NaN"},
					[]string{"1", "2", "7.1", "0", "1", "1"},
					[]string{"4", "NaN", "NaN", "NaN", "2", "0"},
					[]string{"5", "NaN", "NaN", "NaN", "9", "0"},
				},
				DetectTypes(false),
				DefaultType(series.Float),
			),
		},
	}
	for testnum, test := range table {
		c := a.OuterJoin(b, test.keys...)
		if err := c.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), c.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), c.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), c.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), c.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), c.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), c.Records())
		}
	}
}

func TestDataFrame_CrossJoin(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "a", "5.1", "true"},
			[]string{"2", "b", "6.0", "true"},
			[]string{"3", "c", "6.0", "false"},
			[]string{"1", "d", "7.1", "false"},
		},
	)
	b := LoadRecords(
		[][]string{
			[]string{"A", "F", "D"},
			[]string{"1", "1", "true"},
			[]string{"4", "2", "false"},
			[]string{"2", "8", "false"},
			[]string{"5", "9", "false"},
		},
	)
	c := a.CrossJoin(b)
	expectedCSV := `
A_0,B,C,D_0,A_1,F,D_1
1,a,5.1,true,1,1,true
1,a,5.1,true,4,2,false
1,a,5.1,true,2,8,false
1,a,5.1,true,5,9,false
2,b,6.0,true,1,1,true
2,b,6.0,true,4,2,false
2,b,6.0,true,2,8,false
2,b,6.0,true,5,9,false
3,c,6.0,false,1,1,true
3,c,6.0,false,4,2,false
3,c,6.0,false,2,8,false
3,c,6.0,false,5,9,false
1,d,7.1,false,1,1,true
1,d,7.1,false,4,2,false
1,d,7.1,false,2,8,false
1,d,7.1,false,5,9,false
`
	expected := ReadCSV(
		strings.NewReader(expectedCSV),
		WithTypes(map[string]series.Type{
			"A.1": series.String,
		}))
	if err := c.Err; err != nil {
		t.Errorf("Error:%v", err)
	}
	// Check that the types are the same between both DataFrames
	if !reflect.DeepEqual(expected.Types(), c.Types()) {
		t.Errorf("Different types:\nA:%v\nB:%v", expected.Types(), c.Types())
	}
	// Check that the colnames are the same between both DataFrames
	if !reflect.DeepEqual(expected.Names(), c.Names()) {
		t.Errorf("Different colnames:\nA:%v\nB:%v", expected.Names(), c.Names())
	}
	// Check that the values are the same between both DataFrames
	if !reflect.DeepEqual(expected.Records(), c.Records()) {
		t.Errorf("Different values:\nA:%v\nB:%v", expected.Records(), c.Records())
	}
}

func TestDataFrame_Maps(t *testing.T) {
	a := New(
		series.New([]string{"a", "b", "c"}, series.String, "COL.1"),
		series.New([]string{"", "2", "3"}, series.Int, "COL.2"),
		series.New([]string{"", "", "3"}, series.Int, "COL.3"),
	)
	m := a.Maps()
	expected := []map[string]interface{}{
		map[string]interface{}{
			"COL.1": "a",
			"COL.2": nil,
			"COL.3": nil,
		},
		map[string]interface{}{
			"COL.1": "b",
			"COL.2": 2,
			"COL.3": nil,
		},
		map[string]interface{}{
			"COL.1": "c",
			"COL.2": 3,
			"COL.3": 3,
		},
	}
	if !reflect.DeepEqual(expected, m) {
		t.Errorf("Different values:\nA:%v\nB:%v", expected, m)
	}
}

func TestDataFrame_WriteCSV(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"COL.1", "COL.2", "COL.3"},
			[]string{"NaN", "1", "3"},
			[]string{"b", "2", "2"},
			[]string{"c", "3", "1"},
		},
	)
	buf := new(bytes.Buffer)
	err := a.WriteCSV(buf)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	expected := `COL.1,COL.2,COL.3
NaN,1,3
b,2,2
c,3,1
`
	if expected != buf.String() {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, buf.String())
	}
}

func TestDataFrame_WriteJSON(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"COL.1", "COL.2", "COL.3"},
			[]string{"NaN", "1", "3"},
			[]string{"5", "2", "2"},
			[]string{"6", "3", "1"},
		},
		DetectTypes(false),
		DefaultType(series.Int),
	)
	buf := new(bytes.Buffer)
	err := a.WriteJSON(buf)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	expected := `[{"COL.1":null,"COL.2":1,"COL.3":3},{"COL.1":5,"COL.2":2,"COL.3":2},{"COL.1":6,"COL.2":3,"COL.3":1}]
`
	if expected != buf.String() {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, buf.String())
	}
}

func TestDataFrame_Col(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"COL.1", "COL.2", "COL.3"},
			[]string{"NaN", "1", "3"},
			[]string{"5", "2", "2"},
			[]string{"6", "3", "1"},
		},
		DetectTypes(false),
		DefaultType(series.Int),
	)
	b := a.Col("COL.2")
	expected := series.New([]int{1, 2, 3}, series.Int, "COL.2")
	if !reflect.DeepEqual(b.Records(), expected.Records()) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, b)
	}
}

func TestDataFrame_Set(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"b", "4", "6.0", "true"},
			[]string{"c", "3", "6.0", "false"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	table := []struct {
		indexes   series.Indexes
		newvalues DataFrame
		expDf     DataFrame
	}{
		{
			series.Ints([]int{0, 2}),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"k", "4", "6.0", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			series.Ints(0),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			series.Bools([]bool{true, false, false, false}),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			series.Bools([]bool{false, true, true, false}),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]int{0, 2},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"k", "4", "6.0", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			0,
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]bool{true, false, false, false},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]bool{false, true, true, false},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
				},
			),
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"k", "5", "7.0", "true"},
					[]string{"k", "4", "6.0", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
	}
	for testnum, test := range table {
		a := a.Copy()
		b := a.Set(test.indexes, test.newvalues)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Arrange(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"b", "4", "6.0", "true"},
			[]string{"c", "3", "6.0", "false"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	table := []struct {
		colnames []Order
		expDf    DataFrame
	}{
		{
			[]Order{Sort("A")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
				},
			),
		},
		{
			[]Order{Sort("B")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"b", "4", "6.0", "true"},
				},
			),
		},
		{
			[]Order{Sort("A"), Sort("B")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
				},
			),
		},
		{
			[]Order{Sort("B"), Sort("A")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"b", "4", "6.0", "true"},
				},
			),
		},
		{
			[]Order{RevSort("A")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]Order{RevSort("B")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]Order{Sort("A"), RevSort("B")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"c", "3", "6.0", "false"},
				},
			),
		},
		{
			[]Order{Sort("B"), RevSort("A")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"a", "2", "7.1", "false"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"a", "4", "5.1", "true"},
				},
			),
		},
		{
			[]Order{RevSort("B"), RevSort("A")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
		{
			[]Order{RevSort("A"), RevSort("B")},
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"c", "3", "6.0", "false"},
					[]string{"b", "4", "6.0", "true"},
					[]string{"a", "4", "5.1", "true"},
					[]string{"a", "2", "7.1", "false"},
				},
			),
		},
	}
	for testnum, test := range table {
		b := a.Arrange(test.colnames...)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Capply(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"b", "4", "6.0", "true"},
			[]string{"c", "3", "6.0", "false"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}
	sum := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum)
	}
	table := []struct {
		fun   func(series.Series) series.Series
		expDf DataFrame
	}{
		{
			mean,
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"NaN", "3.25", "6.05", "0.5"},
				},
				DefaultType(series.Float),
				DetectTypes(false),
			),
		},
		{
			sum,
			LoadRecords(
				[][]string{
					[]string{"A", "B", "C", "D"},
					[]string{"NaN", "13", "24.2", "2"},
				},
				DefaultType(series.Float),
				DetectTypes(false),
			),
		},
	}
	for testnum, test := range table {
		b := a.Capply(test.fun)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_String(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "C", "D"},
			[]string{"1", "5.1", "true"},
			[]string{"NaN", "6.0", "true"},
			[]string{"2", "6.0", "false"},
			[]string{"2", "7.1", "false"},
		},
	)
	received := a.String()
	expected := `[4x3] DataFrame

    A     C        D     
 0: 1     5.100000 true  
 1: NaN   6.000000 true  
 2: 2     6.000000 false 
 3: 2     7.100000 false 
    <int> <float>  <bool>
`
	if expected != received {
		t.Errorf("Different values:\nExpected: \n%v\nReceived: \n%v\n", expected, received)
	}
}

func TestDataFrame_Rapply(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "true"},
			[]string{"1", "4", "6.0", "true"},
			[]string{"2", "3", "6.0", "false"},
			[]string{"2", "2", "7.1", "false"},
		},
	)
	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		ret := series.Floats(sum / float64(len(floats)))
		return ret
	}
	sum := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum)
	}
	table := []struct {
		fun   func(series.Series) series.Series
		expDf DataFrame
	}{
		{
			mean,
			LoadRecords(
				[][]string{
					[]string{"X0"},
					[]string{"2.775"},
					[]string{"3"},
					[]string{"2.75"},
					[]string{"2.775"},
				},
				DefaultType(series.Float),
				DetectTypes(false),
			),
		},
		{
			sum,
			LoadRecords(
				[][]string{
					[]string{"X0"},
					[]string{"11.1"},
					[]string{"12"},
					[]string{"11"},
					[]string{"11.1"},
				},
				DefaultType(series.Float),
				DetectTypes(false),
			),
		},
	}
	for testnum, test := range table {
		b := a.Rapply(test.fun)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddrDf(a, b); err != nil {
			t.Error(err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}

func TestDataFrame_Matrix(t *testing.T) {
	a := LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"1", "4", "5.1", "true"},
			[]string{"1", "4", "6.0", "true"},
			[]string{"2", "3", "6.0", "false"},
			[]string{"2", "2", "7.1", "false"},
		},
	)
	m := a.Matrix()
	sum := mat64.Sum(m)
	expected := 45.2
	if sum != expected {
		t.Errorf("\nExpected: %v\nReceived: %v", expected, sum)
	}
}

func TestLoadMatrix(t *testing.T) {
	table := []struct {
		b     DataFrame
		expDf DataFrame
	}{
		{
			LoadRecords(
				[][]string{
					{"A", "B", "C", "D"},
					{"4", "1", "true", "0"},
					{"3", "2", "true", "0.5"},
				},
			),
			New(
				series.New([]string{"4", "3"}, series.Float, "X0"),
				series.New([]int{1, 2}, series.Float, "X1"),
				series.New([]bool{true, true}, series.Float, "X2"),
				series.New([]float64{0, 0.5}, series.Float, "X3"),
			),
		},
	}
	for testnum, test := range table {
		b := LoadMatrix(test.b.Matrix())
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		// Check that the types are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Types(), b.Types()) {
			t.Errorf("Different types:\nA:%v\nB:%v", test.expDf.Types(), b.Types())
		}
		// Check that the colnames are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Names(), b.Names()) {
			t.Errorf("Different colnames:\nA:%v\nB:%v", test.expDf.Names(), b.Names())
		}
		// Check that the values are the same between both DataFrames
		if !reflect.DeepEqual(test.expDf.Records(), b.Records()) {
			t.Errorf("Different values:\nA:%v\nB:%v", test.expDf.Records(), b.Records())
		}
	}
}
