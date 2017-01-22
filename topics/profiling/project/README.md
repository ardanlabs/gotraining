## Profiling a Larger Web Service

We have a web application that extends a web service. Let's profile this application and attempt to understand how it is working.

### Building and Running the Project

We have a website that we will use to learn and explore more about profiling. This project is a search engine for RSS feeds. Run the website and validate it is working.

	go build
	./project

	http://localhost:5000/search

### Adding Load

To add load to the service while running profiling we can run these command.

	// Send 1M request using 8 connections.
	hey -m POST -c 8 -n 1000000 "http://localhost:5000/search?term=politics&cnn=on&bbc=on&nyt=on"

### GODEBUG

#### GC Trace

Run the website redirecting stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=gctrace=1 ./project > /dev/null

#### Scheduler Trace

Run the website redirecting stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.
	
	GODEBUG=schedtrace=1000 GOMAXPROCS=2 ./project > /dev/null

### PPROF

We already added the following import so we can include the profiling route to our web service.

	import _ "net/http/pprof"

#### Raw http/pprof

Look at the basic profiling stats from the new endpoint.

	http://localhost:5000/debug/pprof

Run a single search from the Browser and then refresh the profile information.

	http://localhost:5000/search?term=politics&cnn=on

Put some load of the web application and look at the heap profile.

	http://localhost:5000/debug/pprof/heap?debug=1

#### Interactive Profiling

Run the Go pprof tool in another window or tab to review heap information.

	go tool pprof -<PICK_MEM_PROFILE> ./project http://localhost:5000/debug/pprof/heap

	// Useful to see current status of heap.
	-inuse_space  : Allocations live at the time of profile  	** default
	-inuse_objects: Number of bytes allocated at the time of profile

	// Useful to see pressure on heap over time.
	-alloc_space  : All allocations happened since program start
	-alloc_objects: Number of object allocated at the time of profile

	If you want to reduce memory consumption, look at the `-inuse_space` profile collected during
	normal program operation.
	
	If you want to improve execution speed, look at the `-alloc_objects` profile collected after
	significant running time or at program end.

Run the Go pprof tool in another window or tab to review cpu information.

	go tool pprof ./project http://localhost:5000/debug/pprof/profile

	_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not
	(except for goroutines that called runtime.LockOSThread, which is, unfortunately, not
	visible in the profile). Note that goroutines in "IO wait" state also do not consume
	threads, they are parked on non-blocking network poller
	(which uses epoll/kqueue/GetQueuedCompletionStatus to unpark goroutines later)._

Explore using the **top**, **list**, **web** and **web list** commands.

#### Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.

    curl -s http://localhost:5000/debug/pprof/heap > base.heap

After some time, take another snapshot:

    curl -s http://localhost:5000/debug/pprof/heap > current.heap

Now compare both snapshots against the binary and get into the pprof tool:

    go tool pprof -inuse_space -base base.heap memory_trace current.heap

#### Flame Graphs

Go-Torch is a tool for stochastically profiling Go programs. Collects stack traces and synthesizes them into a flame graph.

	https://github.com/uber/go-torch

Put some load of the web application and run the torch tool and visualize the profile.

	go-torch -u http://localhost:5000/

### Benchmark Profiling

Run the benchmarks and produce a cpu and memory profile.

	cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/profiling/project/search
	
	go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
	go tool pprof ./search.test cpu.out
	(pprof) web list rssSearch

	go test -run none -bench . -benchtime 3s -benchmem -memprofile mem.out
	go tool pprof -inuse_space ./search.test mem.out
	(pprof) web list rssSearch

### Trace Profiles

#### Trace Web Application

Capture a trace file for a brief duration.

	curl -s http://localhost:5000/debug/pprof/trace?seconds=2 > trace.out

Run the Go trace tool.

	go tool trace trace.out

Use the RSS Search test instead to create a trace.

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
