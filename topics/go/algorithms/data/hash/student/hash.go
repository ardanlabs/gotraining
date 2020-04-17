// Package hash asks the student to implement a hash table in Go.
package hash

/*
Data diagram

   hashKey(key) ────────┐
                        │
                        ↓
    ┌────┬─────┬─────┬────┬─────┬─────┬─────┬─────┐
    │    │     │     │    │     │     │     │     │  ←── buckets
    └────┴─────┴─────┴────┴─────┴─────┴─────┴─────┘
       │                │
       │                │
       ↓                ↓
   ┌─────────────┐   ┌─────────────┐
   │ Key │ Value │   │ Key │ Value │  ←── entry
   ├─────────────┤   ├─────────────┤
   │ Key │ Value │   │ Key │ Value │
   ├─────────────┤   └─────────────┘
   │ Key │ Value │
   ├─────────────┤
   │ Key │ Value │
   ├─────────────┤
   │ Key │ Value │
   └─────────────┘
*/

// entry in the hash table
type entry struct {
	key   string
	value int
}

// Hash is a simple Hash table
type Hash struct {
	buckets [][]*entry
}

// NewHash returns a new hash table
func NewHash() *Hash {
	return nil
}

// Set sets the key to value in the hash
func (h *Hash) Set(key string, value int) {
}

// Get gets the value associated with key, return a error is key not found
func (h *Hash) Get(key string) (int, error) {
	return 0, nil
}

// Delete deletes an entry from the hash, return an error if not found
func (h *Hash) Delete(key string) error {
	return nil
}

// Len return the number of elements in the hash
func (h *Hash) Len() int {
	return 0
}

// Do calls fn on each key/value. If fn return false stops the iteration
func (h *Hash) Do(fn func(key string, value int) bool) {
}
