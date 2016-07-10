## Profiling Code

We can use the go tooling in conjunction with the Graph Visualization Tools and Ghostscript. These tools will allow us to graph the profiles we create.

## Installing Tools

**Graph Visualization Tools**    
Download the package for your target OS/Arch:
[http://www.graphviz.org/Download.php](http://www.graphviz.org/Download.php)

**Ghostscript**    
Download and uncompress the source code:
[http://ghostscript.com/download/gsdnld.html](http://ghostscript.com/download/gsdnld.html)

	./configure
	make
	sudo make install

**go-wrk**  
go-wrk is a modern HTTP benchmarking tool capable of generating significant load when run on a single multi-core CPU. It builds on go language go routines and scheduler for behind the scenes async IO and concurrency.

	go get -u github.com/tsliwowicz/go-wrk

## Building and Running the Project

We have a website that we will use to learn and explore more about profiling. This project is a search engine for RSS feeds. Run the website and validate it is working.

	go build
	./project

	http://localhost:5000/search

## Adding Load

To add load to the service while running profiling we can run these command.

	Use 10 connections for 2 minute on CNN, BBC and NYT about house:
	go-wrk -M POST -c 10 -d 120 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

## GODEBUG

GODEBUG is an environment variable that allows us to get information from the runtime about the scheduler and the garabage collector.

### The Basics

Learn the basics of using GODEBUG for tracing.  
[Memory Tracing](godebug/gctrace) | [Scheduler Tracing](godebug/schedtrace)

### Memory Trace for Project

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=gctrace=1 ./project > /dev/null

Put some load of the web application for 20 seconds.

	go-wrk -M POST -c 10 -d 20 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

### Scheduler Trace for Project

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=schedtrace=1000 GOMAXPROCS=2 ./project > /dev/null

Put some load of the web application for 20 seconds.

	go-wrk -M POST -c 10 -d 20 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

## PPROF

Go provides built in support for retrieving profiling data from your running Go applications.

### Raw http/pprof

Add the following import so we can include the profiling route to our web service.

	import _ "net/http/pprof"

Build the project again and start it.

	go build
	./project

Look at the basic profiling stats from the new endpoint.

	http://localhost:5000/debug/pprof

Run a single search from the Browser and then refresh the profile information.

	http://localhost:5000/search?term=house&cnn=on

Put some load of the web application for 10 seconds. Review the raw profiling information once again.

	go-wrk -M POST -c 10 -d 10 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

### Interactive Profiling

Using the Go pprof tool we can interact with the profiling data.

Put some load of the web application for 2 minutes using a single connection.

 	go-wrk -M POST -c 1 -d 120 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the Go pprof tool in another window or tab to review heap information.

	go tool pprof ./project http://localhost:5000/debug/pprof/heap

Run the Go pprof tool in another window or tab to review cpu information.

	go tool pprof ./project http://localhost:5000/debug/pprof/profile

Explore using the **top** and **list** commands.

### Generate PDF Call Graph

Generate call graphs for both the cpu and memory profiles.

	// Create output files.
	curl -s http://localhost:5000/debug/pprof/profile > cpu.out
	go tool pprof ./project cpu.out
	(pprof) web

	curl -s http://localhost:5000/debug/pprof/heap > mem.out
	go tool pprof ./project mem.out
	(pprof) web

	// Call into the endpoints directly and generate graphs.
	go tool pprof -web ./project http://localhost:5000/debug/pprof/heap
	go tool pprof -web ./project http://localhost:5000/debug/pprof/profile

## Go Torch

Tool for stochastically profiling Go programs. Collects stack traces and synthesizes them into a flame graph.

	https://github.com/uber/go-torch

Put some load of the web application for 2 minutes.

	go-wrk -M POST -c 10 -d 120 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the torch tool and visualize the profile.

	go-torch -u http://localhost:5000/

## Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.

    curl -s http://localhost:5000/debug/pprof/heap > base.heap

After some time, take another snapshot:

    curl -s http://localhost:5000/debug/pprof/heap > current.heap

Now compare both snapshots against the binary and get into the pprof tool:

    go tool pprof -alloc_space -base base.heap memory_trace current.heap

    -inuse_space  : Display in-use memory size
    -inuse_objects: Display in-use object counts
    -alloc_space  : Display allocated memory size
    -alloc_objects: Display allocated object counts

## Tracing

Run the web application with tracing on. Show the code that produces the trace.

	./project trace

Put some load of the web application for 2 minutes.

	go-wrk -M POST -c 10 -d 120 -no-ka "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Kill the web application to produce the trace.out file.

Run the Go trace tool.

	go tool trace trace.out

## Links

http://golang.org/blog/profiling-go-programs  
http://golang.org/pkg/runtime/pprof/  
https://golang.org/pkg/net/http/pprof/  
https://godoc.org/rsc.io/pprof_mac_fix

## Code Review

[HTTP Service](helloHTTP.go) ([Go Playground](http://play.golang.org/p/XcpEreJ9zg))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
