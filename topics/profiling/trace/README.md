## Tracing Code

Tracing allows us to see blocking and latency profiles.

## Running A Trace

Once you have a test established you can use the **-trace file.out** option with the go test tool.

	go test -v -trace trace.out
	go tool trace trace.out

Run these command and let's explore each trace.

## Links

No Extra links at this time.

## Code Review

[Delay Trace](delay/delay_test.go) ([Go Playground](http://play.golang.org/p/_i4hBx2Pzu)) | [Latency Trace](latency/latency_test.go) ([Go Playground](http://play.golang.org/p/b50cFOkrMd)) 
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
