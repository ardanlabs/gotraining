## Tracing

The tracing can help identify not only what is happening but also what is not happening when your program is running. We will use a simple program to learn how to navigate and read some of the tracing information you can find in the trace tool.

## Basic Skills

Review this post to gain basic skills.

[go tool trace](https://making.pusher.com/go-tool-trace/) - Will Sewell

## Trace Command

Run the program to download a file. Use the `LoadWrite()` function first and then try the `StreamWrite` function.

Build and run the program.

    $ go build
    $ ./trace

Run run the trace tool and inspect the trace.

    $ go tool trace trace.out

Generate a CPU profile.

    $ go tool trace -pprof=[net,syscall,sync,sched] trace.out > cpu.out
    
View the profile.

    $ go tool pprof ./trace cpu.out  

## Code Review
 
[Profiling Test](trace.go) ([Go Playground](https://play.golang.org/p/QJahKPIydE))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
