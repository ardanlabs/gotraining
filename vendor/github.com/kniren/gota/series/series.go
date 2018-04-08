package series

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"math"

	"github.com/gonum/stat"
)

// Series is a data structure designed for operating on arrays of elements that
// should comply with a certain type structure. They are flexible enough that can
// be transformed to other Series types and account for missing or non valid
// elements. Most of the power of Series resides on the ability to compare and
// subset Series of different types.
type Series struct {
	Name     string   // The name of the series
	elements Elements // The values of the elements
	t        Type     // The type of the series
	Err      error    // If there are errors they are stored here
}

// Elements is the interface that represents the array of elements contained on
// a Series.
type Elements interface {
	Elem(int) Element
	Len() int
}

// Element is the interface that defines the types of methods to be present for
// elements of a Series
type Element interface {
	// Setter method
	Set(interface{})

	// Comparation methods
	Eq(Element) bool
	Neq(Element) bool
	Less(Element) bool
	LessEq(Element) bool
	Greater(Element) bool
	GreaterEq(Element) bool

	// Accessor/conversion methods
	Copy() Element     // FIXME: Returning interface is a recipe for pain
	Val() ElementValue // FIXME: Returning interface is a recipe for pain
	String() string
	Int() (int, error)
	Float() float64
	Bool() (bool, error)

	// Information methods
	IsNA() bool
	Type() Type
}

// intElements is the concrete implementation of Elements for Int elements.
type intElements []intElement

func (e intElements) Len() int           { return len(e) }
func (e intElements) Elem(i int) Element { return &e[i] }

// stringElements is the concrete implementation of Elements for String elements.
type stringElements []stringElement

func (e stringElements) Len() int           { return len(e) }
func (e stringElements) Elem(i int) Element { return &e[i] }

// floatElements is the concrete implementation of Elements for Float elements.
type floatElements []floatElement

func (e floatElements) Len() int           { return len(e) }
func (e floatElements) Elem(i int) Element { return &e[i] }

// boolElements is the concrete implementation of Elements for Bool elements.
type boolElements []boolElement

func (e boolElements) Len() int           { return len(e) }
func (e boolElements) Elem(i int) Element { return &e[i] }

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

// New is the generic Series constructor
func New(values interface{}, t Type, name string) Series {
	ret := Series{
		Name: name,
		t:    t,
	}

	// Pre-allocate elements
	preAlloc := func(n int) {
		switch t {
		case String:
			ret.elements = make(stringElements, n)
		case Int:
			ret.elements = make(intElements, n)
		case Float:
			ret.elements = make(floatElements, n)
		case Bool:
			ret.elements = make(boolElements, n)
		default:
			panic(fmt.Sprintf("unknown type %v", t))
		}
	}

	if values == nil {
		preAlloc(1)
		ret.elements.Elem(0).Set(nil)
		return ret
	}

	switch values.(type) {
	case []string:
		v := values.([]string)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements.Elem(i).Set(v[i])
		}
	case []float64:
		v := values.([]float64)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements.Elem(i).Set(v[i])
		}
	case []int:
		v := values.([]int)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements.Elem(i).Set(v[i])
		}
	case []bool:
		v := values.([]bool)
		l := len(v)
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements.Elem(i).Set(v[i])
		}
	case Series:
		v := values.(Series)
		l := v.Len()
		preAlloc(l)
		for i := 0; i < l; i++ {
			ret.elements.Elem(i).Set(v.elements.Elem(i))
		}
	default:
		switch reflect.TypeOf(values).Kind() {
		case reflect.Slice:
			v := reflect.ValueOf(values)
			l := v.Len()
			preAlloc(v.Len())
			for i := 0; i < l; i++ {
				val := v.Index(i).Interface()
				ret.elements.Elem(i).Set(val)
			}
		default:
			preAlloc(1)
			v := reflect.ValueOf(values)
			val := v.Interface()
			ret.elements.Elem(0).Set(val)
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
	return New([]int{}, s.t, s.Name)
}

// Append adds new elements to the end of the Series. When using Append, the
// Series is modified in place.
func (s *Series) Append(values interface{}) {
	if err := s.Err; err != nil {
		return
	}
	news := New(values, s.t, s.Name)
	switch s.t {
	case String:
		s.elements = append(s.elements.(stringElements), news.elements.(stringElements)...)
	case Int:
		s.elements = append(s.elements.(intElements), news.elements.(intElements)...)
	case Float:
		s.elements = append(s.elements.(floatElements), news.elements.(floatElements)...)
	case Bool:
		s.elements = append(s.elements.(boolElements), news.elements.(boolElements)...)
	}
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
	ret := Series{
		Name: s.Name,
		t:    s.t,
	}
	switch s.t {
	case String:
		elements := make(stringElements, len(idx))
		for k, i := range idx {
			elements[k] = s.elements.(stringElements)[i]
		}
		ret.elements = elements
	case Int:
		elements := make(intElements, len(idx))
		for k, i := range idx {
			elements[k] = s.elements.(intElements)[i]
		}
		ret.elements = elements
	case Float:
		elements := make(floatElements, len(idx))
		for k, i := range idx {
			elements[k] = s.elements.(floatElements)[i]
		}
		ret.elements = elements
	case Bool:
		elements := make(boolElements, len(idx))
		for k, i := range idx {
			elements[k] = s.elements.(boolElements)[i]
		}
		ret.elements = elements
	default:
		panic("unknown series type")
	}
	return ret
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
		s.elements.Elem(i).Set(newvalues.elements.Elem(k))
	}
	return s
}

