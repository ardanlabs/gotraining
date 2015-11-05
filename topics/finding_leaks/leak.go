// I am experimenting with this code. There is no way to identify where in the
// code a leak is occuring. It is a treasure hunt at the end of the day.
//
// pprof
// -----------------------------------------------------------------------------
// We can use the pprof tooling but all it will show are places where lots of
// allocations are taking place.
// 		go tool pprof http://localhost:6060/debug/pprof/heap
//
// Take a snapshot of the current heap profile:
// 		curl -s http://localhost:6060/debug/pprof/heap > base.heap
//
// After some time, take another snapshot:
// 		curl -s http://localhost:6060/debug/pprof/heap > current.heap
//
// Now compare both snapshots against the binary and get into the pprof tool:
// 		go tool pprof -inuse_objects -base base.heap \
//                              /PATH_TO_BINARY/finding_leaks \
//                              current.heap
//
// Run the top command to see the functions allocating the most objects:
// (pprof) top
// 3182575 of 3182575 total (  100%)
// Dropped 5 nodes (cum <= 15912)
//       flat  flat%   sum%        cum   cum%
//    3182575   100%   100%    3182575   100%  main.main.func1
//          0     0%   100%    3182575   100%  runtime.goexit
//
// Run the list command against the goroutine declared in main:
// (pprof) list main.main.func1
// Total: 3182575
// ROUTINE ======================== main.main.func1 in /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/finding_leaks/leak.go
//    3182575    3182575 (flat, cum)   100% of Total
//          .          .     10:func main() {
//          .          .     11:	go func() {
//          .          .     12:		m := make(map[int]int)
//          .          .     13:
//          .          .     14:		for i := 0; ; i++ {
//    3182575    3182575     15:			m[i] = i
//          .          .     16:		}
//          .          .     17:	}()
//          .          .     18:
//          .          .     19:	go func() {
//          .          .     20:		http.ListenAndServe(":6060", nil)
//
// GODEBUG=gctrace=1
// -----------------------------------------------------------------------------
// gctrace: setting gctrace=1 causes the garbage collector to emit a single line
// to standard error at each collection, summarizing the amount of memory collected
// and the length of the pause. Setting gctrace=2 emits the same summary but also
// repeats each collection. The format of this line is subject to change.
// Currently, it is:
// 		gc # @#s #%: #+...+# ms clock, #+...+# ms cpu, #->#-># MB, # MB goal, # P
// where the fields are as follows:
// 		gc #        the GC number, incremented at each GC
// 		@#s         time in seconds since program start
// 		#%          percentage of time spent in GC since program start
// 		#+...+#     wall-clock/CPU times for the phases of the GC
// 		#->#-># MB  heap size at GC start, at GC end, and live heap
// 		# MB goal   goal heap size
// 		# P         number of processors used
// The phases are stop-the-world (STW) sweep termination, scan,
// synchronize Ps, mark, and STW mark termination. The CPU times
// for mark are broken down in to assist time (GC performed in
// line with allocation), background GC time, and idle GC time.
// If the line ends with "(forced)", this GC was forced by a
// runtime.GC() call and all phases are STW.
//
// In C++, a memory leak is memory you have lost a reference to.
// In Go, a memory leak is memory you retain a reference to.
//
// export GODEBUG=gctrace=1
// ./finding_leak
//
// gc 1 @0.009s 1%: 0.059+0.17+0.005+0.24+0.12 ms clock, 0.17+0.17+0+0/0.36/0.067+0.38 ms cpu, 5->5->3 MB, 4 MB goal, 8 P
// gc 2 @0.017s 1%: 0.037+0.096+0.098+0.21+0.086 ms clock, 0.22+0.096+0+0.10/0.31/0.091+0.51 ms cpu, 8->8->7 MB, 7 MB goal, 8 P
// gc 3 @0.032s 1%: 0.020+0.16+0.007+0.25+0.090 ms clock, 0.14+0.16+0+0/0.20/0.27+0.63 ms cpu, 17->17->14 MB, 14 MB goal, 8 P
// gc 4 @0.066s 0%: 0.029+0.17+0.074+0.48+0.10 ms clock, 0.20+0.17+0+0/0.42/0.26+0.76 ms cpu, 35->35->29 MB, 29 MB goal, 8 P
//
package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	// Create a goroutine that leaks memory. Since the map needs
	// to be constantly shuffled around, this becomes very expensive.
	go func() {
		m := make(map[int]int)

		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	// Start a listener for the pprof support.
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// On a <ctrl> C shutdown the program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
}
