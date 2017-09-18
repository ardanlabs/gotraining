package mat64_test

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func ExampleSymDense_SubsetSym() {
	n := 5
	s := mat64.NewSymDense(5, nil)
	count := 1.0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			s.SetSym(i, j, count)
			count++
		}
	}
	fmt.Println("Original matrix:")
	fmt.Printf("%0.4v\n\n", mat64.Formatted(s))

	// Take the subset {0, 2, 4}
	var sub mat64.SymDense
	sub.SubsetSym(s, []int{0, 2, 4})
	fmt.Println("Subset {0, 2, 4}")
	fmt.Printf("%0.4v\n\n", mat64.Formatted(&sub))

	// Take the subset {0, 0, 4}
	sub.SubsetSym(s, []int{0, 0, 4})
	fmt.Println("Subset {0, 0, 4}")
	fmt.Printf("%0.4v\n\n", mat64.Formatted(&sub))

	// Output:
	// Original matrix:
	// ⎡ 1   2   3   4   5⎤
	// ⎢ 2   6   7   8   9⎥
	// ⎢ 3   7  10  11  12⎥
	// ⎢ 4   8  11  13  14⎥
	// ⎣ 5   9  12  14  15⎦
	//
	// Subset {0, 2, 4}
	// ⎡ 1   3   5⎤
	// ⎢ 3  10  12⎥
	// ⎣ 5  12  15⎦
	//
	// Subset {0, 0, 4}
	// ⎡ 1   1   5⎤
	// ⎢ 1   1   5⎥
	// ⎣ 5   5  15⎦
}
