// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package bce shows a sample function that does not take into consideration
// the extra bounds checks that the compiler places in code for integrity. By
// using information provided by the ssa backend of the compiler, we can
// change the code to remove the checks and improve performance.
package bce

import "encoding/binary"

// go build -gcflags -d=ssa/check_bce/debug=1

func hash64(buffer []byte, seed uint64) uint64 {
	const (
		k0 = 0xD6D018F5
		k1 = 0xA2AA033B
		k2 = 0x62992FC1
		k3 = 0x30BC5B29
	)

	ptr := buffer

	hash := (seed + k2) * k0

	if len(ptr) >= 32 {
		v := [4]uint64{hash, hash, hash, hash}

		for len(ptr) >= 32 {
			v[0] += binary.LittleEndian.Uint64(ptr) * k0
			ptr = ptr[8:]
			v[0] = rotateRight(v[0], 29) + v[2]
			v[1] += binary.LittleEndian.Uint64(ptr) * k1
			ptr = ptr[8:]
			v[1] = rotateRight(v[1], 29) + v[3]
			v[2] += binary.LittleEndian.Uint64(ptr) * k2
			ptr = ptr[8:]
			v[2] = rotateRight(v[2], 29) + v[0]
			v[3] += binary.LittleEndian.Uint64(ptr) * k3
			ptr = ptr[8:]
			v[3] = rotateRight(v[3], 29) + v[1]
		}
	}

	return hash
}

func rotateRight(v uint64, k uint) uint64 {
	return (v >> k) | (v << (64 - k))
}
