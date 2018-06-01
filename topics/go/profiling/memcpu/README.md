## Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Profiling Commands

#### CPU Profiling

Run the benchmark.
   
    $ go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out

Run the pprof tool.
    
    $ go tool pprof cpu.out

Run these pprof commands.

    (pprof) list algOne
    (pprof) web list algOne

_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not (except for goroutines that called runtime.LockOSThread, which is, unfortunately, not visible in the profile). Note that goroutines in "IO wait" state also do not consume threads, they are parked on non-blocking network poller (which uses epoll/kqueue/GetQueuedCompletionStatus to unpark goroutines later)._

#### Memory Profiling

Run the benchmark.

    $ go test -run none -bench . -benchtime 3s -benchmem -memprofile mem.out

Run the pprof tool.

    $ go tool pprof -<PICK_MEM_PROFILE> mem.out

Run these pprof commands.

    (pprof) list algOne
    (pprof) web list algOne

Documentation of memory profile options.

    // Useful to see current status of heap.
	-inuse_space  : Allocations live at the time of profile  	** default
	-inuse_objects: Number of bytes allocated at the time of profile

	// Useful to see pressure on heap over time.
	-alloc_space  : All allocations happened since program start
	-alloc_objects: Number of object allocated at the time of profile

If you want to reduce memory consumption, look at the `-inuse_space` profile collected during normal program operation.
	
If you want to improve execution speed, look at the `-alloc_objects` profile collected after significant running time or at program end.

## Code Review

[Profiling](stream.go) ([Go Playground](https://play.golang.org/p/U7oJS85y5g-)) | 
[Profiling Test](stream_test.go) ([Go Playground](https://play.golang.org/p/n0pvooFGkcl))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
