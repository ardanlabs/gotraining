## Profiling Code

We can use the go tooling in conjunction with the Graph Visualization Tools and Ghostscript. These tools will allow us to graph the profiles we create.

Note: Unless you are running OS X El Capitan, profiling on the Mac is broken. This post talks about how to hack to OS X Kernel to make it work.  
[http://research.swtch.com/macpprof](http://research.swtch.com/macpprof)  
[https://godoc.org/rsc.io/pprof_mac_fix](https://godoc.org/rsc.io/pprof_mac_fix)

## Installing Tools

[download files](https://drive.google.com/?pli=1&authuser=0#folders/0B8nQmHFH90Pkck13MVVLcko5OGc)

### Graph Visualization Tools
Download the package for your target OS/Arch:
[http://www.graphviz.org/Download.php](http://www.graphviz.org/Download.php)

### Ghostscript
This is not an easy step for Mac users since there is no prebuilt distribution.

Download and uncompress the source code:
[http://ghostscript.com/download/gsdnld.html](http://ghostscript.com/download/gsdnld.html)

	./configure
	make
	sudo make install

## Static Profiling

We can have our program produce profiling data and once the program quits, write the profiling data to disk so it can be visualized.

### Code Changes

`go get` Dave Cheney's profiling package. He has done a nice job abstracting all the boilerplate code required.

	go get github.com/davecheney/profile

We need to make some changes to main to get the profiling data we need.

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

### Running and Creating Profile Graphs

	Build and run the service:
		go build
		./profiling
	
	In a separate terminal generate requests:
		while true && do curl http://localhost:6060/english && done
		<control> C
    
    Generate the call graphs:
    	go tool pprof --pdf ./profiling cpu.pprof > cpugraph.pdf
		go tool pprof --pdf ./profiling mem.pprof > memgraph.pdf

## Dynamic Profiling

We can have our program produce profiling data and during runtime, write the profiling data to disk so it can be visualized.

### Code Changes

Reset the example back to the original code. Then add this import and this will add a route to the existing service we can leverage to get the profile data.

	_ "net/http/pprof"

### Running and Creating Profile Graphs

	Build and run the service:
		go build
		./profiling
	
	In a separate terminal generate requests:
		while true && do curl http://localhost:6060/english && done

	In a separate terminal run a pprof tool:
		go tool pprof http://localhost:6060/debug/pprof/heap		(Memory Profile)
		go tool pprof http://localhost:6060/debug/pprof/profile   	(CPU Profile)

### pprof Commands

We will use the CPU profile data:

	go tool pprof http://localhost:6060/debug/pprof/profile

The `top` command will show top functions and methods based on CPU Profile

	(pprof) top
	810ms of 810ms total (  100%)
	Showing top 10 nodes out of 98 (cum >= 10ms)
	      flat  flat%   sum%        cum   cum%
	     320ms 39.51% 39.51%      330ms 40.74%  syscall.Syscall
	     140ms 17.28% 56.79%      140ms 17.28%  runtime.kevent
	     120ms 14.81% 71.60%      120ms 14.81%  runtime.mach_semaphore_signal
	      70ms  8.64% 80.25%       70ms  8.64%  runtime.usleep
	      60ms  7.41% 87.65%       60ms  7.41%  runtime.mach_semaphore_wait
	      30ms  3.70% 91.36%       30ms  3.70%  runtime.mach_semaphore_timedwait
	      30ms  3.70% 95.06%       30ms  3.70%  runtime.memmove
	      20ms  2.47% 97.53%       20ms  2.47%  syscall.Syscall6
	      10ms  1.23% 98.77%       10ms  1.23%  runtime.cas64
	      10ms  1.23%   100%       10ms  1.23%  syscall.RawSyscall

The `web` command will generate a call graph.

## Links

http://golang.org/blog/profiling-go-programs

http://golang.org/pkg/runtime/pprof/

https://golang.org/pkg/net/http/pprof/

https://godoc.org/rsc.io/pprof_mac_fix

## Code Review

[HTTP Service](helloHTTP.go) ([Go Playground](https://play.golang.org/p/fcU9jQX2Qz))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
