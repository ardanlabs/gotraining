/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.

	package vlq

	// DecodeVarint takes a variable length VLQ based integer and
	// decodes it into a 32 bit integer.
	func DecodeVarint(input []byte) (uint32, error)

	// EncodeVarint takes a 32 bit integer and encodes it into
	// a variable length VLQ based integer.
	func EncodeVarint(n uint32) []byte
*/

package vlq

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeVarint(t *testing.T) {
	testCases := []struct {
		input  []byte
		output uint32
	}{
		0:  {[]byte{0x7F}, 127},
		1:  {[]byte{0x81, 0x00}, 128},
		2:  {[]byte{0xC0, 0x00}, 8192},
		3:  {[]byte{0xFF, 0x7F}, 16383},
		4:  {[]byte{0x81, 0x80, 0x00}, 16384},
		5:  {[]byte{0xFF, 0xFF, 0x7F}, 2097151},
		6:  {[]byte{0x81, 0x80, 0x80, 0x00}, 2097152},
		7:  {[]byte{0xC0, 0x80, 0x80, 0x00}, 134217728},
		8:  {[]byte{0xFF, 0xFF, 0xFF, 0x7F}, 268435455},
		9:  {[]byte{0x82, 0x00}, 256},
		10: {[]byte{0x81, 0x10}, 144},
	}

	for i, tc := range testCases {
		t.Logf("test case %d - %#v\n", i, tc.input)
		if o, _ := DecodeVarint(tc.input); o != tc.output {
			t.Fatalf("expected %d\ngot\n%d\n", tc.output, o)
		}
		if encoded := EncodeVarint(tc.output); bytes.Compare(encoded, tc.input) != 0 {
			t.Fatalf("%d - expected %#v\ngot\n%#v\n", tc.output, tc.input, encoded)
		}
	}
}
