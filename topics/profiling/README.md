## Profiling Code

We can use the go tooling to inspect and profile our programs. Profiling is more of a journey and detective work. It requires some understanding about the application and expectations. The profiling data in and of itself is just raw numbers. We have to give it meaning and understanding.

## The Basics of Profiling

### How does a profiler work?

A profiler runs your program and configures the operating system to interrupt it at regular intervals. This is done by sending SIGPROF to the program being profiled, which suspends and transfers execution to the profiler. The profiler then grabs the program counter for each executing thread and restarts the program.

### Profiling do's and don't's

Before you profile, you must have a stable environment to get repeatable results.

* The machine must be idle—don't profile on shared hardware, don't browse the web while waiting for a long benchmark to run.
* Watch out for power saving and thermal scaling.
* Avoid virtual machines and shared cloud hosting; they are too noisy for consistent measurements.
* There is a kernel bug on OS X versions less than El Capitan; upgrade or avoid profiling on OS X.

If you can afford it, buy dedicated performance test hardware. Rack them, disable all the power management and thermal scaling and never update the software on those machines.

For everyone else, have a before and after sample and run them multiple times to get consistent results.

### Types of Profiling

**CPU profiling**  
CPU profiling is the most common type of profile. When CPU profiling is enabled, the runtime will interrupt itself every 10ms and record the stack trace of the currently running goroutines. Once the profile is saved to disk, we can analyse it to determine the hottest code paths. The more times a function appears in the profile, the more time that code path is taking as a percentage of the total runtime.

**Memory profiling**  
Memory profiling records the stack trace when a heap allocation is made. Memory profiling, like CPU profiling is sample based. By default memory profiling samples 1 in every 1000 allocations. This rate can be changed. Stack allocations are assumed to be free and are not tracked in the memory profile. Because of memory profiling is sample based and because it tracks allocations not use, using memory profiling to determine your application's overall memory usage is difficult.

**Block profiling**  
Block profiling is quite unique. A block profile is similar to a CPU profile, but it records the amount of time a goroutine spent waiting for a shared resource. This can be useful for determining concurrency bottlenecks in your application. Block profiling can show you when a large number of goroutines could make progress, but were blocked. 

Blocking includes:

* Sending or receiving on a unbuffered channel.
* Sending to a full channel, receiving from an empty one.
* Trying to Lock a sync.Mutex that is locked by another goroutine.
* Block profiling is a very specialised tool, it should not be used until you believe you have eliminated all your CPU and memory usage bottlenecks.

**One profile at at time**  
Profiling is not free. Profiling has a moderate, but measurable impact on program performance—especially if you increase the memory profile sample rate. Most tools will not stop you from enabling multiple profiles at once. If you enable multiple profiles at the same time, they will observe their own interactions and skew your results.

**Do not enable more than one kind of profile at a time.**

### Hints to interpret what you see in the profile

If you see lots of time spent in `runtime.mallocgc` function, the program potentially makes excessive amount of small memory allocations. The profile will tell you where the allocations are coming from. See the memory profiler section for suggestions on how to optimize this case.

If lots of time is spent in channel operations, `sync.Mutex` code and other synchronization primitives or System component, the program probably suffers from contention. Consider to restructure program to eliminate frequently accessed shared resources. Common techniques for this include sharding/partitioning, local buffering/batching and copy-on-write technique.

If lots of time is spent in `syscall.Read/Write`, the program potentially makes excessive amount of small reads and writes. Bufio wrappers around os.File or net.Conn can help in this case.

If lots of time is spent in GC component, the program either allocates too many transient objects or heap size is very small so garbage collections happen too frequently.

* Large objects affect memory consumption and GC time, while large number of tiny allocations affects execution speed.

* Combine values into larger values. This will reduce number of memory allocations (faster) and also reduce pressure on garbage collector (faster garbage collections).

* Values that do not contain any pointers are not scanned by garbage collector. Removing pointers from actively used value can positively impact garbage collection time.

## Installing Tools

**boom**  
boom is a modern HTTP benchmarking tool capable of generating the load you need to run tests. It's built using the Go language and leverages goroutines for behind the scenes async IO and concurrency.

	go get -u github.com/rakyll/boom

## Dave Cheney's Profiling Presentation:  

Much of what I have learned comes from Dave and working on solving problems. This slide deck is a great place to start. Much of this material can be found in the material below.

http://go-talks.appspot.com/github.com/davecheney/presentations/seven.slide#1

## Profiling, Debugging and Optimization Reading

Here is more reading and videos to also help get you started.

