package flaws

import "testing"

func BenchmarkSliceMapAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]*int)
		var x1 int
		m[0] = &x1 // BAD: cause of x1 escape

		s := make([]*int, 1)
		var x2 int
		s[0] = &x2 // BAD: cause of x2 escape
	}
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3
./example3_test.go:5:6: cannot inline BenchmarkSliceMapAssignment: unhandled op for
./example3_test.go:9:10: &x1 escapes to heap
./example3_test.go:9:10: 	from m[0] (value of map put) at ./example3_test.go:9:8
./example3_test.go:8:7: moved to heap: x1
./example3_test.go:13:10: &x2 escapes to heap
./example3_test.go:13:10: 	from s[0] (slice-element-equals) at ./example3_test.go:13:8
./example3_test.go:12:7: moved to heap: x2
./example3_test.go:5:37: BenchmarkSliceMapAssignment b does not escape
./example3_test.go:7:12: BenchmarkSliceMapAssignment make(map[int]*int) does not escape
./example3_test.go:11:12: BenchmarkSliceMapAssignment make([]*int, 1) does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3
BenchmarkSliceMapAssignment-8   	10000000	       104 ns/op	      16 B/op	       2 allocs/op

$ go tool pprof -alloc_space mem.out
Type: alloc_space
(pprof) list Benchmark
Total: 162.50MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3.BenchmarkSliceMapAssignment in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3/example3_test.go
	162.50MB   162.50MB (flat, cum)   100% of Total
				 .          .      5:func BenchmarkSliceMapAssignment(b *testing.B) {
				 .          .      6:	for i := 0; i < b.N; i++ {
				 .          .      7:		m := make(map[int]*int)
	107.50MB   107.50MB      8:		var x1 int
				 .          .      9:		m[0] = &x1 // BAD: cause of x1 escape
				 .          .     10:
				 .          .     11:		s := make([]*int, 1)
			55MB       55MB     12:		var x2 int
				 .          .     13:		s[0] = &x2 // BAD: cause of x2 escape
				 .          .     14:	}
				 .          .     15:}
*/
