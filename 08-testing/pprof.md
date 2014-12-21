## Profiling Code

We can use the go tooling in conjunction with the Graph Visualization Tools and Ghostscript. These tools will allow us to graph the profiles we create.

## Installing Tools

[download files](https://drive.google.com/?pli=1&authuser=0#folders/0B8nQmHFH90Pkck13MVVLcko5OGc)

### Dave Cheney's Profile Package
`go get` Dave Cheney's profiling package. He has done a nice job abstracting all the boilerplate code required. If you are interested in understanding how to do this without the profile package, read this (http://saml.rilspace.org/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof)

	go get github.com/davecheney/profile

### Graph Visualization Tools
Download the package for your target OS/Arch:
[http://www.graphviz.org/Download.php](http://www.graphviz.org/Download.php)

### Ghostscript
This is not an easy step for Mac users since there is no prebuilt distribution.

Download and uncompress the source code:
[http://ghostscript.com/download/gsdnld.html](http://ghostscript.com/download/gsdnld.html)

	./configure
	make
	make install

### Code Changes
We need to add some changes to main to get the profiling data we need.

    import "github.com/davecheney/profile"

	// main is the entry point for the application.
	func main() {
		cfg := profile.Config{
			MemProfile:     true,
			CPUProfile:     true,
			ProfilePath:    ".",  // store profiles in current directory
			NoShutdownHook: true, // do not hook SIGINT
		}

		// p.Stop() must be called before the program exits to
		// ensure profiling information is written to disk.
		p := profile.Start(&cfg)
		defer p.Stop()

		. . .
	}

### Running and Creating Profile Graph
	go build
	./example1
    go tool pprof --pdf ./example1 cpu.pprof > callgraph.pdf
    go tool pprof --pdf ./example1 mem.pprof > callgraph.pdf

### Peek into the runtime scheduler:
[http://golang.org/pkg/runtime/](http://golang.org/pkg/runtime/)

	GODEBUG=schedtrace=1000,scheddetail=1 ./example1.go

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

	SCHED 11951ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=6 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
	Scheduler States:
		gomaxprocs=2: Number of contexts being used.
		threads=4:    Number of threads being used.
		runqueue=6:   Number of goroutines in the global queue.
		gcwaiting=0:  Is the scheduled blocking waiting for GC to finish.


  	P0: status=1 schedtick=11 syscalltick=0 m=3 runqsize=1 gfreecnt=0
  	P1: status=1 schedtick=13 syscalltick=1 m=2 runqsize=1 gfreecnt=0
  	Context Stats:
  		m=3:        The thread being used.
  		runqsize=1: Number of goroutines assigned to the context.

  	M3: p=0 curg=5 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
  	M2: p=1 curg=12 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
  	M1: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=1 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
  	M0: p=-1 curg=-1 mallocing=0 throwing=0 gcing=0 locks=0 dying=0 helpgc=0 spinning=0 blocked=0 lockedg=-1
  	Thread Stats:
  		p=0: The Context this thread is attached to.

  	G1: status=4(semacquire) m=-1 lockedm=-1
  	G2: status=4(force gc (idle)) m=-1 lockedm=-1
  	G3: status=4(GC sweep wait) m=-1 lockedm=-1
  	G4: status=1(stack growth) m=-1 lockedm=-1
  	G5: status=2(sleep) m=3 lockedm=-1
  	G6: status=1(sleep) m=-1 lockedm=-1
  	G7: status=1(stack growth) m=-1 lockedm=-1
  	G8: status=1(stack growth) m=-1 lockedm=-1
  	G9: status=1(stack growth) m=-1 lockedm=-1
  	G10: status=1(stack growth) m=-1 lockedm=-1
  	G11: status=1(stack growth) m=-1 lockedm=-1
  	G12: status=2(sleep) m=2 lockedm=-1
  	G13: status=1(sleep) m=-1 lockedm=-1
  	G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
  	Goroutine Stats:
  		m=3: The thread being used.
  		status: http://golang.org/src/runtime/runtime.h
  			Gidle,                                 // 0
   			Grunnable,                             // 1 runnable and on a run queue
   			Grunning,                              // 2
   			Gsyscall,                              // 3
   			Gwaiting,                              // 4
   			Gmoribund_unused,                      // 5 currently unused, but hardcoded in gdb scripts
   			Gdead,                                 // 6
   			Genqueue,                              // 7 Only the Gscanenqueue is used.
   			Gcopystack,                            // 8 in this state when newstack is moving the stack
   			View runtime.h for more

### Important Read
[Go Debugging](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)
