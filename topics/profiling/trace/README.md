## Tracing / Blocking Profiling

Testing and Tracing allows us to see blocking and latency profiles.

## Running A Test

	We can get blocking profiles by running a test.

	cd delay
	go test -blockprofile block.out
	go tool pprof delay.test block.out
	list TestDelay

	cd latency
	go test -blockprofile block.out
	go tool pprof latency.test block.out
	list TestLatency

## Running A Trace

Once you have a test established you can use the **-trace file.out** option with the go test tool.

	go test -trace trace.out
	go tool trace trace.out

Run these command and let's explore each trace.

## Links

No Extra links at this time.

## Code Review

[Delay Trace](delay/delay_test.go) ([Go Playground](http://play.golang.org/p/_i4hBx2Pzu)) | [Latency Trace](latency/latency_test.go) ([Go Playground](http://play.golang.org/p/b50cFOkrMd)) 
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
