package bloom

import (
	"hash/maphash"
	"math/big"
	"math/rand"
)

// Bloom is an implementation of a bloom filter.
type Bloom struct {
	m    int      // Number of bits
	k    int      // Number of probes
	mask *big.Int // Underlying bitmask

	hash maphash.Hash // Used to generate seed for key
	r    *rand.Rand   // Used to generate "k" bit numbers
}

// New creates a new Bloom filter with "m" bits and "k" has functions.
func New(m, k int) *Bloom {
	b := Bloom{
		m:    m,
		k:    k,
		r:    rand.New(rand.NewSource(0)),
		mask: big.NewInt(0),
	}
	return &b
}

// seedFor return a seed for key.
// For a given key, the returned seed will be the same.
func (b *Bloom) seedFor(key string) int64 {
	b.hash.Reset()
	b.hash.WriteString(key)
	return int64(b.hash.Sum64())
}

// sample returns sample of b.k out of b.m for seed.
// For a given seed, we'll get back the same bits.
func (b *Bloom) sample(seed int64) []int {
	b.r.Seed(seed)
	samples := make([]int, 0, b.k)
	seen := make(map[int]bool)
	for len(samples) < b.k {
		i := b.r.Intn(b.m)
		if seen[i] {
			continue
		}
		seen[i] = true
		samples = append(samples, i)
	}

	return samples
}

// bitsFor return b.k bits for key.
func (b *Bloom) bitsFor(key string) []int {
	seed := b.seedFor(key)
	bits := b.sample(seed)
	return bits
}

// Add adds a key to the bloom filter.
func (b *Bloom) Add(key string) {
	for _, bitNum := range b.bitsFor(key) {
		b.mask.SetBit(b.mask, bitNum, 1)
	}
}

// Contains return true is key is probably in b.
func (b *Bloom) Contains(key string) bool {
	for _, bitNum := range b.bitsFor(key) {
		if b.mask.Bit(bitNum) == 0 {
			return false
		}
	}
	return true
}
