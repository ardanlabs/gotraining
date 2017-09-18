package base

import (
	"math"
	"unsafe"
)

// PackU64ToBytesInline fills ret with the byte values of
// val. Ret must have length at least 8.
func PackU64ToBytesInline(val uint64, ret []byte) {
	ret[7] = byte(val & (0xFF << 56) >> 56)
	ret[6] = byte(val & (0xFF << 48) >> 48)
	ret[5] = byte(val & (0xFF << 40) >> 40)
	ret[4] = byte(val & (0xFF << 32) >> 32)
	ret[3] = byte(val & (0xFF << 24) >> 24)
	ret[2] = byte(val & (0xFF << 16) >> 16)
	ret[1] = byte(val & (0xFF << 8) >> 8)
	ret[0] = byte(val & (0xFF << 0) >> 0)
}

// PackFloatToBytesInline fills ret with the byte values of
// the float64 argument. ret must be at least 8 bytes in size.
func PackFloatToBytesInline(val float64, ret []byte) {
	PackU64ToBytesInline(math.Float64bits(val), ret)
}

// PackU64ToBytes allocates a return value of appropriate length
// and fills it with the values of val.
func PackU64ToBytes(val uint64) []byte {
	ret := make([]byte, 8)
	ret[7] = byte(val & (0xFF << 56) >> 56)
	ret[6] = byte(val & (0xFF << 48) >> 48)
	ret[5] = byte(val & (0xFF << 40) >> 40)
	ret[4] = byte(val & (0xFF << 32) >> 32)
	ret[3] = byte(val & (0xFF << 24) >> 24)
	ret[2] = byte(val & (0xFF << 16) >> 16)
	ret[1] = byte(val & (0xFF << 8) >> 8)
	ret[0] = byte(val & (0xFF << 0) >> 0)
	return ret
}

// UnpackBytesToU64 converst a given byte slice into
// a uint64 value.
func UnpackBytesToU64(val []byte) uint64 {
	pb := unsafe.Pointer(&val[0])
	return *(*uint64)(pb)
}

// PackFloatToBytes returns a 8-byte slice containing
// the byte values of a float64.
func PackFloatToBytes(val float64) []byte {
	return PackU64ToBytes(math.Float64bits(val))
}

// UnpackBytesToFloat converts a given byte slice into an
// equivalent float64.
func UnpackBytesToFloat(val []byte) float64 {
	pb := unsafe.Pointer(&val[0])
	return *(*float64)(pb)
}

func byteSeqEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
