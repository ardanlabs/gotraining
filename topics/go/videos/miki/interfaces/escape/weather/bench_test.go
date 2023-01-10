package weather

import (
	"testing"
)

const (
	testV         = 32.1
	testC Celsius = testV
)

type Value interface {
	Unit() string
	Amount() float64
}

func valueOf(c Celsius) float64 {
	return c.Amount()
}

func valueOfIface(v Value) float64 {
	return v.Amount()
}

func BenchmarkConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := valueOf(testC)
		if v != testV {
			b.Fatal(v)
		}
	}
}

func BenchmarkIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := valueOfIface(testC)
		if v != testV {
			b.Fatal(v)
		}
	}
}
