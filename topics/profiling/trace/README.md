## Trace Profiling

The trace profiling can help identify the health of your programs. We will use a simple program to learn how to navigate and read some of the tracing information you can find in the trace tool.

## Trace Command

    // Run the test and produce a trace profile. Add the gctrace as well so
    // we can compare what we see in the trace with the gctrace outout.
    GODEBUG=gctrace=1 go test -run TestSTWTrace -trace trace.out

    // Run the trace tool with the generated profile.
    go tool trace trace.out

Navigating the tracing tool and interpreting the data requires in class instruction.

## Code Review
 
[Profiling Test](trace_test.go) ([Go Playground](https://play.golang.org/p/nmDRsb4Dhj))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
