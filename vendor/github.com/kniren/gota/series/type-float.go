package series

import (
	"fmt"
	"math"
	"strconv"
)

type floatElement struct {
	e *float64
}

func (e floatElement) Set(value interface{}) Element {
	var val float64
	switch value.(type) {
	case string:
		if value.(string) == "NaN" {
			e.e = nil
			return e
		}
		f, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			e.e = nil
			return e
		}
		val = f
	case int:
		val = float64(value.(int))
	case float64:
		val = float64(value.(float64))
	case bool:
		b := value.(bool)
		if b {
			val = 1
		} else {
			val = 0
		}
	case Element:
		val = value.(Element).Float()
	default:
		e.e = nil
		return e
	}
	e.e = &val
	return e
}

func (e floatElement) Copy() Element {
	if e.e == nil {
		return floatElement{nil}
	}
	copy := float64(*e.e)
	return floatElement{&copy}
}

func (e floatElement) IsNA() bool {
	if e.e == nil || math.IsNaN(*e.e) {
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
	return float64(*e.e)
}

func (e floatElement) String() string {
	if e.e == nil {
		return "NaN"
	}
	return fmt.Sprintf("%f", *e.e)
}

func (e floatElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	f := *e.e
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
	return float64(*e.e)
}

func (e floatElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch *e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Float \"%v\" to bool", *e.e)
}

func (e floatElement) Addr() string {
	return fmt.Sprint(e.e)
}

func (e floatElement) Eq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e == f
}

func (e floatElement) Neq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e != f
}

func (e floatElement) Less(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e < f
}

func (e floatElement) LessEq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e <= f
}

func (e floatElement) Greater(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e > f
}

func (e floatElement) GreaterEq(elem Element) bool {
	f := elem.Float()
	if e.IsNA() || math.IsNaN(f) {
		return false
	}
	return *e.e >= f
}
