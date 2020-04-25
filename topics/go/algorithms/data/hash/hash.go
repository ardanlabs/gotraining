// Package hash implements a hash table.
package hash

import (
	"fmt"
	"hash/maphash"
)

const numBuckets = 256

// An entry where we store key and value in the hash.
type entry struct {
	key   string
	value int
}

// Hash is a simple Hash table implementation.
type Hash struct {
	buckets [][]entry
	hash    maphash.Hash
}

// New returns a new hash table.
func New() *Hash {
	return &Hash{
		buckets: make([][]entry, numBuckets),
	}
}

// Store adds a value in the hash table based on the key.
func (h *Hash) Store(key string, value int) {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	idx := h.hashKey(key)

	// Extract a copy of the bucket from the hash table.
	bucket := h.buckets[idx]

	// Iterate over the indexes for the specified bucket.
	for idx := range bucket {

		// Compare the keys and if there is a match replace the
		// existing entry value for the new value.
		if bucket[idx].key == key {
			bucket[idx].value = value
			return
		}
	}

	// This key does not exist, so add this new value.
	h.buckets[idx] = append(bucket, entry{key, value})
}

// Retrieve extracts a value from the hash table based on the key.
func (h *Hash) Retrieve(key string) (int, error) {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	idx := h.hashKey(key)

	// Iterate over the entries for the specified bucket.
	for _, entry := range h.buckets[idx] {

		// Compare the keys and if there is a match return
		// the value associated with the key.
		if entry.key == key {
			return entry.value, nil
		}
	}

	// The key was not found so return the error.
	return 0, fmt.Errorf("%q not found", key)
}

// Delete deletes an entry from the hash table.
func (h *Hash) Delete(key string) error {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	bucketIdx := h.hashKey(key)

	// Extract a copy of the bucket from the hash table.
	bucket := h.buckets[bucketIdx]

	// Iterate over the entries for the specified bucket.
	for entryIdx, entry := range bucket {

		// Compare the keys and if there is a match remove
		// the entry from the bucket.
		if entry.key == key {

			// Remove the entry based on its index position.
			bucket = removeEntry(bucket, entryIdx)

			// Replace the existing bucket for the new one.
			h.buckets[bucketIdx] = bucket
			return nil
		}
	}

	// The key was not found so return the error.
	return fmt.Errorf("%q not found", key)
}

// Len return the number of elements in the hash. This function currently
// uses a linear traversal but could be improved with meta-data.
func (h *Hash) Len() int {
	sum := 0
	for _, bucket := range h.buckets {
		sum += len(bucket)
	}
	return sum
}

// Do calls fn on each key/value. If fn return false stops the iteration.
func (h *Hash) Do(fn func(key string, value int) bool) {
	for _, bucket := range h.buckets {
		for _, entry := range bucket {
			if ok := fn(entry.key, entry.value); !ok {
				return
			}
		}
	}
}

// hashKey calculates the bucket index position to use
// for the specified key.
func (h *Hash) hashKey(key string) int {

	// Reset the maphash to initial state so we'll get the same
	// hash value for the same key.
	h.hash.Reset()

	// Write the key to the maphash to update the current state.
	// We don't check error value since WriteString never fails.
	h.hash.WriteString(key)

	// Ask the maphash for its current state which we will
	// use to calculate the final bucket index.
	n := h.hash.Sum64()

	// Use the modulu operator to return a value in the range
	// of our bucket length defined by the const numBuckets.
	return int(n % numBuckets)
}

// removeEntry performs the physical act of removing an
// entry from a bucket,
func removeEntry(bucket []entry, idx int) []entry {

	// https://github.com/golang/go/wiki/SliceTricks
	// Cut out the entry by taking all entries from
	// infront of the index and moving them behind the
	// index specified.
	copy(bucket[idx:], bucket[idx+1:])

	// Set the proper length for the new slice since
	// an entry was removed. The length needs to be
	// reduced by 1.
	bucket = bucket[:len(bucket)-1]

	// Look to see if the current allocation for the
	// bucket can be reduced due to the amount of
	// entries removed from this bucket.
	return reduceAllocation(bucket)
}

// reduceAllocation looks to see if memory can be freed to
// when a bucket has lost a percent of entries.
func reduceAllocation(bucket []entry) []entry {

	// If the bucket if more than ½ full, do nothing.
	if cap(bucket) < 2*len(bucket) {
		return bucket
	}

	// Free memory when the bucket shrinks a lot. If we don't do that,
	// the underlying bucket array will stay in memory and will be in
	// the biggest size the bucket ever was
	newBucket := make([]entry, len(bucket))
	copy(newBucket, bucket)
	return newBucket
}
