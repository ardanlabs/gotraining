package series

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Series is a data structure designed for operating on arrays of elements that
// should comply with a certain type structure. They are flexible enough that can
// be transformed to other Series types and account for missing or non valid
// elements. Most of the power of Series resides on the ability to compare and
// subset Series of different types.
type Series struct {
	Name     string    // The name of the series
	elements []Element // The values of the elements
	t        Type      // The type of the series
	Err      error     // If there are errors they are stored here
}

// Element is the interface that defines the types of methods to be present for
// elements of a Series
type Element interface {
	Set(interface{}) Element
	Copy() Element
	IsNA() bool
	Type() Type
	Val() ElementValue
	String() string
	Int() (int, error)
	Float() float64
	Bool() (bool, error)
	Addr() string
	Eq(Element) bool
	Neq(Element) bool
	Less(Element) bool
	LessEq(Element) bool
	Greater(Element) bool
	GreaterEq(Element) bool
}

// ElementValue represents the value that can be used for marshaling or
// unmarshaling Elements.
type ElementValue interface{}

// Comparator is a convenience alias that can be used for a more type safe way of
// reason and use comparators.
type Comparator string

// Supported Comparators
const (
	Eq        Comparator = "==" // Equal
	Neq                  = "!=" // Non equal
	Greater              = ">"  // Greater than
	GreaterEq            = ">=" // Greater or equal than
	Less                 = "<"  // Lesser than
	LessEq               = "<=" // Lesser or equal than
	In                   = "in" // Inside
)

// Type is a convenience alias that can be used for a more type safe way of
// reason and use Series types.
type Type string

// Supported Series Types
const (
	String Type = "string"
	Int         = "int"
	Float       = "float"
	Bool        = "bool"
)

// Indexes represent the elements that can be used for selecting a subset of
// elements within a Series. Currently supported are:
//
//     int            // Matches the given index number
//     []int          // Matches all given index numbers
//     []bool         // Matches all elements in a Series marked as true
//     Series [Int]   // Same as []int
//     Series [Bool]  // Same as []bool
type Indexes interface{}

// TODO:  New series values as an Alias (type Values interface{}) for better documentation

// New is the generic Series constructor
func New(values interface{}, t Type, name string) Series {
	var elements []Element
	ret := Series{
		Name:     name,
		elements: elements,
		t:        t,
	}

	// Pre-allocate elements
	preAlloc := func(n int) {
		ret.elements = make([]Element, n)
		for i := 0; i < n; i++ {
			switch t {
			case String:
				ret.elements[i] = stringElement{}
			case Int:
				ret.elements[i] = intElement{}
			case Float:
				ret.elements[i] = floatElement{}
			case Bool:
				ret.elements[i] = boolElement{}
			default:
				panic(fmt.Sprintf("unknown type %v", t))
			}
		}
	}
	if values == nil {
		preAlloc(1)
		return ret
	}
	switch values.(type) {
	case []string:
		v := values.([]string)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements[i] = ret.elements[i].Set(v[i])
		}
	case []float64:
		v := values.([]float64)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements[i] = ret.elements[i].Set(v[i])
		}
	case []int:
		v := values.([]int)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements[i] = ret.elements[i].Set(v[i])
		}
	case []bool:
		v := values.([]bool)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements[i] = ret.elements[i].Set(v[i])
		}
	case Series:
		v := values.(Series)
		l := v.Len()
		preAlloc(l)
		for i, e := range v.elements {
			ret.elements[i] = ret.elements[i].Set(e)
		}
	default:
		switch reflect.TypeOf(values).Kind() {
		case reflect.Slice:
			v := reflect.ValueOf(values)
			l := v.Len()
			preAlloc(v.Len())
			for i := 0; i < l; i++ {
				val := v.Index(i).Interface()
				ret.elements[i] = ret.elements[i].Set(val)
			}
		default:
			preAlloc(1)
			v := reflect.ValueOf(values)
			val := v.Interface()
			ret.elements[0] = ret.elements[0].Set(val)
		}
	}
	return ret
}

// Strings is a constructor for a String Series
func Strings(values interface{}) Series {
	return New(values, String, "")
}

// Ints is a constructor for an Int Series
func Ints(values interface{}) Series {
	return New(values, Int, "")
}

// Floats is a constructor for a Float Series
func Floats(values interface{}) Series {
	return New(values, Float, "")
}

// Bools is a constructor for a Bool Series
func Bools(values interface{}) Series {
	return New(values, Bool, "")
}

// Empty returns an empty Series of the same type
func (s Series) Empty() Series {
	var elements []Element
	return Series{
		Name:     s.Name,
		t:        s.t,
		elements: elements,
	}
}

