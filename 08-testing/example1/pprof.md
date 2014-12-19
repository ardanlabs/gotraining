## Profiling Code

We can use the go tooling in conjunction with the Graph Visualization Tools and Ghostscript. These tools will allow us to graph the profiles we create.

## Installing Tools

[downloads](https://drive.google.com/?pli=1&authuser=0#folders/0B8nQmHFH90Pkck13MVVLcko5OGc)

### Dave Cheney's Profile Package
`go get` Dave Cheney's profiling package. He has done a nice job abstracting all the boilerplate code required. If you are interested in understanding how to do this without the profile package, read this (http://saml.rilspace.org/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof)

	go get github.com/davecheney/profile

### Graph Visualization Tools
Download the package for your target OS/Arch

	http://www.graphviz.org/Download.php

### Ghostscript
This is not an easy step for Mac users since there is no prebuilt distribution.

	Download and uncompress the source code:
	http://ghostscript.com/download/gsdnld.html

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
GODEBUG=schedtrace=1000,scheddetail=1 ./example1.go

### Important Read
[Go Debugging](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)
