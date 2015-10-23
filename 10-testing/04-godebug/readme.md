## GODEBUG

There is a special environmental variable named GODEBUG that will emit debugging information about the runtime as your program executes. You can request summary and detailed information for both the garbage collector and the scheduler. What’s great is you don't need to build your program with any special switches for it to work.

## Notes

* View the internals of the runtime and scheduler.
* Look and details about memory and goroutines.
* Helps to determine how your concurrent program is running.

### GODEBUG Documentation

[http://golang.org/pkg/runtime/](http://golang.org/pkg/runtime/)

	GODEBUG=schedtrace=1000,scheddetail=1 ./example3.go

	*allocfreetrace*: setting allocfreetrace=1 causes every allocation to be
	profiled and a stack trace printed on each object's allocation and free.

	*efence*: setting efence=1 causes the allocator to run in a mode
	where each object is allocated on a unique page and addresses are
	never recycled.

	*gctrace*: setting gctrace=1 causes the garbage collector to emit a single line to standard
	error at each collection, summarizing the amount of memory collected and the
	length of the pause. Setting gctrace=2 emits the same summary but also
	repeats each collection.

	*gcdead*: setting gcdead=1 causes the garbage collector to clobber all stack slots
	that it thinks are dead.

	*invalidptr*: defaults to invalidptr=1, causing the garbage collector and stack
	copier to crash the program if an invalid pointer value (for example, 1)
	is found in a pointer-typed location. Setting invalidptr=0 disables this check.
	This should only be used as a temporary workaround to diagnose buggy code.
	The real fix is to not store integers in pointer-typed locations.

	*scheddetail*: setting schedtrace=X and scheddetail=1 causes the scheduler to emit
	detailed multiline info every X milliseconds, describing state of the scheduler,
	processors, threads and goroutines.

	*schedtrace*: setting schedtrace=X causes the scheduler to emit a single line to standard
	error every X milliseconds, summarizing the scheduler state.

	*scavenge*: scavenge=1 enables debugging mode of heap scavenger.

	Example
	http://golang.org/src/runtime/proc.c

	GODEBUG=schedtrace=1000 ./example3.go

	GOMAXPROCS = 1
		SCHED 0ms: gomaxprocs=1 idleprocs=0 threads=2 spinningthreads=0 idlethreads=0 runqueue=0 [1]
		SCHED 1009ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 [9]

	GOMAXPROCS = 2
		SCHED 1001ms: gomaxprocs=2 idleprocs=2 threads=4 spinningthreads=0 idlethreads=2 runqueue=0 [0 0]
		SCHED 2002ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [4 4]
	
	Scheduler States:
		gomaxprocs=1:  Contexts configured.
		idleprocs=0:   Contexts not in use. Goroutine running.
		threads=3:     Threads in use.
		idlethreads=0: Threads not in use.
		runqueue=0:    Goroutines in the global queue.
		[9]:           Goroutines in a context's run queue.
		[4 4]:         Goroutines in each of the context's run queue.

	GODEBUG=schedtrace=1000,scheddetail=1 ./example3.go

	GOMAXPROCS = 1
		SCHED 2016ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
		P0: status=1 schedtick=20 syscalltick=14 m=0 runqsize=9 gfreecnt=0
		M2: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
		M1: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=1 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
		M0: p=0 curg=5 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
		G1: status=4(semacquire) m=-1 lockedm=-1
		G2: status=4(force gc (idle)) m=-1 lockedm=-1
		G3: status=4(GC sweep wait) m=-1 lockedm=-1
		G4: status=4(finalizer wait) m=-1 lockedm=-1
		G5: status=2(sleep) m=0 lockedm=-1
		G6: status=1(sleep) m=-1 lockedm=-1
		G7: status=1(sleep) m=-1 lockedm=-1
		G8: status=1(sleep) m=-1 lockedm=-1
		G9: status=1(sleep) m=-1 lockedm=-1
		G10: status=1(sleep) m=-1 lockedm=-1
		G11: status=1(sleep) m=-1 lockedm=-1
		G12: status=1(sleep) m=-1 lockedm=-1
		G13: status=1(sleep) m=-1 lockedm=-1
		G14: status=1(sleep) m=-1 lockedm=-1
		G15: status=4(timer goroutine (idle)) m=-1 lockedm=-1

	GOMAXPROCS = 2
		SCHED 2007ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
		P0: status=1 schedtick=20 syscalltick=12 m=3 runqsize=4 gfreecnt=0
	  	P1: status=1 schedtick=8 syscalltick=2 m=2 runqsize=4 gfreecnt=0
	  	M3: p=0 curg=11 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M2: p=1 curg=6 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M1: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=1 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M0: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	G1: status=4(semacquire) m=-1 lockedm=-1
	  	G2: status=4(force gc (idle)) m=-1 lockedm=-1
	  	G3: status=4(GC sweep wait) m=-1 lockedm=-1
	  	G4: status=4(finalizer wait) m=-1 lockedm=-1
	  	G5: status=1(sleep) m=-1 lockedm=-1
	  	G6: status=2(sleep) m=2 lockedm=-1
	  	G7: status=1(sleep) m=-1 lockedm=-1
	  	G8: status=1(sleep) m=-1 lockedm=-1
	  	G9: status=1(sleep) m=-1 lockedm=-1
	  	G10: status=1(sleep) m=-1 lockedm=-1
	  	G11: status=2(sleep) m=3 lockedm=-1
	  	G12: status=1(sleep) m=-1 lockedm=-1
	  	G13: status=1(sleep) m=-1 lockedm=-1
	  	G14: status=1(sleep) m=-1 lockedm=-1
	  	G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1

  	Scheduler States:  
		gcwaiting=0: Is the scheduled blocking waiting for GC to finish.

	P's represent contexts:  
  		P0: status=1 schedtick=20 syscalltick=12 m=3 runqsize=4 gfreecnt=0
  		P1: status=1 schedtick=8 syscalltick=2 m=2 runqsize=4 gfreecnt=0
  		Context Stats:
  			m=3:        The thread being used.
  			runqsize=1: Number of goroutines in a context's run queue.

  	M's represent a thread:
	  	M3: p=0 curg=11 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M2: p=1 curg=6 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M1: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=1 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
	  	M0: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
  		Thread Stats:  
  			p=0: The Context this thread is attached to.

  	G's represent a goroutine:  
  		G1: status=4(semacquire) m=-1 lockedm=-1
	  	G2: status=4(force gc (idle)) m=-1 lockedm=-1
	  	G3: status=4(GC sweep wait) m=-1 lockedm=-1
	  	G4: status=4(finalizer wait) m=-1 lockedm=-1
	  	G5: status=1(sleep) m=-1 lockedm=-1
	  	G6: status=2(sleep) m=2 lockedm=-1
	  	G7: status=1(sleep) m=-1 lockedm=-1
	  	G8: status=1(sleep) m=-1 lockedm=-1
	  	G9: status=1(sleep) m=-1 lockedm=-1
	  	G10: status=1(sleep) m=-1 lockedm=-1
	  	G11: status=2(sleep) m=3 lockedm=-1
	  	G12: status=1(sleep) m=-1 lockedm=-1
	  	G13: status=1(sleep) m=-1 lockedm=-1
	  	G14: status=1(sleep) m=-1 lockedm=-1
	  	G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
  		Goroutine Stats:  
  			m=3: The thread being used.
  			status: http://golang.org/src/runtime/runtime.h
	  			Gidle,                                 // 0 
	   			Grunnable,                             // 1 runnable and on a run queue
	   			Grunning,                              // 2 running
	   			Gsyscall,                              // 3 performing a syscall
	   			Gwaiting,                              // 4 waiting for the runtime
	   			Gmoribund_unused,                      // 5 currently unused, but hardcoded in gdb scripts
	   			Gdead,                                 // 6 goroutine is dead
	   			Genqueue,                              // 7 only the Gscanenqueue is used.
	   			Gcopystack,                            // 8 in this state when newstack is moving the stack
	   			View runtime.h for more

## Links

https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs

http://www.goinggo.net/2015/02/scheduler-tracing-in-go.html

## Code Review

[Scheduler Stats](godebug.go) ([Go Playground](https://play.golang.org/p/M6_9Ir79EB))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
