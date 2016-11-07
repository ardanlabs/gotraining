## Schedule Tracing

We can get specific information about the scheduler using the GODEBUG environmental variable. The variable will cause the schedule to emit information about the health of the logical processors.

## GODEBUG

	export GODEBUG=schedtrace=1000

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

## Running GODEBUG Scheduler Trace

This example shows a simple web api running with some basic load.

	Build and run the example program using a single logical processor.

		go build
		GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./schedtrace

	Put some load of the web application.

		boom -m POST -c 8 -n 10000 "http://localhost:4000/sendjson"
	
	Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds
	we don't see any more goroutines in the trace.

		SCHED 8047ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [62]
		SCHED 9056ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=32 [0]
		SCHED 10065ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]
		SCHED 11068ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]

	Run the example program but leak goroutines.

		GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./schedtrace leak

	Put some load of the web application.

		boom -m POST -c 8 -n 10000 "http://localhost:4000/sendjson"
	
	Look at the load on the logical processor. We can only see runnable goroutines. After 5 seconds
	we still see goroutines in the trace.

		SCHED 13074ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [37]
		SCHED 14084ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
		SCHED 15091ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
		SCHED 16097ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=129 [225]

	Let's run with two logical processors now:

		GOMAXPROCS=2 GODEBUG=schedtrace=1000 ./schedtrace

## Links

[Tour of Go's env variables](http://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables)   
[http://golang.org/pkg/runtime/](http://golang.org/pkg/runtime/)  
https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs    
http://www.goinggo.net/2015/02/scheduler-tracing-in-go.html  

## Code Review

[Scheduler Trace](trace.go) ([Go Playground](https://play.golang.org/p/iyRaSsjQSS))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
