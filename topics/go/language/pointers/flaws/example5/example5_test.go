package flaws

import (
	"bytes"
	"testing"
)

func BenchmarkUnknown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.Write([]byte{1})
		_ = buf.Bytes()
	}
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example5
./example5_test.go:8:6: cannot inline Benchmark: unhandled op for
./example5_test.go:12:16: inlining call to bytes.(*Buffer).Bytes method(*bytes.Buffer) func() []byte { return bytes.b·2.buf[bytes.b·2.off:] }
./example5_test.go:11:6: buf escapes to heap
./example5_test.go:11:6: 	from buf (passed to call[argument escapes]) at ./example5_test.go:11:12
./example5_test.go:10:7: moved to heap: buf
./example5_test.go:8:19: Benchmark b does not escape
./example5_test.go:11:19: Benchmark []byte literal does not escape
./example5_test.go:12:16: Benchmark buf does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example5
Benchmark-8   	20000000	        50.8 ns/op	     112 B/op	       1 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 2.19GB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example5.BenchmarkUnknown in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example5/example5_test.go
	2.19GB     2.19GB (flat, cum)   100% of Total
		 .          .      8:func BenchmarkUnknown(b *testing.B) {
		 .          .      9:	for i := 0; i < b.N; i++ {
	2.19GB     2.19GB     10:		var buf bytes.Buffer
		 .          .     11:		buf.Write([]byte{1})
		 .          .     12:		_ = buf.Bytes()
		 .          .     13:	}
		 .          .     14:}
*/
