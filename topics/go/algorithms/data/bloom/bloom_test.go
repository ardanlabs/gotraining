package bloom

import (
	"hash/maphash"
	"testing"
	"unsafe"
)

// Hack to generate fixed seed
// See https://github.com/golang/go/issues/43043
func fixedSeed() maphash.Seed {
	var s maphash.Seed
	// s.s is unexported
	ptr := unsafe.Pointer(&s)
	val := (*uint64)(ptr)
	*val = 353

	return s
}

func TestBloom(t *testing.T) {
	b := New(60, 7)
	b.hash.SetSeed(fixedSeed())
	loons := []string{"bugs", "daffy", "elmer", "tweety", "taz", "porky"}
	for _, name := range loons {
		b.Add(name)
	}

	for _, name := range loons {
		if !b.Contains(name) {
			t.Fatalf("%s not found", name)
		}
	}

	for _, name := range []string{"mickey", "pluto"} {
		if b.Contains(name) {
			t.Fatalf("%s found", name)
		}
	}
}
