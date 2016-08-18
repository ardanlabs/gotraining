## Profiling Code

We can use the go tooling to inspect and profile our programs. Profiling is more of a journey and detective work. It requires some understanding about the application and expectations. The profiling data in and of itself is just raw numbers. We have to give it meaning and understanding.

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

**boom**  
boom is a modern HTTP benchmarking tool capable of generating the load you need to run tests. It's built using the Go language and leverages goroutines for behind the scenes async IO and concurrency.

	go get -u github.com/rakyll/boom

## The Basics

Let's learn the basics of viewing and working with profiling data.

Learn the basics of using GODEBUG.  
[Memory Tracing](godebug/gctrace) | [Scheduler Tracing](godebug/schedtrace)

Learn the basics of using benchmarking.  
[Benchmarking](benchmarks)

Learn the basics of using tests/tracing.  
[Tracing](trace)

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

#### Memory Tracing

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=gctrace=1 ./project > /dev/null

Put some load of the web application.

	boom -m POST -c 8 -n 10000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

#### Scheduler Tracing

Run the website redirecting the stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=schedtrace=1000 GOMAXPROCS=2 ./project > /dev/null

Put some load of the web application.

	boom -m POST -c 8 -n 10000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

### PPROF

#### Raw http/pprof

We already added the following import so we can include the profiling route to our web service.

	import _ "net/http/pprof"

Look at the basic profiling stats from the new endpoint.

	http://localhost:5000/debug/pprof

Run a single search from the Browser and then refresh the profile information.

	http://localhost:5000/search?term=house&cnn=on

Put some load of the web application. Review the raw profiling information once again.

	boom -m POST -c 8 -n 10000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

#### Interactive Profiling

Put some load of the web application using a single connection.

 	boom -m POST -c 1 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the Go pprof tool in another window or tab to review heap information.

	go tool pprof ./project http://localhost:5000/debug/pprof/heap

Run the Go pprof tool in another window or tab to review cpu information.

	go tool pprof ./project http://localhost:5000/debug/pprof/profile

Explore using the **top**, **list**, **web** and **web list** commands.

#### Comparing Profiles

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

#### Go Torch

Tool for stochastically profiling Go programs. Collects stack traces and synthesizes them into a flame graph.

	https://github.com/uber/go-torch

Put some load of the web application.

	boom -m POST -c 8 -n 100000 "http://localhost:5000/search?term=house&cnn=on&bbc=on&nyt=on"

Run the torch tool and visualize the profile.

	go-torch -u http://localhost:5000/

### Benchmarks

Run the test and produce a cpu and memory profile.

	cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/profiling/project/search
	
	go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
	go tool pprof ./search.test cpu.out
	(pprof) web list rssSearch

	go test -run none -bench . -benchtime 3s -benchmem -memprofile mem.out
	go tool pprof -alloc_space ./search.test mem.out
	(pprof) web list rssSearch

### Tracing / Blocking Profiles

#### Tracing Web Application

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

The `godoc` tool can help you perform analysis on your code.

[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team

## Links

http://golang.org/pkg/runtime/pprof/  
https://golang.org/pkg/net/http/pprof/  
[Profiling & Optimizing in Go](https://www.youtube.com/watch?v=xxDZuPEgbBU) - Brad Fitzpatrick  
[Go Dynamic Tools](https://www.youtube.com/watch?v=a9xrxRsIbSU) - Dmitry Vyukov  
[How NOT to Measure Latency](https://www.youtube.com/watch?v=lJ8ydIuPFeU&feature=youtu.be) - Gil Tene  
[Go Performance Tales](http://jmoiron.net/blog/go-performance-tales) - Jason Moiron  
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov  
[Reduce allocation in Go code](https://methane.github.io/2015/02/reduce-allocation-in-go-code) - Python Bytes  
[Write High Performance Go](http://go-talks.appspot.com/github.com/davecheney/presentations/writing-high-performance-go.slide) - Dave Cheney  
[Profiling Go Programs](http://golang.org/blog/profiling-go-programs) - Go Team  
[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team

## Code Review

[HTTP Service](helloHTTP.go) ([Go Playground](http://play.golang.org/p/XcpEreJ9zg))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
