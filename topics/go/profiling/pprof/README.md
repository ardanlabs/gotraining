## http/pprof Profiling

Using the http/pprof support you can profile your web applications and services to see exactly where your performance or memory is being taken.

### Building and Running the Project

We have a website that we will use to learn and explore more about profiling. This project is just creating a struct type value, marshaling it to JSON and using that as the response.

	go build
	./pprof

	http://localhost:4000/sendjson

### Adding Load

To add load to the service while running profiling we can run these command.

	// Send 1M request using 8 connections.
	hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"

### Raw http/pprof

We already added the following import so we can include the profiling route to our web service.

	import _ "net/http/pprof"

Look at the basic profiling stats from the new endpoint.

	http://localhost:4000/debug/pprof

Run a single search from the Browser and then refresh the profile information.

	http://localhost:4000/sendjson

Put some load of the web application. Review the raw profiling information once again.

	hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"

### Looking at Heap Profiles

	http://localhost:4000/debug/pprof/heap?debug=1

At the top of the heap profile you will see this information.
```
heap profile: 0: 0 [9318: 3545424] @ heap/1048576
[0:] 		    Currently live objects,
[0] 		    Amount of memory occupied by live objects
[9318:] 	    Total number of allocations
[3545424] 	    Amount of memory occupied by all allocations
```

At the bottom of the heap profile you will see this information.
```
// General statistics.
Alloc      uint64 // bytes allocated and not yet freed
TotalAlloc uint64 // bytes allocated (even if freed)
Sys        uint64 // bytes obtained from system (sum of XxxSys below)
Lookups    uint64 // number of pointer lookups
Mallocs    uint64 // number of mallocs
Frees      uint64 // number of frees

// Main allocation heap statistics.
HeapAlloc    uint64 // bytes allocated and not yet freed (same as Alloc above)
HeapSys      uint64 // bytes obtained from system
HeapIdle     uint64 // bytes in idle spans
HeapInuse    uint64 // bytes in non-idle span
HeapReleased uint64 // bytes released to the OS
HeapObjects  uint64 // total number of allocated objects

// Low-level fixed-size structure allocator statistics.
//	Inuse is bytes used now.
//	Sys is bytes obtained from system.
StackInuse  uint64 // bytes used by stack allocator
StackSys    uint64
MSpanInuse  uint64 // mspan structures
MSpanSys    uint64
MCacheInuse uint64 // mcache structures
MCacheSys   uint64
BuckHashSys uint64 // profiling bucket hash table
GCSys       uint64 // GC metadata
OtherSys    uint64 // other system allocations

// Garbage collector statistics.
NextGC        uint64 // next collection will happen when HeapAlloc ≥ this amount
PauseNs       [256]uint64 // circular buffer of recent GC pause durations, most recent at [(NumGC+255)%256]
NumGC         uint32 // Number of GC that have run
```

### Interactive Profiling

#### Heap Profiling

Run the Go pprof tool in another window or tab to review heap information.

	go tool pprof -<PICK_MEM_PROFILE> ./pprof http://localhost:4000/debug/pprof/heap

	// Useful to see current status of heap.
	-inuse_space  : Allocations live at the time of profile  	** default
	-inuse_objects: Number of bytes allocated at the time of profile

	// Useful to see pressure on heap over time.
	-alloc_space  : All allocations happened since program start
	-alloc_objects: Number of object allocated at the time of profile

	If you want to reduce memory consumption, look at the `-inuse_space` profile collected during
	normal program operation.
	
	If you want to improve execution speed, look at the `-alloc_objects` profile collected after
	significant running time or at program end.

#### CPU Profiling

Run the Go pprof tool in another window or tab to review cpu information.

	go tool pprof ./pprof http://localhost:4000/debug/pprof/profile

	_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not
	(except for goroutines that called runtime.LockOSThread, which is, unfortunately, not
	visible in the profile). Note that goroutines in "IO wait" state also do not consume
	threads, they are parked on non-blocking network poller
	(which uses epoll/kqueue/GetQueuedCompletionStatus to unpark goroutines later)._

Explore using the **top**, **list**, **web** and **web list** commands.

### Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.

    curl -s http://localhost:4000/debug/pprof/heap > base.heap

After some time, take another snapshot:

    curl -s http://localhost:4000/debug/pprof/heap > current.heap

Now compare both snapshots against the binary and get into the pprof tool:

    go tool pprof -inuse_space -base base.heap ./pprof current.heap
