package bloom

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	bs := NewBitSet(19)
	bitNum := 13
	if bs.Get(bitNum) != 0 {
		t.Fatalf("%d not 0", bitNum)
	}
	bs.Set(bitNum, 1)
	if bs.Get(bitNum) != 1 {
		t.Fatalf("%d not 1", bitNum)
	}
	bs.Set(bitNum, 0)
	if bs.Get(bitNum) != 0 {
		t.Fatalf("%d not 0", bitNum)
	}
}
