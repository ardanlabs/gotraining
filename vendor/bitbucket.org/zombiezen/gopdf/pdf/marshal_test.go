// Copyright (C) 2011, Ross Light

package pdf

import (
	"testing"
)

type marshalTest struct {
	Value    interface{}
	Expected string
}

type fooStruct struct {
	Size    int64
	Params  map[name]string
	NotHere string  `pdf:"-"`
	Rename  float64 `pdf:"Pi"`
	TheVoid []int   `pdf:",omitempty"`
}

var marshalTests = []marshalTest{
	{nil, "null"},
	{"", "()"},
	{"This is a string", "(This is a string)"},
	{"Strings may contain newlines\nand such.", "(Strings may contain newlines\nand such.)"},
	{"Escape (this).", `(Escape \(this\).)`},
	{int(123), "123"},
	{int(-321), "-321"},
	{float64(-3.141599), "-3.14160"},
	{float64(1e9), "1000000000.00000"},
	{name(""), "/"},
	{name("foo"), "/foo"},
	{[]interface{}{}, `[ ]`},
	{[]string{"foo", "(parens)"}, `[ (foo) (\(parens\)) ]`},
	{map[name]string{}, `<< >>`},
	{map[name]string{name("foo"): "bar"}, `<< /foo (bar) >>`},
	{indirectObject{Reference{42, 0}, "foo"}, "42 0 obj\r\n(foo)\r\nendobj"},
	{Reference{42, 0}, `42 0 R`},
	{
		fooStruct{
			Size:    42,
			Params:  map[name]string{name("this"): "that"},
			NotHere: "XXX",
			Rename:  3.141592,
			TheVoid: []int{1, 2, 3},
		},
		`<< /Size 42 /Params << /this (that) >> /Pi 3.14159 /TheVoid [ 1 2 3 ] >>`,
	},
	{
		fooStruct{
			Size:    42,
			Params:  map[name]string{name("this"): "that"},
			NotHere: "XXX",
			Rename:  3.141592,
			TheVoid: nil,
		},
		`<< /Size 42 /Params << /this (that) >> /Pi 3.14159 >>`,
	},
	{
		Rectangle{Point{1, 2}, Point{3, 4}},
		`[ 1.00000 2.00000 3.00000 4.00000 ]`,
	},
}

func TestMarshal(t *testing.T) {
	for i, tt := range marshalTests {
		result, err := marshal(nil, tt.Value)
		switch {
		case err != nil:
			t.Errorf("%d. Marshal(%#v) error: %v", i, tt.Value, err)
		case string(result) != tt.Expected:
			t.Errorf("%d. Marshal(%#v) != %q (got %q)", i, tt.Value, tt.Expected, result)
		}
	}
}
