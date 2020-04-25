/*
Package hash implements a hash table.

Hash table

        hashKey(key) ────────┐
                             │
                             ↓
    ┌────┬─────┬─────┬────┬─────┬─────┬─────┬─────┐
    │    │     │     │    │     │     │     │     │  ←── bucket
    └────┴─────┴─────┴────┴─────┴─────┴─────┴─────┘
             │               │
             ↓               ↓
       ┌─────────────┐  ┌─────────────┐
       │ key │ value │  │ key │ value │  ←── entry
       ├─────────────┤  ├─────────────┤
       │ key │ value │  │ key │ value │
       ├─────────────┤  └─────────────┘
       │ key │ value │
       ├─────────────┤
       │ key │ value │
       ├─────────────┤
       │ key │ value │
       └─────────────┘


- hashKey(key) returns a number between 0 to len(buckets)-1
- We use a slice of entries as a bucket to handles cases where two or more keys
  are hashed to the same bucket
- See more at https://en.wikipedia.org/wiki/Hash_table
*/
package hash
