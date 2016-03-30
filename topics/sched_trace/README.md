## Schedule Tracing

We can get specific information about the scheduler using the GODEBUG environmental variable. The variable will cause the schedule to emit information about the health of the logical processors.

We will need this program to place load on the example web api's. Download and build this tool.
	https://github.com/wg/wrk

## GODEBUG

[http://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables](Tour of Go's env variables)

[http://golang.org/pkg/runtime/](http://golang.org/pkg/runtime/)

	export GODEBUG=schedtrace=1000,scheddetail=1

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

## Links

https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs

http://www.goinggo.net/2015/02/scheduler-tracing-in-go.html

## Running Summary Trace: Example1

This example shows a simple web api running with some basic load.

	Build and run the example program using a single logical processor.

		go build
		export GOMAXPROCS=1
		export GODEBUG=schedtrace=1000
		./example1

	Start to apply load to the web api:
	
		wrk -t8 -c500 -d5s http://localhost:4000/sendjson

	Look at the load on the logical processor. We can only see runnable goroutines:

		SCHED 1009ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]
		SCHED 4026ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=32 [25]

	Look at the general stats from the profile:

		http://localhost:4000/debug/pprof
		/debug/pprof/
			profiles:
				0	block
				499	goroutine
				19	heap
				4	threadcreate

	Let's run with two logical processors now:

		export GOMAXPROCS=2
		./example1

	Look at the load on the logical processor. We can only see runnable goroutines:

		SCHED 9056ms: gomaxprocs=2 idleprocs=0 threads=6 spinningthreads=0 idlethreads=2 runqueue=79 [13 0]
		SCHED 10062ms: gomaxprocs=2 idleprocs=0 threads=6 spinningthreads=0 idlethreads=2 runqueue=4 [9 4]

	Look at the general stats from the profile:
	
		http://localhost:4000/debug/pprof
		/debug/pprof/
			profiles:
				0	block
				499	goroutine
				22	heap
				5	threadcreate

## Running Summary Trace: Example2

This example shows a simple web api running that is leaking goroutines.

	Build and run the example program using a single logical processor.

		go build
		export GOMAXPROCS=1
		export GODEBUG=schedtrace=1000
		./example2

	Start to apply load to the web api:
	
		wrk -t8 -c500 -d5s http://localhost:4000/sendjson

	Look at the load on the logical processor. We can only see runnable goroutines:

		SCHED 1009ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [0]
		SCHED 4026ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=32 [25]

	Look at the general stats from the profile:

		http://localhost:4000/debug/pprof
		/debug/pprof/
			profiles:
				0	block
				499	goroutine
				19	heap
				4	threadcreate

		http://localhost:4000/debug/pprof
		/debug/pprof/
			profiles:
				0	block
				694	goroutine
				23	heap
				7	threadcreate

## Code Review

[Web API](example1/example1.go) ([Go Playground](https://play.golang.org/p/70aRkw59zH))

[Web API Leaking Goroutines](example2/example2.go) ([Go Playground](https://play.golang.org/p/jWMR2oudEN))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
