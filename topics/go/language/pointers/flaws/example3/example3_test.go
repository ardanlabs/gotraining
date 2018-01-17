package flaws

import "testing"

func BenchmarkSliceMapAssignment(b *testing.B) {
	type X struct {
		name string
	}
	for i := 0; i < b.N; i++ {
		m := make(map[int]*X)
		var x1 X
		m[0] = &x1 // BAD: cause of x1 escape

		s := make([]*X, 1)
		var x2 X
		s[0] = &x2 // BAD: cause of x2 escape
	}
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3
./example3_test.go:5:6: cannot inline BenchmarkSliceMapAssignment: unhandled op DCLTYPE
./example3_test.go:12:10: &x1 escapes to heap
./example3_test.go:12:10: 	from m[0] (value of map put) at ./example3_test.go:12:8
./example3_test.go:11:11: moved to heap: x1
./example3_test.go:16:10: &x2 escapes to heap
./example3_test.go:16:10: 	from s[0] (slice-element-equals) at ./example3_test.go:16:8
./example3_test.go:15:11: moved to heap: x2
./example3_test.go:5:39: BenchmarkSliceMapAssignment b does not escape
./example3_test.go:10:12: BenchmarkSliceMapAssignment make(map[int]*X) does not escape
./example3_test.go:14:12: BenchmarkSliceMapAssignment make([]*X, 1) does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3
BenchmarkSliceMapAssignment-8   	10000000	       131 ns/op	      32 B/op	       2 allocs/op

$ go tool pprof -alloc_space mem.out
Type: alloc_space
(pprof) list Benchmark
Total: 345.01MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3.BenchmarkSliceMapAssignment in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example3/example3_test.go
  345.01MB   345.01MB (flat, cum)   100% of Total
         .          .      6:	type X struct {
         .          .      7:		name string
         .          .      8:	}
         .          .      9:	for i := 0; i < b.N; i++ {
         .          .     10:		m := make(map[int]*X)
     175MB      175MB     11:		var x1 X
         .          .     12:		m[0] = &x1 // BAD: cause of x1 escape
         .          .     13:
         .          .     14:		s := make([]*X, 1)
     170MB      170MB     15:		var x2 X
         .          .     16:		s[0] = &x2 // BAD: cause of x2 escape
         .          .     17:	}
         .          .     18:}
*/
