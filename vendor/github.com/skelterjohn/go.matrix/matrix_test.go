package matrix

import (
	"testing"
)

func TestParse(t *testing.T) {
	s := `[1 2 3;4 5 6]`
	A, err := ParseMatlab(s)
	
	if err != nil {
		t.Fatal(err)
	}
	
	Ar := MakeDenseMatrix([]float64{1,2,3,4,5,6}, 2, 3)
	if !Equals(A, Ar) {
		t.Error()
	}
}
