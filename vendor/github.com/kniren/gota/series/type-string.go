package series

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type stringElement struct {
	e *string
}

func (e stringElement) Set(value interface{}) Element {
	var val string
	switch value.(type) {
	case string:
		val = string(value.(string))
		if val == "NaN" {
			e.e = nil
			return e
		}
	case int:
		val = strconv.Itoa(value.(int))
	case float64:
		val = strconv.FormatFloat(value.(float64), 'f', 6, 64)
	case bool:
		b := value.(bool)
		if b {
			val = "true"
		} else {
			val = "false"
		}
	case Element:
		val = value.(Element).String()
	default:
		e.e = nil
		return e
	}
	e.e = &val
	return e
}

func (e stringElement) Copy() Element {
	if e.e == nil {
		return stringElement{nil}
	}
	copy := string(*e.e)
	return stringElement{&copy}
}

func (e stringElement) IsNA() bool {
	if e.e == nil {
		return true
	}
	return false
}

func (e stringElement) Type() Type {
	return String
}

func (e stringElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return string(*e.e)
}

func (e stringElement) String() string {
	if e.e == nil {
		return "NaN"
	}
	return string(*e.e)
}

func (e stringElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return strconv.Atoi(*e.e)
}

func (e stringElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	f, err := strconv.ParseFloat(*e.e, 64)
	if err != nil {
		return math.NaN()
	}
	return f
}

func (e stringElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch strings.ToLower(*e.e) {
	case "true", "t", "1":
		return true, nil
	case "false", "f", "0":
		return false, nil
	}
	return false, fmt.Errorf("can't convert String \"%v\" to bool", *e.e)
}

func (e stringElement) Addr() string {
	return fmt.Sprint(e.e)
}

func (e stringElement) Eq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e == elem.String()
}

func (e stringElement) Neq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e != elem.String()
}

func (e stringElement) Less(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e < elem.String()
}

func (e stringElement) LessEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e <= elem.String()
}

func (e stringElement) Greater(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e > elem.String()
}

func (e stringElement) GreaterEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return *e.e >= elem.String()
}
