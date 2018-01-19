package flaws

import "testing"

func BenchmarkAssignmentIndirect(b *testing.B) {
	type X struct {
		p *int
	}
	for i := 0; i < b.N; i++ {
		var i1 int
		x1 := &X{
			p: &i1, // GOOD: i1 does not escape
		}
		_ = x1

		var i2 int
		x2 := &X{}
		x2.p = &i2 // BAD: Cause of i2 escape
	}
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1
./example1_test.go:5:6: cannot inline BenchmarkAssignmentIndirect: unhandled op DCLTYPE
./example1_test.go:18:10: &i2 escapes to heap
./example1_test.go:18:10: 	from x2.p (star-dot-equals) at ./example1_test.go:18:8
./example1_test.go:16:7: moved to heap: i2
./example1_test.go:5:37: BenchmarkAssignmentIndirect b does not escape
./example1_test.go:12:7: BenchmarkAssignmentIndirect &i1 does not escape
./example1_test.go:12:5: BenchmarkAssignmentIndirect &X literal does not escape
./example1_test.go:17:12: BenchmarkAssignmentIndirect &X literal does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1
BenchmarkAssignmentIndirect-8   	100000000	        14.2 ns/op	       8 B/op	       1 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 759.51MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1.BenchmarkAssignmentIndirect in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1/example1_test.go
  759.51MB   759.51MB (flat, cum)   100% of Total
         .          .     11:		x1 := &X{
         .          .     12:			p: &i1, // GOOD: i1 does not escape
         .          .     13:		}
         .          .     14:		_ = x1
         .          .     15:
  759.51MB   759.51MB     16:		var i2 int
         .          .     17:		x2 := &X{}
         .          .     18:		x2.p = &i2 // BAD: Cause of i2 escape
         .          .     19:	}
         .          .     20:}
*/
