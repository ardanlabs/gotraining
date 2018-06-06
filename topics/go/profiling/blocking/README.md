## Blocking Profiling

Testing and Tracing allows us to see blocking profiles.

## Running a Test Based Blocking Profile

We can get blocking profiles by running a test.

Generate a block profile from running the test.

	$ go test -blockprofile block.out

Run the pprof tool to view the blocking profile.

	$ go tool pprof block.out

Review the TestLatency function.

	$ list TestLatency

## Running a Trace

Once you have a test established you can use the **-trace trace.out** option with the go test tool.

Generate a trace from running the test.

	$ go test -trace trace.out

Run the trace tool to review the trace.

	$ go tool trace trace.out

## Links

No Extra links at this time.

## Code Review

[Blocking Trace](blocking_test.go) ([Go Playground](https://play.golang.org/p/e8J13dIxWe6)) 
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
