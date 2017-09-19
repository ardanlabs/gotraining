// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"testing"
)

func TestAiry(t *testing.T) {
	for _, test := range []struct {
		z, ans complex128
	}{
		// Results computed using Octave.
		{5, 1.08344428136074e-04},
		{5i, 29.9014823980070 + 21.6778315987835i},
	} {
		ans := AiryAi(test.z)
		if math.Abs(real(ans)-real(test.ans)) > 1e-10 {
			t.Errorf("Real part mismatch. Got %v, want %v", real(ans), real(test.ans))
		}
		if math.Abs(imag(ans)-imag(test.ans)) > 1e-10 {
			t.Errorf("Imaginary part mismatch. Got %v, want %v", imag(ans), imag(test.ans))
		}
	}
}
