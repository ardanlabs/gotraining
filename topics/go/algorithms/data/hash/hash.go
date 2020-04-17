// Package hash implements a hash table
package hash

import (
	"fmt"
	"hash/maphash"
)

const (
	numBuckets = 256 // Number of buckets in hash
)

// An entry where we store key and value in the hash
type entry struct {
	key   string
	value int
}

// Hash is a simple Hash table
type Hash struct {
	buckets [][]*entry
	hash    maphash.Hash
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

	// Iterate over the entries for the specified bucket.
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
			bucket = removeEntry(bucket, i)
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
	h.hash.Reset()
	h.hash.WriteString(key)
	n := h.hash.Sum64()

	// Return a value in [0:(len(buckets)-1]
	return int(n % uint64(len(h.buckets)))
}

// Remve an entry from a bucket
func removeEntry(bucket []*entry, i int) []*entry {
	copy(bucket[i:], bucket[i+1:])
	bucket = bucket[:len(bucket)-1]

	if cap(bucket) < 2*len(bucket) {
		return bucket
	}

	// Free memory when the bucket shrinks a lot. If we don't do that,
	// the underlying bucket array will stay in memory and will be in
	// the biggest size the bucket ever was
	newBucket := make([]*entry, len(bucket))
	copy(newBucket, bucket)
	return newBucket
}