// Append adds new elements to the end of the Series. When using Append, the
// Series is modified in place.
func (s *Series) Append(values interface{}) {
	if err := s.Err; err != nil {
		return
	}
	news := New(values, s.t, s.Name)
	s.elements = append(s.elements, news.elements...)
}

// Concat concatenates two series together. It will return a new Series with the
// combined elements of both Series.
func (s Series) Concat(x Series) Series {
	if err := s.Err; err != nil {
		return s
	}
	if err := x.Err; err != nil {
		s.Err = fmt.Errorf("concat error: argument has errors: %v", err)
		return s
	}
	y := s.Copy()
	y.Append(x)
	return y
}

// Subset returns a subset of the series based on the given Indexes.
func (s Series) Subset(indexes Indexes) Series {
	if err := s.Err; err != nil {
		return s
	}
	idx, err := parseIndexes(s.Len(), indexes)
	if err != nil {
		s.Err = err
		return s
	}
	elements := make([]Element, len(idx))
	for k, i := range idx {
		if i < 0 || i >= s.Len() {
			s.Err = fmt.Errorf("subsetting error: index out of range")
			return s
		}
		elements[k] = s.elements[i].Copy()
	}
	return Series{
		Name:     s.Name,
		t:        s.t,
		elements: elements,
	}
}

// Set sets the values on the indexes of a Series and returns the reference
// for itself. The original Series is modified.
func (s Series) Set(indexes Indexes, newvalues Series) Series {
	if err := s.Err; err != nil {
		return s
	}
	if err := newvalues.Err; err != nil {
		s.Err = fmt.Errorf("set error: argument has errors: %v", err)
		return s
	}
	idx, err := parseIndexes(s.Len(), indexes)
	if err != nil {
		s.Err = err
		return s
	}
	if len(idx) != newvalues.Len() {
		s.Err = fmt.Errorf("set error: dimensions mismatch")
		return s
	}
	for k, i := range idx {
		if i < 0 || i >= s.Len() {
			s.Err = fmt.Errorf("set error: index out of range")
			return s
		}
		s.elements[i] = s.elements[i].Set(newvalues.elements[k])
	}
	return s
}

// HasNaN checks whether the Series contain NaN elements.
func (s Series) HasNaN() bool {
	for _, e := range s.elements {
		if e.IsNA() {
			return true
		}
	}
	return false
}

// IsNaN returns an array that identifies which of the elements are NaN.
func (s Series) IsNaN() []bool {
	ret := make([]bool, s.Len())
	for i, e := range s.elements {
		ret[i] = e.IsNA()
	}
	return ret
}

// Compare compares the values of a Series with other elements. To do so, the
// elements with are to be compared are first transformed to a Series of the same
// type as the caller.
func (s Series) Compare(comparator Comparator, comparando interface{}) Series {
	if err := s.Err; err != nil {
		return s
	}
	compareElements := func(a, b Element, c Comparator) (bool, error) {
		var ret bool
		switch c {
		case Eq:
			ret = a.Eq(b)
		case Neq:
			ret = a.Neq(b)
		case Greater:
			ret = a.Greater(b)
		case GreaterEq:
			ret = a.GreaterEq(b)
		case Less:
			ret = a.Less(b)
		case LessEq:
			ret = a.LessEq(b)
		default:
			return false, fmt.Errorf("unknown comparator: %v", c)
		}
		return ret, nil
	}

	comp := New(comparando, s.t, "")
	bools := make([]bool, s.Len())
	// In comparator comparation
	if comparator == In {
		for i, e := range s.elements {
			b := false
			for _, m := range comp.elements {
				c, err := compareElements(e, m, Eq)
				if err != nil {
					s = s.Empty()
					s.Err = err
					return s
				}
				if c {
					b = true
					break
				}
			}
			bools[i] = b
		}
		return Bools(bools)
	}

	// Single element comparison
	if comp.Len() == 1 {
		for i, e := range s.elements {
			c, err := compareElements(e, comp.elements[0], comparator)
			if err != nil {
				s = s.Empty()
				s.Err = err
				return s
			}
			bools[i] = c
		}
		return Bools(bools)
	}

	// Multiple element comparison
	if s.Len() != comp.Len() {
		s := s.Empty()
		s.Err = fmt.Errorf("can't compare: length mismatch")
		return s
	}
	for i, e := range s.elements {
		c, err := compareElements(e, comp.elements[i], comparator)
		if err != nil {
			s = s.Empty()
			s.Err = err
			return s
		}
		bools[i] = c
	}
	return Bools(bools)
}

// Copy will return a copy of the Series.
func (s Series) Copy() Series {
	name := s.Name
	t := s.t
	err := s.Err
	elements := make([]Element, s.Len())
	for i, e := range s.elements {
		elements[i] = e.Copy()
	}
	ret := Series{
		Name:     name,
		t:        t,
		elements: elements,
		Err:      err,
	}
	return ret
}

