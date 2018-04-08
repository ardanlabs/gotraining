package series

import (
	"fmt"
	"math"
	"strconv"
)

type floatElement struct {
	e   float64
	nan bool
}

func (e *floatElement) Set(value interface{}) {
	e.nan = false
	switch value.(type) {
	case string:
		if value.(string) == "NaN" {
			e.nan = true
			return
		}
		f, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			e.nan = true
			return
		}
		e.e = f
	case int:
		e.e = float64(value.(int))
	case float64:
		e.e = float64(value.(float64))
	case bool:
		b := value.(bool)
		if b {
			e.e = 1
		} else {
			e.e = 0
		}
	case Element:
		e.e = value.(Element).Float()
	default:
		e.nan = true
		return
	}
	return
}

func (e floatElement) Copy() Element {
	if e.IsNA() {
		return &floatElement{0.0, true}
	}
	return &floatElement{e.e, false}
}

func (e floatElement) IsNA() bool {
	if e.nan || math.IsNaN(e.e) {
		return true
	}
	return false
}

func (e floatElement) Type() Type {
	return Float
}

func (e floatElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return float64(e.e)
}

func (e floatElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprintf("%f", e.e)
}

func (e floatElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	f := e.e
	if math.IsInf(f, 1) || math.IsInf(f, -1) {
		return 0, fmt.Errorf("can't convert Inf to int")
	}
	if math.IsNaN(f) {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(f), nil
}

func (e floatElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e)
}

func (e floatElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Float \"%v\" to bool", e.e)
}

func (e floatElement) Eq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e == f
}

func (e floatElement) Neq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e != f
}

func (e floatElement) Less(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e < f
}

func (e floatElement) LessEq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e <= f
}

func (e floatElement) Greater(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e > f
}

func (e floatElement) GreaterEq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return e.e >= f
}
