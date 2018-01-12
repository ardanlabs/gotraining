package flaws

import "testing"

func BenchmarkLiteralFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var y1 int
		func(p *int, x int) {
			*p = x
		}(&y1, 42) // BAD: Cause of y escape

		var y2 int
		foo(&y2, 42)
	}
}

func foo(p *int, x int) {
	*p = x
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1
./example1_test.go:17:6: can inline foo as: func(*int, int) { *p = x }
./example1_test.go:5:6: cannot inline BenchmarkLiteralFunctions: unhandled op for
./example1_test.go:13:6: inlining call to foo func(*int, int) { *p = x }
./example1_test.go:8:3: can inline BenchmarkLiteralFunctions.func1 as: func(*int, int) { *p = x }
./example1_test.go:10:5: &y1 escapes to heap
./example1_test.go:10:5: 	from (func literal)(&y1, 42) (parameter to indirect call) at ./example1_test.go:10:4
./example1_test.go:7:7: moved to heap: y1
./example1_test.go:5:40: BenchmarkLiteralFunctions b does not escape
./example1_test.go:8:3: BenchmarkLiteralFunctions func literal does not escape
./example1_test.go:13:7: BenchmarkLiteralFunctions &y2 does not escape
./example1_test.go:8:18: BenchmarkLiteralFunctions.func1 p does not escape
./example1_test.go:17:20: foo p does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1
BenchmarkLiteralFunctions-8   	100000000	        15.7 ns/op	       8 B/op	       1 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 754.51MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1.BenchmarkLiteralFunctions in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example1/example1_test.go
  754.51MB   754.51MB (flat, cum)   100% of Total
         .          .      5:func BenchmarkLiteralFunctions(b *testing.B) {
         .          .      6:	for i := 0; i < b.N; i++ {
  754.51MB   754.51MB      7:		var y1 int
         .          .      8:		func(p *int, x int) {
         .          .      9:			*p = x
         .          .     10:		}(&y1, 42) // BAD: Cause of y escape
         .          .     11:
         .          .     12:		var y2 int
*/