// HasNaN checks whether the Series contain NaN elements.
func (s Series) HasNaN() bool {
	for i := 0; i < s.Len(); i++ {
		if s.elements.Elem(i).IsNA() {
			return true
		}
	}
	return false
}

// IsNaN returns an array that identifies which of the elements are NaN.
func (s Series) IsNaN() []bool {
	ret := make([]bool, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.elements.Elem(i).IsNA()
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
		for i := 0; i < s.Len(); i++ {
			e := s.elements.Elem(i)
			b := false
			for j := 0; j < comp.Len(); j++ {
				m := comp.elements.Elem(j)
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
		for i := 0; i < s.Len(); i++ {
			e := s.elements.Elem(i)
			c, err := compareElements(e, comp.elements.Elem(0), comparator)
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
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
		c, err := compareElements(e, comp.elements.Elem(i), comparator)
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
	var elements Elements
	switch s.t {
	case String:
		elements = make(stringElements, s.Len())
		copy(elements.(stringElements), s.elements.(stringElements))
	case Float:
		elements = make(floatElements, s.Len())
		copy(elements.(floatElements), s.elements.(floatElements))
	case Bool:
		elements = make(boolElements, s.Len())
		copy(elements.(boolElements), s.elements.(boolElements))
	case Int:
		elements = make(intElements, s.Len())
		copy(elements.(intElements), s.elements.(intElements))
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
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
		ret[i] = e.String()
	}
	return ret
}

// Float returns the elements of a Series as a []float64. If the elements can not
// be converted to float64 or contains a NaN returns the float representation of
// NaN.
func (s Series) Float() []float64 {
	ret := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
		ret[i] = e.Float()
	}
	return ret
}

// Int returns the elements of a Series as a []int or an error if the
// transformation is not possible.
func (s Series) Int() ([]int, error) {
	ret := make([]int, s.Len())
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
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
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
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
	return s.elements.Len()
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

// Val returns the value of a series for the given index. Will panic if the index
// is out of bounds.
func (s Series) Val(i int) interface{} {
	return s.elements.Elem(i).Val()
}

// Elem returns the element of a series for the given index. Will panic if the
// index is out of bounds.
func (s Series) Elem(i int) Element {
	return s.elements.Elem(i)
}

// parseIndexes will parse the given indexes for a given series of length `l`. No
// out of bounds checks is performed.
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
	for i := 0; i < s.Len(); i++ {
		e := s.elements.Elem(i)
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

// StdDev calculates the standard deviation of a series
func (s Series) StdDev() float64 {
	stdDev := stat.StdDev(s.Float(), nil)
	return stdDev
}

// Mean calculates the average value of a series
func (s Series) Mean() float64 {
	stdDev := stat.Mean(s.Float(), nil)
	return stdDev
}

// Max return the biggest element in the series
func (s Series) Max() float64 {
	if s.elements.Len() == 0 || s.Type() == String {
		return math.NaN()
	}

	max := s.elements.Elem(0)
	for i := 1; i < s.elements.Len(); i++ {
		elem := s.elements.Elem(i)
		if elem.Greater(max) {
			max = elem
		}
	}
	return max.Float()
}

// MaxStr return the biggest element in a series of type String
func (s Series) MaxStr() string {
	if s.elements.Len() == 0 || s.Type() != String {
		return ""
	}

	max := s.elements.Elem(0)
	for i := 1; i < s.elements.Len(); i++ {
		elem := s.elements.Elem(i)
		if elem.Greater(max) {
			max = elem
		}
	}
	return max.String()
}

// Min return the lowest element in the series
func (s Series) Min() float64 {
	if s.elements.Len() == 0 || s.Type() == String {
		return math.NaN()
	}

	min := s.elements.Elem(0)
	for i := 1; i < s.elements.Len(); i++ {
		elem := s.elements.Elem(i)
		if elem.Less(min) {
			min = elem
		}
	}
	return min.Float()
}

// MinStr return the lowest element in a series of type String
func (s Series) MinStr() string {
	if s.elements.Len() == 0 || s.Type() != String {
		return ""
	}

	min := s.elements.Elem(0)
	for i := 1; i < s.elements.Len(); i++ {
		elem := s.elements.Elem(i)
		if elem.Less(min) {
			min = elem
		}
	}
	return min.String()
}

// Quantile returns the sample of x such that x is greater than or
// equal to the fraction p of samples.
// Note: gonum/stat panics when called with strings
func (s Series) Quantile(p float64) float64 {
	if s.Type() == String || s.Len() == 0 {
		return math.NaN()
	}

	ordered := s.Subset(s.Order(false)).Float()

	return stat.Quantile(p, stat.Empirical, ordered, nil)
}