// Records returns the elements of a Series as a []string
func (s Series) Records() []string {
	ret := make([]string, s.Len())
	for i, e := range s.elements {
		ret[i] = e.String()
	}
	return ret
}

// Float returns the elements of a Series as a []float64. If the elements can not
// be converted to float64 or contains a NaN returns the float representation of
// NaN.
func (s Series) Float() []float64 {
	ret := make([]float64, s.Len())
	for i, e := range s.elements {
		ret[i] = e.Float()
	}
	return ret
}

// Int returns the elements of a Series as a []int or an error if the
// transformation is not possible.
func (s Series) Int() ([]int, error) {
	ret := make([]int, s.Len())
	for i, e := range s.elements {
		val, err := e.Int()
		if err != nil {
			return nil, err
		}
		ret[i] = val
	}
	return ret, nil
}

// Bool returns the elements of a Series as a []bool or an error if the
// transformation is not possible.
func (s Series) Bool() ([]bool, error) {
	ret := make([]bool, s.Len())
	for i, e := range s.elements {
		val, err := e.Bool()
		if err != nil {
			return nil, err
		}
		ret[i] = val
	}
	return ret, nil
}

// Type returns the type of a given series
func (s Series) Type() Type {
	return s.t
}

// Len returns the length of a given Series
func (s Series) Len() int {
	return len(s.elements)
}

// String implements the Stringer interface for Series
func (s Series) String() string {
	return fmt.Sprint(s.elements)
}

// Str prints some extra information about a given series
func (s Series) Str() string {
	var ret []string
	// If name exists print name
	if s.Name != "" {
		ret = append(ret, "Name: "+s.Name)
	}
	ret = append(ret, "Type: "+fmt.Sprint(s.t))
	ret = append(ret, "Length: "+fmt.Sprint(s.Len()))
	if s.Len() != 0 {
		ret = append(ret, "Values: "+fmt.Sprint(s))
	}
	return strings.Join(ret, "\n")
}

// Val returns the value of a series for the given index
func (s Series) Val(i int) (interface{}, error) {
	if i >= s.Len() || i < 0 {
		return nil, fmt.Errorf("index out of bounds")
	}
	return s.elements[i].Val(), nil
}

// Elem returns the element of a series for the given index or nil if the index is
// out of bounds
func (s Series) Elem(i int) Element {
	if i >= s.Len() || i < 0 {
		return nil
	}
	return s.elements[i]
}

// Addr returns the string representation of the memory address that store the
// values of a given Series.
func (s Series) Addr() []string {
	ret := make([]string, s.Len())
	for i, e := range s.elements {
		ret[i] = e.Addr()
	}
	return ret
}

func parseIndexes(l int, indexes Indexes) ([]int, error) {
	var idx []int
	switch indexes.(type) {
	case []int:
		idx = indexes.([]int)
	case int:
		idx = []int{indexes.(int)}
	case []bool:
		bools := indexes.([]bool)
		if len(bools) != l {
			return nil, fmt.Errorf("indexing error: index dimensions mismatch")
		}
		for i, b := range bools {
			if b {
				idx = append(idx, i)
			}
		}
	case Series:
		s := indexes.(Series)
		if err := s.Err; err != nil {
			return nil, fmt.Errorf("indexing error: new values has errors: %v", err)
		}
		if s.HasNaN() {
			return nil, fmt.Errorf("indexing error: indexes contain NaN")
		}
		switch s.t {
		case Int:
			return s.Int()
		case Bool:
			bools, err := s.Bool()
			if err != nil {
				return nil, fmt.Errorf("indexing error: %v", err)
			}
			return parseIndexes(l, bools)
		default:
			return nil, fmt.Errorf("indexing error: unknown indexing mode")
		}
	default:
		return nil, fmt.Errorf("indexing error: unknown indexing mode")
	}
	return idx, nil
}

// Order returns the indexes for sorting a Series. NaN elements are pushed to the
// end by order of appearance.
func (s Series) Order(reverse bool) []int {
	var ie indexedElements
	var nasIdx []int
	for i, e := range s.elements {
		if e.IsNA() {
			nasIdx = append(nasIdx, i)
		} else {
			ie = append(ie, indexedElement{i, e})
		}
	}
	var srt sort.Interface
	srt = ie
	if reverse {
		srt = sort.Reverse(srt)
	}
	sort.Sort(srt)
	var ret []int
	for _, e := range ie {
		ret = append(ret, e.index)
	}
	return append(ret, nasIdx...)
}

type indexedElement struct {
	index   int
	element Element
}

type indexedElements []indexedElement

func (e indexedElements) Len() int           { return len(e) }
func (e indexedElements) Less(i, j int) bool { return e[i].element.Less(e[j].element) }
func (e indexedElements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
