package grpcutil

import (
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// This buffer size is:
		// 1. Reasonably smaller than the max gRPC size
		// 2. Small enough that having hundreds of these buffers won't
		// overwhelm the node
		// 3. Large enough for message-sending to be efficient
		return make([]byte, MaxMsgSize/10)
	},
}

// GetBuffer returns a buffer.  The buffer may or may not be freshly
// allocated, and it may or may not be zero-ed.
func GetBuffer() []byte {
	return bufPool.Get().([]byte)
}

// PutBuffer returns the buffer to the pool.
func PutBuffer(buf []byte) {
	bufPool.Put(buf)
}
