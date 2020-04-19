package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	h := NewHash()
	k1, v1 := "key1", 1
	k2, v2 := "key2", 2
	h.Store(k1, v1)
	h.Store(k2, v2)

	if h.Len() != 2 {
		t.Fatalf("len, expected 2, got %d", h.Len())
	}

	v, err := h.Retrieve(k1)
	if err != nil {
		t.Fatal(err)
	}

	if v != v1 {
		t.Fatalf("Get(%q), expected %d, got %d", k1, v1, v)
	}

	v1a := 11
	h.Store(k1, v1a)

	v, err = h.Retrieve(k1)
	if err != nil {
		t.Fatal(err)
	}

	if v != v1a {
		t.Fatalf("Get(%q) after change, expected %d, got %d", k1, v1a, v)
	}

	err = h.Delete(k1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = h.Retrieve(k1)
	if err == nil {
		t.Fatalf("found %q after delete", k1)
	}

	k3 := "key3"
	_, err = h.Retrieve(k3)
	if err == nil {
		t.Fatalf("found non existing key")
	}

	count := 0
	fn := func(key string, value int) bool {
		count++
		return true
	}
	h.Do(fn)
	if count != h.Len() {
		t.Fatalf("Do ran %d times, expected %d", count, h.Len())
	}
}
