package flaws

import (
	"net/http"
	"testing"
)

func BenchmarkHandler(b *testing.B) {

	// Setup route with specific handler.
	h := func(w http.ResponseWriter, r *http.Request) error {
		// fmt.Println("Specific Request Handler")
		return nil
	}
	route := wrapHandler(h)

	// Execute route.
	for i := 0; i < b.N; i++ {
		var r http.Request
		route(nil, &r) // BAD: Cause of r escape
	}
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func wrapHandler(h Handler) Handler {
	f := func(w http.ResponseWriter, r *http.Request) error {
		// fmt.Println("Boilerplate Code")
		return h(w, r)
	}
	return f
}

/*
$ go test -gcflags "-m -m" -run none -bench BenchmarkHandler -benchmem -memprofile mem.out

# github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2
./example2_http_test.go:20:14: &r escapes to heap
./example2_http_test.go:20:14: 	from route(nil, &r) (parameter to indirect call) at ./example2_http_test.go:20:8
./example2_http_test.go:19:7: moved to heap: r

goos: darwin
goarch: amd64
pkg: github.com/ardanlabs/gotraining/topics/go/language/pointers/flaws/example2
BenchmarkHandler-8   	20000000	        72.4 ns/op	     256 B/op	       1 allocs/op

$ go tool pprof -alloc_space mem.out

Type: alloc_space
(pprof) list Benchmark
Total: 5.07GB
ROUTINE ========================
    5.07GB     5.07GB (flat, cum)   100% of Total
         .          .     14:	}
         .          .     15:	route := wrapHandler(h)
         .          .     16:
         .          .     17:	// Execute route.
         .          .     18:	for i := 0; i < b.N; i++ {
    5.07GB     5.07GB     19:		var r http.Request
         .          .     20:		route(nil, &r) // BAD: Cause of r escape
         .          .     21:	}
         .          .     22:}
*/