[Profiling Go Programs](http://golang.org/blog/profiling-go-programs) - Go Team  
[Profiling & Optimizing in Go](https://www.youtube.com/watch?v=xxDZuPEgbBU) - Brad Fitzpatrick  
[Go Dynamic Tools](https://www.youtube.com/watch?v=a9xrxRsIbSU) - Dmitry Vyukov  
[How NOT to Measure Latency](https://www.youtube.com/watch?v=lJ8ydIuPFeU&feature=youtu.be) - Gil Tene  
[Go Performance Tales](http://jmoiron.net/blog/go-performance-tales) - Jason Moiron  
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov  
[Reduce allocation in Go code](https://methane.github.io/2015/02/reduce-allocation-in-go-code) - Python Bytes  
[Write High Performance Go](http://go-talks.appspot.com/github.com/davecheney/presentations/writing-high-performance-go.slide) - Dave Cheney  
[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team   
[Seven ways to profile a Go program](https://www.bigmarker.com/remote-meetup-go/Seven-ways-to-profile-a-Go-program) - Dave Cheney   
[runtime: goroutine execution stalled during GC](https://github.com/golang/go/issues/16293) - Caleb Spare  

## Go and OS Tooling

### time

The **time** command provide information that can help you get a sense how your program is performing.

Use the **time** command to see data about building the program.
	cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/profiling/project
	/usr/bin/time -lp go build		-- Mac OS X
	/usr/bin/time -v go build		-- Linux

### perf

If you're a linux user, then perf(1) is a great tool for profiling applications. Now we have frame pointers, perf can profile Go applications.

	% go build -toolexec="perf stat" cmd/compile/internal/gc
	# cmd/compile/internal/gc

 	Performance counter stats for '/home/dfc/go/pkg/tool/linux_amd64/compile -o $WORK/cmd/compile/internal/gc.a -trimpath $WORK -p cmd/compile/internal/gc -complete -buildid 87cd803267511b4d9e753d68b5b66a70e2f878c4 -D _/home/dfc/go/src/cmd/compile/internal/gc -I $WORK -pack ./alg.go ./align.go ./bexport.go ./bimport.go ./builtin.go ./bv.go ./cgen.go ./closure.go ./const.go ./cplx.go ./dcl.go ./esc.go ./export.go ./fmt.go ./gen.go ./go.go ./gsubr.go ./init.go ./inl.go ./lex.go ./magic.go ./main.go ./mpfloat.go ./mpint.go ./obj.go ./opnames.go ./order.go ./parser.go ./pgen.go ./plive.go ./popt.go ./racewalk.go ./range.go ./reflect.go ./reg.go ./select.go ./sinit.go ./sparselocatephifunctions.go ./ssa.go ./subr.go ./swt.go ./syntax.go ./type.go ./typecheck.go ./universe.go ./unsafe.go ./util.go ./walk.go':

       7026.140760 task-clock (msec)         #    1.283 CPUs utilized          
             1,665 context-switches          #    0.237 K/sec                  
                39 cpu-migrations            #    0.006 K/sec                  
            77,362 page-faults               #    0.011 M/sec                  
    21,769,537,949 cycles                    #    3.098 GHz                     [83.41%]
    11,671,235,864 stalled-cycles-frontend   #   53.61% frontend cycles idle    [83.31%]
     6,839,727,058 stalled-cycles-backend    #   31.42% backend  cycles idle    [66.65%]
    27,157,950,447 instructions              #    1.25  insns per cycle        
                                             #    0.43  stalled cycles per insn [83.25%]
     5,351,057,260 branches                  #  761.593 M/sec                   [83.49%]
       118,150,150 branch-misses             #    2.21% of all branches         [83.15%]

       5.476816754 seconds time elapsed

## Basic Go Profiling

Learn the basics of using GODEBUG.  
[Memory Tracing](godebug/gctrace) | [Scheduler Tracing](godebug/schedtrace)

Learn the basics of using memory and cpu profiling.  
[Memory and CPU Profiling](memcpu)

Learn the basics of blocking profiling.  
[Blocking Profiling](blocking)

Learn the basics of trace profiling.  
[Trace Profiling](trace)

## Profiling a Web Service

We have a web application that extends a web service. Let's profile this application and attempt to understand how it is working.

### Building and Running the Project

We have a website that we will use to learn and explore more about profiling. This project is a search engine for RSS feeds. Run the website and validate it is working.

	go build
	./project

	http://localhost:5000/search

### Adding Load

To add load to the service while running profiling we can run these command.

	// Send 100k request using 8 connections.
	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

### GODEBUG

#### GC Trace

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=gctrace=1 ./project > /dev/null

Put some load of the web application.

	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

#### Scheduler Trace

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=schedtrace=1000 GOMAXPROCS=2 ./project > /dev/null

Put some load of the web application.

	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

### PPROF

pprof descends from the Google Performance Tools suite. pprof profiling is built into the Go runtime and comes in two parts:

* runtime/pprof package built into every Go program
* go tool pprof for investigating profiles.

#### Raw http/pprof

We already added the following import so we can include the profiling route to our web service.

	import _ "net/http/pprof"

Look at the basic profiling stats from the new endpoint.

	http://localhost:5000/debug/pprof

Run a single search from the Browser and then refresh the profile information.

	http://localhost:5000/search?term=house&cnn=on

Put some load of the web application. Review the raw profiling information once again.

	boom -m POST -c 8 -n 10000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Look at the heap profile

	http://localhost:5000/debug/pprof/heap?debug=1

	heap profile: 4: 3280 [1645: 13352240] @ heap/1048576

	[4:] 		Currently live objects,
	[3280] 		Amount of memory occupied by live objects
	[1645:] 	Total number of allocations
	[13352240] 	Amount of memory occupied by all allocations

#### Interactive Profiling

Put some load of the web application using a single connection.

 	boom -m POST -c 1 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the Go pprof tool in another window or tab to review heap information.

	go tool pprof -<PICK_MEM_PROFILE> ./project http://localhost:5000/debug/pprof/heap

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

Run the Go pprof tool in another window or tab to review cpu information.

	go tool pprof ./project http://localhost:5000/debug/pprof/profile

	_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not
	(except for goroutines that called runtime.LockOSThread, which is, unfortunately, not
	visible in the profile). Note that goroutines in "IO wait" state also do not consume
	threads, they are parked on non-blocking network poller
	(which uses epoll/kqueue/GetQueuedCompletionStatus to unpark goroutines later)._

Explore using the **top**, **list**, **web** and **web list** commands.

#### Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.

    curl -s http://localhost:5000/debug/pprof/heap > base.heap

After some time, take another snapshot:

    curl -s http://localhost:5000/debug/pprof/heap > current.heap

Now compare both snapshots against the binary and get into the pprof tool:

    go tool pprof -inuse_space -base base.heap memory_trace current.heap

#### Flame Graphs

Go-Torch is a tool for stochastically profiling Go programs. Collects stack traces and synthesizes them into a flame graph.

	https://github.com/uber/go-torch

Put some load of the web application.

	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the torch tool and visualize the profile.

	go-torch -u http://localhost:5000/

### Benchmark Profiling

Run the benchmarks and produce a cpu and memory profile.

	cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/profiling/project/search
	
	go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
	go tool pprof ./search.test cpu.out
	(pprof) web list rssSearch

	go test -run none -bench . -benchtime 3s -benchmem -memprofile mem.out
	go tool pprof -inuse_space ./search.test mem.out
	(pprof) web list rssSearch

### Trace Profiles

#### Trace Web Application

Put some load of the web application.

	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Capture a trace file for a brief duration.

	curl -s http://localhost:5000/debug/pprof/trace?seconds=2 > trace.out

Run the Go trace tool.

	go tool trace trace.out

Use the RSS Search test instead.

	cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/profiling/project/search
	go test -run none -bench . -benchtime 3s -trace trace.out
	go tool trace trace.out

## Expvar

Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.

### Adding New Variables

	import "expvar"

	// expvars is adding the goroutine counts to the variable set.
	func expvars() {

		// Add goroutine counts to the variable set.
		gr := expvar.NewInt("Goroutines")
		go func() {
			for _ = range time.Tick(time.Millisecond * 250) {
				gr.Set(int64(runtime.NumGoroutine()))
			}
		}()
	}

	// main is the entry point for the application.
	func main() {
		expvars()
		service.Run()
	}

### Expvarmon

TermUI based Go apps monitor using expvars variables (/debug/vars). Quickest way to monitor your Go app.

	go get github.com/divan/expvarmon

Running expvarmon

	expvarmon -ports=":5000" -vars="requests,goroutines,mem:memstats.Alloc"

## Godoc Analysis

The `godoc` tool can help you perform static analysis on your code.

	// Perform a pointer analysis and then run the godoc website.
	godoc -analysis pointer -http=:8080

[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team

## HTTP Tracing

HTTP tracing facilitate the gathering of fine-grained information throughout the lifecycle of an HTTP client request.

[HTTP Tracing Package](http_trace)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
