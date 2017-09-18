// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright ©2011-2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette

import (
	"image/color"
	"testing"
)

func withinEpsilon(a, b, epsilon uint32) bool {
	d := int64(a) - int64(b)
	if d < 0 {
		d = -d
	}
	return d <= int64(epsilon)
}

func TestColor(t *testing.T) {
	for r := 0; r < 256; r += 5 {
		for g := 0; g < 256; g += 5 {
			for b := 0; b < 256; b += 5 {
				col := color.RGBA{uint8(r), uint8(g), uint8(b), 0}
				cDirectR, cDirectG, cDirectB, cDirectA := col.RGBA()
				hsva := rgbaToHsva(col.RGBA())
				if hsva.H < 0 || hsva.H > 1 {
					t.Errorf("unexpected values for H: want [0, 1] got:%f", hsva.H)
				}
				if hsva.S < 0 || hsva.S > 1 {
					t.Errorf("unexpected values for S: want [0, 1] got:%f", hsva.S)
				}
				if hsva.V < 0 || hsva.V > 1 {
					t.Errorf("unexpected values for V: want [0, 1] got:%f", hsva.V)
				}

				cFromHSVR, cFromHSVG, cFromHSVB, cFromHSVA := hsva.RGBA()
				if cFromHSVR < 0 || cFromHSVR > 0xFFFF {
					t.Errorf("unexpected values for H: want [0x0, 0xFFFF] got:%f", hsva.H)
				}
				if cFromHSVG < 0 || cFromHSVG > 0xFFFF {
					t.Errorf("unexpected values for S: want [0x0, 0xFFFF] got:%f", hsva.S)
				}
				if cFromHSVB < 0 || cFromHSVB > 0xFFFF {
					t.Errorf("unexpected values for V: want [0x0, 0xFFFF] got:%f", hsva.V)
				}
				if cFromHSVA < 0 || cFromHSVA > 0xFFFF {
					t.Errorf("unexpected values for V: want [0x0, 0xFFFF] got:%f", hsva.V)
				}

				back := rgbaToHsva(color.RGBA{uint8(cFromHSVR >> 8), uint8(cFromHSVG >> 8), uint8(cFromHSVB >> 8), uint8(cFromHSVA)}.RGBA())
				if hsva != back {
					t.Errorf("roundtrip error: got:%#v want:%#v", back, hsva)
				}
				const epsilon = 1
				if !withinEpsilon(cFromHSVR, cDirectR, epsilon) {
					t.Errorf("roundtrip error for R: got:%d want:%d", cFromHSVR, cDirectR)
				}
				if !withinEpsilon(cFromHSVG, cDirectG, epsilon) {
					t.Errorf("roundtrip error for G: got:%d want:%d %d", cFromHSVG, cDirectG, cFromHSVG-cDirectG)
				}
				if !withinEpsilon(cFromHSVB, cDirectB, epsilon) {
					t.Errorf("roundtrip error for B: got:%d want:%d", cFromHSVB, cDirectB)
				}
				if cFromHSVA != cDirectA {
					t.Errorf("roundtrip error for A: got:%d want:%d", cFromHSVA, cDirectA)
				}
				if t.Failed() {
					return
				}
			}
		}
	}
}
