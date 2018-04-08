// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// PCGSource is an implementation of a 64-bit permuted congruential
// generator as defined in
//
// 	PCG: A Family of Simple Fast Space-Efficient Statistically Good
// 	Algorithms for Random Number Generation
// 	Melissa E. O’Neill, Harvey Mudd College
// 	http://www.pcg-random.org/pdf/toms-oneill-pcg-family-v1.02.pdf
//
// The generator here is the congruential generator PCG XSL RR 128/64 (LCG)
// as found in the software available at http://www.pcg-random.org/.
// It has period 2^128 with 128 bits of state, producing 64-bit values.
// Is state is represented by two uint64 words.
type PCGSource struct {
	low  uint64
	high uint64
}

const (
	maxUint32 = (1 << 32) - 1

	multiplier = 47026247687942121848144207491837523525

	increment = 117397592171526113268558934119004209487
	incHigh   = increment >> 64
	incLow    = increment & maxUint64

	// TODO: Use these?
	initializer = 245720598905631564143578724636268694099
	initHigh    = initializer >> 64
	initLow     = initializer & maxUint64
)

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (pcg *PCGSource) Seed(seed uint64) {
	pcg.low = seed
	pcg.high = seed // TODO: What is right?
}

func (pcg *PCGSource) add() {
	old := pcg.low
	pcg.low += incLow
	if pcg.low < old {
		// Carry occurred.
		pcg.high++
	}
	pcg.high += incHigh
}

func (pcg *PCGSource) multiply() {
	// Break each lower word into two separate 32-bit 'digits' each stored
	// in a 64-bit word with 32 high zero bits.  This allows the overflow
	// into the high word to be computed.
	s0 := (pcg.low >> 00) & maxUint32
	s1 := (pcg.low >> 32) & maxUint32

	const (
		m0    = (multiplier >> 00) & maxUint32
		m1    = (multiplier >> 32) & maxUint32
		mLow  = multiplier & (1<<64 - 1)
		mHigh = multiplier >> 64 & (1<<64 - 1)
	)

	high := pcg.low*mHigh + pcg.high*mLow
	s0m0 := s0 * m0
	s0m1 := s0 * m1
	s1m0 := s1 * m0
	s1m1 := s1 * m1
	high += (s0m1 >> 32) + (s1m0 >> 32)
	carry := (s0m1 & maxUint32) + (s1m0 & maxUint32) + s0m0>>32
	high += (carry >> 32)

	pcg.low *= mLow
	pcg.high = high + s1m1
}
