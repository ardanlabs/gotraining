package flaws

import (
	"testing"
)

func BenchmarkLiteralFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var y1 int
		foo(&y1, 42) // GOOD: y1 does not escape

		var y2 int
		func(p *int, x int) {
			*p = x
		}(&y2, 42) // BAD: Cause of y2 escape

		var y3 int
		p := foo
		p(&y3, 42) // BAD: Cause of y3 escape
	}
}

func foo(p *int, x int) {
	*p = x
}

/*
$ go test -gcflags "-m -m" -run none -bench BenchmarkLiteralFunctions -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2
./example2_test.go:21:6: can inline foo as: func(*int, int) { *p = x }
./example2_test.go:5:6: cannot inline BenchmarkLiteralFunctions: unhandled op for
./example2_test.go:8:6: inlining call to foo func(*int, int) { *p = x }
./example2_test.go:11:3: can inline BenchmarkLiteralFunctions.func1 as: func(*int, int) { *p = x }
./example2_test.go:13:5: &y2 escapes to heap
./example2_test.go:13:5: 	from (func literal)(&y2, 42) (parameter to indirect call) at ./example2_test.go:13:4
./example2_test.go:10:7: moved to heap: y2
./example2_test.go:17:5: &y3 escapes to heap
./example2_test.go:17:5: 	from p(&y3, 42) (parameter to indirect call) at ./example2_test.go:17:4
./example2_test.go:15:7: moved to heap: y3
./example2_test.go:5:35: BenchmarkLiteralFunctions b does not escape
./example2_test.go:8:7: BenchmarkLiteralFunctions &y1 does not escape
./example2_test.go:11:3: BenchmarkLiteralFunctions func literal does not escape
./example2_test.go:11:18: BenchmarkLiteralFunctions.func1 p does not escape
./example2_test.go:21:20: foo p does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2
BenchmarkLiteralFunctions-8   	50000000	        30.7 ns/op	      16 B/op	       2 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 768.01MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2.BenchmarkLiteralFunctions in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2/example2_test.go
	768.01MB   768.01MB (flat, cum)   100% of Total
				 .          .      5:func BenchmarkLiteralFunctions(b *testing.B) {
				 .          .      6:	for i := 0; i < b.N; i++ {
				 .          .      7:		var y1 int
				 .          .      8:		foo(&y1, 42) // GOOD: y1 does not escape
				 .          .      9:
	380.51MB   380.51MB     10:		var y2 int
				 .          .     11:		func(p *int, x int) {
				 .          .     12:			*p = x
				 .          .     13:		}(&y2, 42) // BAD: Cause of y2 escape
				 .          .     14:
	387.51MB   387.51MB     15:		var y3 int
				 .          .     16:		p := foo
				 .          .     17:		p(&y3, 42) // BAD: Cause of y3 escape
				 .          .     18:	}
				 .          .     19:}
*/
