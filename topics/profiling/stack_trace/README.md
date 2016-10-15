## Stack Traces

Having some basic skills in debugging Go programs can save any programmer a good amount of time trying to identify problems. I believe in logging as much information as you can, but sometimes a panic occurs and what you logged is not enough. Understanding the information in a stack trace can sometimes mean the difference between finding the bug now or needing to add more logging and waiting for it to happen again.

## Notes

* Stack traces are an important tool in debugging an application.
* The runtime should never panic so the trace is everything.
* You can see every goroutine and the call stack for each routine.
* You can see every value passed into each function on the stack.

## Links

http://www.goinggo.net/2015/01/stack-traces-in-go.html

## Code Review

[Review Stack Trace](example1/example1.go) ([Go Playground](https://play.golang.org/p/dMv5Dx92dB))  
[Packing](example2/example2.go) ([Go Playground](https://play.golang.org/p/SIiN6Y2jTR))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
