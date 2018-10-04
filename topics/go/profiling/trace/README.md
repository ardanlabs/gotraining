## Tracing

The tracing can help identify not only what is happening but also what is not happening when your program is running. We will use a simple program to learn how to navigate and read some of the tracing information you can find in the trace tool.

## Basic Skills

Review this post to gain basic skills.

[go tool trace](https://making.pusher.com/go-tool-trace/) - Will Sewell  
[Debugging Latency in Go 1.11](https://medium.com/observability/debugging-latency-in-go-1-11-9f97a7910d68) - JBD

## Trace Command

You have two options with this code. First uncomment the CPU profile lines to generate a CPU profile.

    pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// trace.Start(os.Stdout)
	// defer trace.Stop()

This will let you run a profile first. Leverage the lessons learned in the other sections.

    $ ./trace > p.out
    $ go tool pprof p.out

Then run a trace by uncommenting the other lines of code.

    // pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

Once you run the program.

    $ ./trace > t.out
    $ go tool trace t.out

Then explore the trace tooling by building the program with these different find functions.

    n := find(topic, docs)
	// n := findConcurrent(topic, docs)
	// n := findConcurrentSem(topic, docs)
	// n := findNumCPU(topic, docs)
	// n := findActor(topic, docs)

Using this function allows you to see how to add custom tasks and regions. This requires Go version 1.11.

	// n := findNumCPUTasks(topic, docs)

_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not (except for goroutines that called runtime.LockOSThread, which is, unfortunately, not visible in the profile)._

_Note that goroutines in "IO wait" state do NOT consume an OS thread. They are parked on the non-blocking network poller._ 

## Code Review
 
[Profiling Test](trace.go) ([Go Playground](https://play.golang.org/p/AbAfYByFys0))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
