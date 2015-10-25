## Profiling Code

We can use the go tooling in conjunction with the Graph Visualization Tools and Ghostscript. These tools will allow us to graph the profiles we create.

Note: Unless you are running OS X El Capitan, profiling on the Mac is broken. This post talks about how to hack to OS X Kernel to make it work.  
[http://research.swtch.com/macpprof](http://research.swtch.com/macpprof)  
[https://godoc.org/rsc.io/pprof_mac_fix](https://godoc.org/rsc.io/pprof_mac_fix)

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
	sudo make install

### Code Changes
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
		./05-profiling
	
	In a separate terminal generate requests:
		while true && do curl http://localhost:6060/english && done
		<control> C
    
    Generate the call graphs:
    	go tool pprof --pdf ./05-profiling cpu.pprof > cpugraph.pdf
		go tool pprof --pdf ./05-profiling mem.pprof > memgraph.pdf

    See all the options:
    	go tool pprof -h

### HTTP Runtime Profile Graphs
	Build and run the service:
		go build
		./05-profiling
	
	In a separate terminal generate requests:
		while true && do curl http://localhost:6060/english && done

	In a separate terminal run a pprof tool:
		go tool pprof http://localhost:6060/debug/pprof/heap
		go tool pprof http://localhost:6060/debug/pprof/profile
		go tool pprof http://localhost:6060/debug/pprof/block

	Play with the different commands:
		go tool pprof -help

## Links

http://golang.org/blog/profiling-go-programs

http://golang.org/pkg/runtime/pprof/

https://golang.org/pkg/net/http/pprof/

https://godoc.org/rsc.io/pprof_mac_fix

## Code Review

[HTTP Service](helloHTTP.go) ([Go Playground](https://play.golang.org/p/fcU9jQX2Qz))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
