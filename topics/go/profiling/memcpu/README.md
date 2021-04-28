## Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Profiling Commands

#### CPU Profiling

Run the benchmark.
   
    $ go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out

Run the pprof tool using the command line tooling.
    
    $ go tool pprof p.out

Run these pprof commands.

    (pprof) list algOne
    (pprof) web list algOne

Run the pprof tool using the browser based tooling.

    $ go tool pprof -http :3000 p.out

Navigate the drop down menu in the UI.

Run the pprof tool using including the `memcpu.test` binary.

    $ go tool pprof -http :3000 memcpu.test p.out

When you do this, you can get profiling information down to the assembly level.

#### Memory Profiling

Run the benchmark.

    $ go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out

Run the pprof tool.

    $ go tool pprof -<OPTIONAL_PICK_MEM_PROFILE> p.out

Run these pprof commands.

    (pprof) list algOne
    (pprof) web list algOne

Run the pprof tool using the browser based tooling.

    $ go tool pprof -<OPTIONAL_PICK_MEM_PROFILE> -http :3000 p.out

Navigate the drop down menu in the UI.

Run the pprof tool using including the `memcpu.test` binary.

    $ go tool pprof -<OPTIONAL_PICK_MEM_PROFILE> -http :3000 memcpu.test p.out

When you do this, you can get profiling information down to the assembly level.

Documentation of memory profile options.

    // Useful to see pressure on heap over time.
	-alloc_space  : All allocations happened since program start ** default
	-alloc_objects: Number of object allocated at the time of profile

    // Useful to see current status of heap.
	-inuse_space  : Allocations live at the time of profile
	-inuse_objects: Number of bytes allocated at the time of profile

If you want to reduce memory consumption, look at the `-inuse_space` profile collected during normal program operation.

If you want to improve execution speed, look at the `-alloc_objects` profile collected after significant running time or at program end.

## Code Review

[Profiling](stream.go) ([Go Playground](https://play.golang.org/p/Cm92cvurEnE)) | 
[Profiling Test](stream_test.go) ([Go Playground](https://play.golang.org/p/9xzn4zOeviO))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
