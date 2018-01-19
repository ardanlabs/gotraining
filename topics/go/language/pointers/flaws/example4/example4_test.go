package flaws

import "testing"

type Iface interface {
	Method()
}

type X struct {
	name string
}

func (x X) Method() {}

func BenchmarkInterfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x1 := X{"bill"}
		var i1 Iface = x1
		var i2 Iface = &x1

		i1.Method() // BAD: cause copy of x1 to escape
		i2.Method() // BAD: cause x1 to escape

		x2 := X{"bill"}
		foo(x2)
		foo(&x2)
	}
}

func foo(i Iface) {
	i.Method() // BAD: cause value passed in to escape
}

/*
$ go test -gcflags "-m -m" -run none -bench . -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example4
./example4_test.go:13:6: can inline X.Method as: method(X) func() {  }
./example4_test.go:30:6: cannot inline foo: non-leaf op CALLINTER
./example4_test.go:15:6: cannot inline BenchmarkInterfaces: unhandled op for
./example4_test.go:13:9: X.Method x does not escape
./example4_test.go:30:12: leaking param: i
./example4_test.go:30:12: 	from i.Method() (receiver in indirect call) at ./example4_test.go:31:10
./example4_test.go:18:7: x1 escapes to heap
./example4_test.go:18:7: 	from i1 (assigned) at ./example4_test.go:18:7
./example4_test.go:18:7: 	from i1.Method() (receiver in indirect call) at ./example4_test.go:21:12
./example4_test.go:19:7: &x1 escapes to heap
./example4_test.go:19:7: 	from i2 (assigned) at ./example4_test.go:19:7
./example4_test.go:19:7: 	from i2.Method() (receiver in indirect call) at ./example4_test.go:22:12
./example4_test.go:19:18: &x1 escapes to heap
./example4_test.go:19:18: 	from &x1 (interface-converted) at ./example4_test.go:19:7
./example4_test.go:19:18: 	from i2 (assigned) at ./example4_test.go:19:7
./example4_test.go:19:18: 	from i2.Method() (receiver in indirect call) at ./example4_test.go:22:12
./example4_test.go:17:17: moved to heap: x1
./example4_test.go:25:6: x2 escapes to heap
./example4_test.go:25:6: 	from x2 (passed to call[argument escapes]) at ./example4_test.go:25:6
./example4_test.go:26:7: &x2 escapes to heap
./example4_test.go:26:7: 	from &x2 (passed to call[argument escapes]) at ./example4_test.go:26:6
./example4_test.go:26:7: &x2 escapes to heap
./example4_test.go:26:7: 	from &x2 (interface-converted) at ./example4_test.go:26:7
./example4_test.go:26:7: 	from &x2 (passed to call[argument escapes]) at ./example4_test.go:26:6
./example4_test.go:24:17: moved to heap: x2
./example4_test.go:15:45: BenchmarkInterfaces b does not escape

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example4
BenchmarkInterfaces-8   	10000000	       126 ns/op	      64 B/op	       4 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 658.01MB
ROUTINE ======================== github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example4.BenchmarkInterfaces in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example4/example4_test.go
	658.01MB   658.01MB (flat, cum)   100% of Total
				 .          .     12:
				 .          .     13:func (x X) Method() {}
				 .          .     14:
				 .          .     15:func BenchmarkInterfaces(b *testing.B) {
				 .          .     16:	for i := 0; i < b.N; i++ {
	167.50MB   167.50MB     17:		x1 := X{"bill"}
	163.50MB   163.50MB     18:		var i1 Iface = x1
				 .          .     19:		var i2 Iface = &x1
				 .          .     20:
				 .          .     21:		i1.Method() // BAD: cause copy of x to escape
				 .          .     22:		i2.Method() // BAD: cause x to escape
				 .          .     23:
	163.50MB   163.50MB     24:		x2 := X{"bill"}
	163.50MB   163.50MB     25:		foo(x2)
				 .          .     26:		foo(&x2)
				 .          .     27:	}
				 .          .     28:}
*/
