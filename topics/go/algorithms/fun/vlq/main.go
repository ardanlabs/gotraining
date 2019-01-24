package main

import (
	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/vlq/varint"
)

func main() {
	varint.DecodeVarint([]byte{0xFF, 0xFF, 0x7F})
}
