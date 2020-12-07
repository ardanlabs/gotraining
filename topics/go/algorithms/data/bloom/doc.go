/* Package bloom contains a simple implementation of a bloom filer

A Bloom filter is a space efficient probabilistic data structure.
It supports two operations - adding a memeber and checking for membership.
Asking for membership will return either "not in set" or "probably in set".

Bloom filters are using when using a map will take too much memory.

Algorithm:

A bloom filter has a bit array of "m" bits and "k" < "m" different hash functions.
Each hash function maps a key (set element) into a location in the bit array, this means each element in the Bloom filter is represented by "k" bits.

The filter starts with all "m" bits set to 0.
When adding a key, the "k" hash function generate "k" bits and these bits are set to 1 in the array.
When checking for membership, we check that all "k" bits from the hash function are set to 1 in the underlying array.
See more at https://en.wikipedia.org/wiki/Bloom_filter

Implementation

This implementation uses a random number generators with a fixed seed to generate the "k" bits for each set member. We use a "hash/maphash" to get a fixed seed for each key (a string) and then get "k" bit numbers from the random number generator after seeding it with the key.
*/
package bloom
