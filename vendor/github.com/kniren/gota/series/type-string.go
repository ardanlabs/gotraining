package series

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type stringElement struct {
	e   string
	nan bool
}

func (e *stringElement) Set(value interface{}) {
	e.nan = false
	switch value.(type) {
	case string:
		e.e = string(value.(string))
		if e.e == "NaN" {
			e.nan = true
			return
		}
	case int:
		e.e = strconv.Itoa(value.(int))
	case float64:
		e.e = strconv.FormatFloat(value.(float64), 'f', 6, 64)
	case bool:
		b := value.(bool)
		if b {
			e.e = "true"
		} else {
			e.e = "false"
		}
	case Element:
		e.e = value.(Element).String()
	default:
		e.nan = true
		return
	}
	return
}

func (e stringElement) Copy() Element {
	if e.IsNA() {
		return &stringElement{"", true}
	}
	return &stringElement{e.e, false}
}

func (e stringElement) IsNA() bool {
	if e.nan {
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
	return string(e.e)
}

func (e stringElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return string(e.e)
}

func (e stringElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return strconv.Atoi(e.e)
}

func (e stringElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	f, err := strconv.ParseFloat(e.e, 64)
	if err != nil {
		return math.NaN()
	}
	return f
}

func (e stringElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch strings.ToLower(e.e) {
	case "true", "t", "1":
		return true, nil
	case "false", "f", "0":
		return false, nil
	}
	return false, fmt.Errorf("can't convert String \"%v\" to bool", e.e)
}

func (e stringElement) Eq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e == elem.String()
}

func (e stringElement) Neq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e != elem.String()
}

func (e stringElement) Less(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e < elem.String()
}

func (e stringElement) LessEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e <= elem.String()
}

func (e stringElement) Greater(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e > elem.String()
}

func (e stringElement) GreaterEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	return e.e >= elem.String()
}
