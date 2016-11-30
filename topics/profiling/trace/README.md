## Trace Profiling

The trace profiling can help identify the health of your programs. We will use a simple program to learn how to navigate and read some of the tracing information you can find in the trace tool.

## Trace Command

    // Run the website and hit the /work handler.
    // The hit <control><C> to shut the service down.

    // Run the trace tool with the generated profile.
    go tool trace trace.out

    // Look at the profile data as well.
    go tool pprof ./trace cpu.pprof

Navigating the tracing tool and interpreting the data requires in class instruction.

## Code Review
 
[Profiling Test](trace.go) ([Go Playground](https://play.golang.org/p/N2Q-djdhIk))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
