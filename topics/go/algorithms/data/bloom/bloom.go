package bloom

import (
	"hash/maphash"
	"math/rand"
)

type Bloom struct {
	hash   maphash.Hash
	r      *rand.Rand
	size   int
	probes int
	mask   *BitSet
}

func New(size, probes int) *Bloom {
	b := Bloom{
		r:      rand.New(rand.NewSource(0)),
		size:   size,
		probes: probes,
		mask:   NewBitSet(size),
	}
	return &b
}

func (b *Bloom) seedFor(key string) int64 {
	b.hash.Reset()
	b.hash.WriteString(key)
	return int64(b.hash.Sum64())
}

func (b *Bloom) sample(seed int64) []int {
	b.r.Seed(seed)
	samples := make([]int, 0, b.probes)
	seen := make(map[int]bool)
	for len(samples) < b.probes {
		i := b.r.Intn(b.size)
		if seen[i] {
			continue
		}
		seen[i] = true
		samples = append(samples, i)
	}

	return samples
}

func (b *Bloom) bitsFor(key string) []int {
	seed := b.seedFor(key)
	bits := b.sample(seed)
	return bits
}

func (b *Bloom) Add(key string) {
	for _, bitNum := range b.bitsFor(key) {
		b.mask.Set(bitNum, 1)
	}
}

func (b *Bloom) Contains(key string) bool {
	for _, bitNum := range b.bitsFor(key) {
		if b.mask.Get(bitNum) == 0 {
			return false
		}
	}
	return true
}
