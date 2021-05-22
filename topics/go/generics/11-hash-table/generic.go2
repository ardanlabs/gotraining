// This code is provided by Matt Layher (@mdlayher)
// https://mdlayher.com/blog/go-generics-draft-design-building-a-hashtable/#a-generic-hashtable

package main

import (
	"fmt"
	"hash/fnv"
)

// =============================================================================

type hashFunc[K comparable] func(key K, buckets int) int

type keyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

type Table[K comparable, V any] struct {
	hashFunc hashFunc[K]
	buckets  int
	data     [][]keyValuePair[K, V]
}

func New[K comparable, V any](buckets int, hf hashFunc[K]) *Table[K, V] {
	return &Table[K, V]{
		hashFunc: hf,
		buckets:  buckets,
		data:     make([][]keyValuePair[K, V], buckets),
	}
}

func (t *Table[K, V]) Insert(key K, value V) {
	bucket := t.hashFunc(key, t.buckets)

	for idx, kvp := range t.data[bucket] {
		if key == kvp.Key {
			t.data[bucket][idx].Value = value
			return
		}
	}

	kvp := keyValuePair[K, V]{
		Key:   key,
		Value: value,
	}
	t.data[bucket] = append(t.data[bucket], kvp)
}

func (t *Table[K, V]) Retrieve(key K) (V, bool) {
	bucket := t.hashFunc(key, t.buckets)

	for idx, kvp := range t.data[bucket] {
		if key == kvp.Key {
			return t.data[bucket][idx].Value, true
		}
	}

	var zero V
	return zero, false
}

// =============================================================================

func main() {
	const buckets = 8

	hashFunc1 := func(key string, buckets int) int {
		h := fnv.New32()
		h.Write([]byte(key))
		return int(h.Sum32()) % buckets
	}
	table1 := New[/*key*/ string, /*value*/ int](buckets, hashFunc1)

	hashFunc2 := func(key int, buckets int) int {
		return key % buckets
	}
	table2 := New[/*key*/ int, /*value*/ string](buckets, hashFunc2)

	words := []string{"foo", "bar", "baz"}
	for i, word := range words {
		table1.Insert(word, i)
		table2.Insert(i, word)
	}

	for i, s := range append(words, "nope!") {
		v1, ok1 := table1.Retrieve(s)
		fmt.Printf("t1.Rtr(%v) = (%v, %v)\n", s, v1, ok1)

		v2, ok2 := table2.Retrieve(i)
		fmt.Printf("t2.Rtr(%v) = (%v, %v)\n", i, v2, ok2)
	}
}
