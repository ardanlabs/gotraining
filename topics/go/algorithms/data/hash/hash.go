// Package hash implements a hash table
package hash

import (
	"fmt"
	"hash/maphash"
)

const (
	numBuckets = 256 // Number of buckets in hash
)

// entry in the hash table
type entry struct {
	key   string
	value int
}

// Hash is a simple Hash table
type Hash struct {
	buckets [][]*entry
	h       maphash.Hash
}

// NewHash returns a new hash table
func NewHash() *Hash {
	return &Hash{
		buckets: make([][]*entry, numBuckets),
	}
}

// Set sets the key to value in the hash
func (h *Hash) Set(key string, value int) {
	b := h.hashKey(key)
	// Look for existing key
	for _, e := range h.buckets[b] {
		if e.key == key {
			e.value = value
			return
		}
	}

	h.buckets[b] = append(h.buckets[b], &entry{key, value})
}

// Get gets the value associated with key, return a error is key not found
func (h *Hash) Get(key string) (int, error) {
	b := h.hashKey(key)
	for _, e := range h.buckets[b] {
		if e.key == key {
			return e.value, nil
		}
	}

	return 0, fmt.Errorf("%q not found", key)
}

// Delete deletes an entry from the hash, return an error if not found
func (h *Hash) Delete(key string) error {
	b := h.hashKey(key)
	bucket := h.buckets[b]
	for i, e := range bucket {
		if e.key == key {
			copy(bucket[i:], bucket[i+1:])
			bucket = bucket[:len(bucket)-1]
			// Free memory when bucket shrinks a lot
			if cap(bucket) > 2*len(bucket) {
				nb := make([]*entry, len(bucket))
				copy(nb, bucket)
				bucket = nb
			}
			h.buckets[b] = bucket
			return nil
		}
	}

	return fmt.Errorf("%q not found", key)
}

// Len return the number of elements in the hash
func (h *Hash) Len() int {
	size := 0
	for _, b := range h.buckets {
		size += len(b)
	}
	return size
}

// Do calls fn on each key/value. If fn return false stops the iteration
func (h *Hash) Do(fn func(key string, value int) bool) {
	for _, b := range h.buckets {
		for _, e := range b {
			ok := fn(e.key, e.value)
			if !ok {
				return
			}
		}
	}
}

// hashKey returns the bucket index for key
func (h *Hash) hashKey(key string) int {
	h.h.Reset()
	h.h.WriteString(key)
	n := h.h.Sum64()
	return int(n % uint64(len(h.buckets)))
}
