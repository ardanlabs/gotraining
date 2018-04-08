// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

package rand

import "math/bits"

// Uint64 returns a pseudo-random 64-bit unsigned integer as a uint64.
func (pcg *PCGSource) Uint64() uint64 {
	pcg.multiply()
	pcg.add()
	// XOR high and low 64 bits together and rotate right by high 6 bits of state.
	return bits.RotateLeft64(pcg.high^pcg.low, -int(pcg.high>>58))
}
