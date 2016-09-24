## Blocking Profiling

Testing and Tracing allows us to see blocking profiles.

## Running a Test Based Blocking Profile

	We can get blocking profiles by running a test.

	go test -blockprofile block.out
	go tool pprof blocking.test block.out
	list TestLatency

## Running a Trace

Once you have a test established you can use the **-trace trace.out** option with the go test tool.

	go test -trace trace.out
	go tool trace trace.out

Run these command and let's explore each trace.

## Links

No Extra links at this time.

## Code Review

[Blocking Trace](blocking_test.go) ([Go Playground](https://play.golang.org/p/cjqIVeAwHz)) 
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
