package series

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// Check that there are no shared memory addreses between the elements of two Series
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

// Check that all the types on a Series are the same type and that it matches with
// Series.t
func checkTypes(s Series) error {
	var types []Type
	for _, e := range s.elements {
		types = append(types, e.Type())
	}
	for _, t := range types {
		if t != s.t {
			return fmt.Errorf("bad types for %v Series:\n%v", s.t, types)
		}
	}
	return nil
}

func TestSeries_Compare(t *testing.T) {
	table := []struct {
		series     Series
		comparator Comparator
		comparando interface{}
		expected   Series
	}{
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Eq,
			"B",
			Bools([]bool{false, true, false, true, false, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Eq,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{false, true, true, false, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Eq,
			"2",
			Bools([]bool{false, true, false, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Eq,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{true, true, false, true, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Eq,
			"2",
			Bools([]bool{false, true, false, false, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Eq,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{true, true, false, true, false}),
		},
		{
			Bools([]bool{true, true, false}),
			Eq,
			"true",
			Bools([]bool{true, true, false}),
		},
		{
			Bools([]bool{true, true, false}),
			Eq,
			[]bool{true, false, false},
			Bools([]bool{true, false, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Neq,
			"B",
			Bools([]bool{true, false, true, false, true, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Neq,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{true, false, false, true, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Neq,
			"2",
			Bools([]bool{true, false, true, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Neq,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{false, false, true, false, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Neq,
			"2",
			Bools([]bool{true, false, true, true, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Neq,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{false, false, true, false, true}),
		},
		{
			Bools([]bool{true, true, false}),
			Neq,
			"true",
			Bools([]bool{false, false, true}),
		},
		{
			Bools([]bool{true, true, false}),
			Neq,
			[]bool{true, false, false},
			Bools([]bool{false, true, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Greater,
			"B",
			Bools([]bool{false, false, true, false, true, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Greater,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{false, false, false, false, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Greater,
			"2",
			Bools([]bool{false, false, false, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Greater,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{false, false, true, false, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Greater,
			"2",
			Bools([]bool{false, false, false, true, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Greater,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{false, false, true, false, false}),
		},
		{
			Bools([]bool{true, true, false}),
			Greater,
			"true",
			Bools([]bool{false, false, false}),
		},
		{
			Bools([]bool{true, true, false}),
			Greater,
			[]bool{true, false, false},
			Bools([]bool{false, true, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			GreaterEq,
			"B",
			Bools([]bool{false, true, true, true, true, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			GreaterEq,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{false, true, true, false, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			GreaterEq,
			"2",
			Bools([]bool{false, true, false, true, true}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			GreaterEq,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{true, true, true, true, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			GreaterEq,
			"2",
			Bools([]bool{false, true, false, true, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			GreaterEq,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{true, true, true, true, false}),
		},
		{
			Bools([]bool{true, true, false}),
			GreaterEq,
			"true",
			Bools([]bool{true, true, false}),
		},
		{
			Bools([]bool{true, true, false}),
			GreaterEq,
			[]bool{true, false, false},
			Bools([]bool{true, true, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Less,
			"B",
			Bools([]bool{true, false, false, false, false, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			Less,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{true, false, false, true, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Less,
			"2",
			Bools([]bool{true, false, true, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			Less,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{false, false, false, false, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Less,
			"2",
			Bools([]bool{true, false, true, false, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			Less,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{false, false, false, false, true}),
		},
		{
			Bools([]bool{true, true, false}),
			Less,
			"true",
			Bools([]bool{false, false, true}),
		},
		{
			Bools([]bool{true, true, false}),
			Less,
			[]bool{true, false, false},
			Bools([]bool{false, false, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			LessEq,
			"B",
			Bools([]bool{true, true, false, true, false, false}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			LessEq,
			[]string{"B", "B", "C", "D", "A", "A"},
			Bools([]bool{true, true, true, true, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			LessEq,
			"2",
			Bools([]bool{true, true, true, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			LessEq,
			[]int{0, 2, 0, 5, 10},
			Bools([]bool{true, true, false, true, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			LessEq,
			"2",
			Bools([]bool{true, true, true, false, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			LessEq,
			[]float64{0.1, 2, 0, 5, 10},
			Bools([]bool{true, true, false, true, true}),
		},
		{
			Bools([]bool{true, true, false}),
			LessEq,
			"true",
			Bools([]bool{true, true, true}),
		},
		{
			Bools([]bool{true, true, false}),
			LessEq,
			[]bool{true, false, false},
			Bools([]bool{true, false, true}),
		},
		{
			Strings([]string{"A", "B", "C", "B", "D", "BADA"}),
			In,
			"B",
			Bools([]bool{false, true, false, true, false, false}),
		},
		{
			Strings([]string{"Hello", "world", "this", "is", "a", "test"}),
			In,
			[]string{"cat", "world", "hello", "a"},
			Bools([]bool{false, true, false, false, true, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			In,
			"2",
			Bools([]bool{false, true, false, false, false}),
		},
		{
			Ints([]int{0, 2, 1, 5, 9}),
			In,
			[]int{2, 99, 1234, 9},
			Bools([]bool{false, true, false, false, true}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			In,
			"2",
			Bools([]bool{false, true, false, false, false}),
		},
		{
			Floats([]float64{0.1, 2, 1, 5, 9}),
			In,
			[]float64{2, 99, 1234, 9},
			Bools([]bool{false, true, false, false, true}),
		},
		{
			Bools([]bool{true, true, false}),
			In,
			"true",
			Bools([]bool{true, true, false}),
		},
		{
			Bools([]bool{true, true, false}),
			In,
			[]bool{false, false, false},
			Bools([]bool{false, false, true}),
		},
	}
	for testnum, test := range table {
		a := test.series
		b := a.Compare(test.comparator, test.comparando)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected.Records()
		received := b.Records()
		if !reflect.DeepEqual(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(b); err != nil {
			t.Errorf(
				"Test:%v\nError:%v",
				testnum, err,
			)
		}
		if err := checkAddr(a.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nA:%v\nB:%v", testnum, err, a.Addr(), b.Addr())
		}
	}
}

func TestSeries_Subset(t *testing.T) {
	table := []struct {
		series   Series
		indexes  Indexes
		expected string
	}{
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			[]int{2, 1, 4, 4, 0, 3},
			"[C B D D A K]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			int(1),
			"[B]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			[]bool{true, false, false, true, true},
			"[A K D]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Ints([]int{3, 2, 1, 0}),
			"[K C B A]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Ints([]int{1}),
			"[B]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Ints(2),
			"[C]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Bools([]bool{true, false, false, true, true}),
			"[A K D]",
		},
	}
	for testnum, test := range table {
		a := test.series
		b := test.series.Subset(test.indexes)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(b)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(b); err != nil {
			t.Errorf(
				"Test:%v\nError:%v",
				testnum, err,
			)
		}
		if err := checkAddr(a.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nA:%v\nB:%v", testnum, err, a.Addr(), b.Addr())
		}
	}
}

func TestSeries_Set(t *testing.T) {
	table := []struct {
		series   Series
		indexes  Indexes
		values   Series
		expected string
	}{
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			[]int{1, 2, 4},
			Ints([]string{"1", "2", "3"}),
			"[A 1 2 K 3]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			[]bool{false, true, true, false, true},
			Ints([]string{"1", "2", "3"}),
			"[A 1 2 K 3]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Ints([]int{1, 2, 4}),
			Ints([]string{"1", "2", "3"}),
			"[A 1 2 K 3]",
		},
		{
			Strings([]string{"A", "B", "C", "K", "D"}),
			Bools([]bool{false, true, true, false, true}),
			Ints([]string{"1", "2", "3"}),
			"[A 1 2 K 3]",
		},
	}
	for testnum, test := range table {
		b := test.series.Set(test.indexes, test.values)
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(b)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(b); err != nil {
			t.Errorf(
				"Test:%v\nError:%v",
				testnum, err,
			)
		}
		if err := checkAddr(test.values.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nNV:%v\nB:%v", testnum, err, test.values.Addr(), b.Addr())
		}
	}
}

func TestStrings(t *testing.T) {
	table := []struct {
		series   Series
		expected string
	}{
		{
			Strings([]string{"A", "B", "C", "D"}),
			"[A B C D]",
		},
		{
			Strings([]string{"A"}),
			"[A]",
		},
		{
			Strings("A"),
			"[A]",
		},
		{
			Strings([]int{1, 2, 3}),
			"[1 2 3]",
		},
		{
			Strings([]int{2}),
			"[2]",
		},
		{
			Strings(-1),
			"[-1]",
		},
		{
			Strings([]float64{1, 2, 3}),
			"[1.000000 2.000000 3.000000]",
		},
		{
			Strings([]float64{2}),
			"[2.000000]",
		},
		{
			Strings(-1.0),
			"[-1.000000]",
		},
		{
			Strings(math.NaN()),
			"[NaN]",
		},
		{
			Strings(math.Inf(1)),
			"[+Inf]",
		},
		{
			Strings(math.Inf(-1)),
			"[-Inf]",
		},
		{
			Strings([]bool{true, true, false}),
			"[true true false]",
		},
		{
			Strings([]bool{false}),
			"[false]",
		},
		{
			Strings(true),
			"[true]",
		},
		{
			Strings([]int{}),
			"[]",
		},
		{
			Strings(nil),
			"[NaN]",
		},
		{
			Strings(Strings([]string{"A", "B", "C"})),
			"[A B C]",
		},
	}
	for testnum, test := range table {
		if err := test.series.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(test.series)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(test.series); err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
	}
}

func TestInts(t *testing.T) {
	table := []struct {
		series   Series
		expected string
	}{
		{
			Ints([]string{"A", "B", "1", "2"}),
			"[NaN NaN 1 2]",
		},
		{
			Ints([]string{"1"}),
			"[1]",
		},
		{
			Ints("2"),
			"[2]",
		},
		{
			Ints([]int{1, 2, 3}),
			"[1 2 3]",
		},
		{
			Ints([]int{2}),
			"[2]",
		},
		{
			Ints(-1),
			"[-1]",
		},
		{
			Ints([]float64{1, 2, 3}),
			"[1 2 3]",
		},
		{
			Ints([]float64{2}),
			"[2]",
		},
		{
			Ints(-1.0),
			"[-1]",
		},
		{
			Ints(math.NaN()),
			"[NaN]",
		},
		{
			Ints(math.Inf(1)),
			"[NaN]",
		},
		{
			Ints(math.Inf(-1)),
			"[NaN]",
		},
		{
			Ints([]bool{true, true, false}),
			"[1 1 0]",
		},
		{
			Ints([]bool{false}),
			"[0]",
		},
		{
			Ints(true),
			"[1]",
		},
		{
			Ints([]int{}),
			"[]",
		},
		{
			Ints(nil),
			"[NaN]",
		},
		{
			Ints(Strings([]string{"1", "2", "3"})),
			"[1 2 3]",
		},
		{
			Ints(Ints([]string{"1", "2", "3"})),
			"[1 2 3]",
		},
	}
	for testnum, test := range table {
		if err := test.series.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(test.series)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(test.series); err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
	}
}

func TestFloats(t *testing.T) {
	table := []struct {
		series   Series
		expected string
	}{
		{
			Floats([]string{"A", "B", "1", "2"}),
			"[NaN NaN 1.000000 2.000000]",
		},
		{
			Floats([]string{"1"}),
			"[1.000000]",
		},
		{
			Floats("2.1"),
			"[2.100000]",
		},
		{
			Floats([]int{1, 2, 3}),
			"[1.000000 2.000000 3.000000]",
		},
		{
			Floats([]int{2}),
			"[2.000000]",
		},
		{
			Floats(-1),
			"[-1.000000]",
		},
		{
			Floats([]float64{1.1, 2, 3}),
			"[1.100000 2.000000 3.000000]",
		},
		{
			Floats([]float64{2}),
			"[2.000000]",
		},
		{
			Floats(-1.0),
			"[-1.000000]",
		},
		{
			Floats(math.NaN()),
			"[NaN]",
		},
		{
			Floats(math.Inf(1)),
			"[+Inf]",
		},
		{
			Floats(math.Inf(-1)),
			"[-Inf]",
		},
		{
			Floats([]bool{true, true, false}),
			"[1.000000 1.000000 0.000000]",
		},
		{
			Floats([]bool{false}),
			"[0.000000]",
		},
		{
			Floats(true),
			"[1.000000]",
		},
		{
			Floats([]int{}),
			"[]",
		},
		{
			Floats(nil),
			"[NaN]",
		},
		{
			Floats(Strings([]string{"1", "2", "3"})),
			"[1.000000 2.000000 3.000000]",
		},
	}
	for testnum, test := range table {
		if err := test.series.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(test.series)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(test.series); err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
	}
}

func TestBools(t *testing.T) {
	table := []struct {
		series   Series
		expected string
	}{
		{
			Bools([]string{"A", "true", "1", "f"}),
			"[NaN true true false]",
		},
		{
			Bools([]string{"t"}),
			"[true]",
		},
		{
			Bools("False"),
			"[false]",
		},
		{
			Bools([]int{1, 2, 0}),
			"[true NaN false]",
		},
		{
			Bools([]int{1}),
			"[true]",
		},
		{
			Bools(-1),
			"[NaN]",
		},
		{
			Bools([]float64{1, 2, 0}),
			"[true NaN false]",
		},
		{
			Bools([]float64{0}),
			"[false]",
		},
		{
			Bools(-1.0),
			"[NaN]",
		},
		{
			Bools(math.NaN()),
			"[NaN]",
		},
		{
			Bools(math.Inf(1)),
			"[NaN]",
		},
		{
			Bools(math.Inf(-1)),
			"[NaN]",
		},
		{
			Bools([]bool{true, true, false}),
			"[true true false]",
		},
		{
			Bools([]bool{false}),
			"[false]",
		},
		{
			Bools(true),
			"[true]",
		},
		{
			Bools([]int{}),
			"[]",
		},
		{
			Bools(nil),
			"[NaN]",
		},
		{
			Bools(Strings([]string{"1", "0", "1"})),
			"[true false true]",
		},
	}
	for testnum, test := range table {
		if err := test.series.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		expected := test.expected
		received := fmt.Sprint(test.series)
		if expected != received {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		if err := checkTypes(test.series); err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
	}
}

func TestSeries_Copy(t *testing.T) {
	tests := []Series{
		Strings([]string{"1", "2", "3", "a", "b", "c"}),
		Ints([]string{"1", "2", "3", "a", "b", "c"}),
		Floats([]string{"1", "2", "3", "a", "b", "c"}),
		Bools([]string{"1", "0", "1", "t", "f", "c"}),
	}
	for testnum, test := range tests {
		a := test
		b := a.Copy()
		if fmt.Sprint(a) != fmt.Sprint(b) {
			t.Error("Different values when copying String elements")
		}
		if err := b.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkTypes(b); err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		if err := checkAddr(a.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nA:%v\nB:%v", testnum, err, a.Addr(), b.Addr())
		}
	}
}

func TestSeries_Records(t *testing.T) {
	tests := []struct {
		series   Series
		expected []string
	}{
		{
			Strings([]string{"1", "2", "3", "a", "b", "c"}),
			[]string{"1", "2", "3", "a", "b", "c"},
		},
		{
			Ints([]string{"1", "2", "3", "a", "b", "c"}),
			[]string{"1", "2", "3", "NaN", "NaN", "NaN"},
		},
		{
			Floats([]string{"1", "2", "3", "a", "b", "c"}),
			[]string{"1.000000", "2.000000", "3.000000", "NaN", "NaN", "NaN"},
		},
		{
			Bools([]string{"1", "0", "1", "t", "f", "c"}),
			[]string{"true", "false", "true", "true", "false", "NaN"},
		},
	}
	for testnum, test := range tests {
		expected := test.expected
		received := test.series.Records()
		if !reflect.DeepEqual(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
	}
}

func TestSeries_Float(t *testing.T) {
	precision := 0.0000001
	floatEquals := func(x, y []float64) bool {
		if len(x) != len(y) {
			return false
		}
		for i := 0; i < len(x); i++ {
			a := x[i]
			b := y[i]
			if (a-b) > precision || (b-a) > precision {
				return false
			}
		}
		return true
	}
	tests := []struct {
		series   Series
		expected []float64
	}{
		{
			Strings([]string{"1", "2", "3", "a", "b", "c"}),
			[]float64{1, 2, 3, math.NaN(), math.NaN(), math.NaN()},
		},
		{
			Ints([]string{"1", "2", "3", "a", "b", "c"}),
			[]float64{1, 2, 3, math.NaN(), math.NaN(), math.NaN()},
		},
		{
			Floats([]string{"1", "2", "3", "a", "b", "c"}),
			[]float64{1, 2, 3, math.NaN(), math.NaN(), math.NaN()},
		},
		{
			Bools([]string{"1", "0", "1", "t", "f", "c"}),
			[]float64{1, 0, 1, 1, 0, math.NaN()},
		},
	}
	for testnum, test := range tests {
		expected := test.expected
		received := test.series.Float()
		if !floatEquals(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
	}
}

func TestSeries_Concat(t *testing.T) {
	tests := []struct {
		a        Series
		b        Series
		expected []string
	}{
		{
			Strings([]string{"1", "2", "3"}),
			Strings([]string{"a", "b", "c"}),
			[]string{"1", "2", "3", "a", "b", "c"},
		},
		{
			Ints([]string{"1", "2", "3"}),
			Ints([]string{"a", "4", "c"}),
			[]string{"1", "2", "3", "NaN", "4", "NaN"},
		},
		{
			Floats([]string{"1", "2", "3"}),
			Floats([]string{"a", "4", "c"}),
			[]string{"1.000000", "2.000000", "3.000000", "NaN", "4.000000", "NaN"},
		},
		{
			Bools([]string{"1", "1", "0"}),
			Bools([]string{"0", "0", "0"}),
			[]string{"true", "true", "false", "false", "false", "false"},
		},
	}
	for testnum, test := range tests {
		ab := test.a.Concat(test.b)
		if err := ab.Err; err != nil {
			t.Errorf("Test:%v\nError:%v", testnum, err)
		}
		received := ab.Records()
		expected := test.expected
		if !reflect.DeepEqual(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
		a := test.a
		b := ab
		if err := checkAddr(a.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nA:%v\nAB:%v", testnum, err, a.Addr(), b.Addr())
		}
		a = test.b
		b = ab
		if err := checkAddr(a.Addr(), b.Addr()); err != nil {
			t.Errorf("Test:%v\nError:%v\nB:%v\nAB:%v", testnum, err, a.Addr(), b.Addr())
		}
	}
}

func TestSeries_Order(t *testing.T) {
	tests := []struct {
		series   Series
		reverse  bool
		expected []int
	}{
		{
			Ints([]string{"2", "1", "3", "NaN", "4", "NaN"}),
			false,
			[]int{1, 0, 2, 4, 3, 5},
		},
		{
			Floats([]string{"2", "1", "3", "NaN", "4", "NaN"}),
			false,
			[]int{1, 0, 2, 4, 3, 5},
		},
		{
			Strings([]string{"c", "b", "a"}),
			false,
			[]int{2, 1, 0},
		},
		{
			Bools([]bool{true, false, false, false, true}),
			false,
			[]int{1, 2, 3, 0, 4},
		},
		{
			Ints([]string{"2", "1", "3", "NaN", "4", "NaN"}),
			true,
			[]int{4, 2, 0, 1, 3, 5},
		},
		{
			Floats([]string{"2", "1", "3", "NaN", "4", "NaN"}),
			true,
			[]int{4, 2, 0, 1, 3, 5},
		},
		{
			Strings([]string{"c", "b", "a"}),
			true,
			[]int{0, 1, 2},
		},
		{
			Bools([]bool{true, false, false, false, true}),
			true,
			[]int{0, 4, 1, 2, 3},
		},
	}
	for testnum, test := range tests {
		received := test.series.Order(test.reverse)
		expected := test.expected
		if !reflect.DeepEqual(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
	}
}

func TestSeries_IsNaN(t *testing.T) {
	tests := []struct {
		series   Series
		expected []bool
	}{
		{
			Ints([]string{"2", "1", "3", "NaN", "4", "NaN"}),
			[]bool{false, false, false, true, false, true},
		},
		{
			Floats([]string{"A", "1", "B", "3"}),
			[]bool{true, false, true, false},
		},
	}
	for testnum, test := range tests {
		received := test.series.IsNaN()
		expected := test.expected
		if !reflect.DeepEqual(expected, received) {
			t.Errorf(
				"Test:%v\nExpected:\n%v\nReceived:\n%v",
				testnum, expected, received,
			)
		}
	}
}
