## GODEBUG

We can get specific trace information about the garbage collection and scheduler using the GODEBUG environmental variable. The variable will cause the runtime to emit tracing information.

## Schedule Tracing

	$ export GODEBUG=schedtrace=1000

	*scheddetail*: setting schedtrace=X and scheddetail=1 causes the scheduler to emit
	detailed multiline info every X milliseconds, describing state of the scheduler,
	processors, threads and goroutines.

	*schedtrace*: setting schedtrace=X causes the scheduler to emit a single line to standard
	error every X milliseconds, summarizing the scheduler state.

	SCHED 1009ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 [4 4]

		gomaxprocs=1:  Contexts configured.
		idleprocs=0:   Contexts not in use. Goroutine running.
		threads=3:     Threads in use.
		idlethreads=0: Threads not in use.
		runqueue=0:    Goroutines in the global queue.
		[4 4]:         Goroutines in each of the logical processors.

### Generating a Scheduler Trace

Build and run the example program using a single logical processor.

	$ go build
	$ GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./godebug

Put some load of the web application.

	$ hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
	
Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds we don't see any more goroutines in the trace.

    SCHED 8047ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [62]
    SCHED 9056ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=32 [0]
    SCHED 10065ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]
    SCHED 11068ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]

Run the example program but leak goroutines.

	$ GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./godebug leak

Put some load of the web application.

	$ hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"
	
Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds we still see goroutines in the trace.

    SCHED 13074ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [37]
    SCHED 14084ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
    SCHED 15091ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
    SCHED 16097ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=129 [225]

Let's run with two logical processors now:

	$ GOMAXPROCS=2 GODEBUG=schedtrace=1000 ./godebug

## GC Tracing

There is no way to identify specifically in the code where a leak is occurring. We can validate if a memory leak is present and which functions or methods are producing the most allocations.

Setting **gctrace=1** causes the garbage collector to emit a single line to standard error at each collection, summarizing the amount of memory collected and the length of the pause. Setting gctrace=2 emits the same summary but also repeats each collection. The format of this line is subject to change:

    $ export GODEBUG=gctrace=1

    gc # @#s #%: #+...+# ms clock, #+...+# ms cpu, #->#-># MB, # MB goal, # P

Where the fields are as follows:

    gc #        the GC number, incremented at each GC
    @#s         time in seconds since program start
    #%          percentage of time spent in GC since program start
    #+...+#     wall-clock/CPU times for the phases of the GC
    #->#-># MB  heap size at GC start, at GC end, and live heap
    # MB goal   goal heap size
    # P         number of processors used

**wall-clock** time is a measure of the real time that elapses from start to end, including time that passes due to programmed (artificial) delays or waiting for resources to become available.
https://en.wikipedia.org/wiki/Wall-clock_time

**CPU time** (or process time) is the amount of time for which a central processing unit (CPU) was used for processing instructions of a computer program or operating system, as opposed to, for example, waiting for input/output (I/O) operations or entering low-power (idle) mode.
https://en.wikipedia.org/wiki/CPU_time

You can get more details by adding the **gcpacertrace=1** flag. This causes the garbage collector to print information about the internal state of the concurrent pacer.

    $ export GODEBUG=gctrace=1,gcpacertrace=1

Sample output:

    gc 5 @0.071s 0%: 0.018+0.46+0.071 ms clock, 0.14+0/0.38/0.14+0.56 ms cpu, 29->29->29 MB, 30 MB goal, 8 P
    pacer: sweep done at heap size 29MB; allocated 0MB of spans; swept 3752 pages at +6.183550e-004 pages/byte
    pacer: assist ratio=+1.232155e+000 (scan 1 MB in 70->71 MB) workers=2+0
    pacer: H_m_prev=30488736 h_t=+2.334071e-001 H_T=37605024 h_a=+1.409842e+000 H_a=73473040 h_g=+1.000000e+000 H_g=60977472 u_a=+2.500000e-001 u_g=+2.500000e-001 W_a=308200 goalΔ=+7.665929e-001 actualΔ=+1.176435e+000 u_a/u_g=+1.000000e+000

Notes:

    In C++, a memory leak is memory you have lost a reference to.  
    In Go, a memory leak is memory you retain a reference to.

### Generating a GC Trace

Build and run the example program.

    $ go build
    $ GODEBUG=gctrace=1 ./godebug

Put some load of the web application.

    $ hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"

Review the gc trace.

    gc 318 @36.750s 0%: 0.022+0.27+0.040 ms clock, 0.13+0.60/0.43/0.031+0.24 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 319 @36.779s 0%: 0.019+0.24+0.035 ms clock, 0.15+0.43/0.26/0+0.28 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 320 @36.806s 0%: 0.023+0.34+0.035 ms clock, 0.18+0.63/0.49/0.014+0.28 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 321 @36.834s 0%: 0.026+0.20+0.044 ms clock, 0.18+0.50/0.34/0.001+0.31 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 322 @36.860s 0%: 0.022+0.29+0.055 ms clock, 0.13+0.54/0.47/0+0.33 ms cpu, 4->4->0 MB, 5 MB goal, 8 P

    gc 318      : First GC run since program started.
    @36.750s    : Nine milliseconds since the program started.
    0%          : One percent of the programs time has been spent in GC.
    
    // wall-clock
    0.022ms     : **STW** Sweep termination - Wait for all Ps to reach a GC safe-point.
    0.27ms      : Mark/Scan
    0.040ms     : **STW** Mark termination - Drain any remaining work and perform housekeeping.

    // CPU time
    0.13ms      : **STW** Sweep termination - Wait for all Ps to reach a GC safe-point.
    0.60ms      : Mark/Scan - Assist Time (GC performed in line with allocation)
    0.43ms      : Mark/Scan - Background GC time
    0.031ms     : Mark/Scan - Idle GC time
    0.24ms      : **STW** Mark termination - Drain any remaining work and perform housekeeping.

    4MB         : Heap size at GC start
    4MB         : Heap size at GC end
    0MB         : Live Heap
    5MB         : Goal heap size
    8P          : Number of logical processors

## Links

[Tour of Go's env variables](https://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables) - Dave Cheney    
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov    
[Scheduler Tracing In Go](https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html) - William Kennedy    

[Go runtime package](http://golang.org/pkg/runtime/)    
[GC Runtime Source Code](https://golang.org/src/runtime/mgc.go)    

[Finding memory leaks in Go](https://www.hakkalabs.co/articles/finding-memory-leaks-go-programs) - Oleg Shaldybin    
[Visualising the Go garbage collector](https://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector) - Dave Cheney    
[Understanding Go memory usage](https://web.archive.org/web/20170925060611/https://deferpanic.com/blog/understanding-golang-memory-usage/)    
[Visualising garbage collection algorithms](https://spin.atomicobject.com/2014/09/03/visualizing-garbage-collection-algorithms/) - Ken Fox    

## Code Review

[Web Service](godebug.go) ([Go Playground](https://play.golang.org/p/wr1FEf0_gKG))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
