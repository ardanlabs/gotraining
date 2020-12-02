package bloom

import (
	"testing"
)

func TestBloom(t *testing.T) {
	b := New(60, 7)
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
