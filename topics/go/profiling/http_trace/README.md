## HTTP Tracing

In Go 1.7 the HTTP tracing package was introduced to facilitate the gathering of fine-grained information throughout the lifecycle of an HTTP client request. Support for HTTP tracing is provided by the net/http/httptrace package. The collected information can be used for debugging latency issues, service monitoring, writing adaptive systems, and more.

## Notes

The httptrace package provides a number of hooks to gather information during an HTTP round trip about a variety of events. These events include:

* Connection creation
* Connection reuse
* DNS lookups
* Writing the request to the wire
* Reading the response

## Links

[Introducing HTTP Tracing](https://blog.golang.org/http-tracing) - Jaana Burcu Dogan  

## Code Review

[Tracing events](example1/example1.go) ([Go Playground](https://play.golang.org/p/9Y3Y3gfgb3j))  
[Tracing with http.Client](example2/example2.go) ([Go Playground](https://play.golang.org/p/Qh6DD-VCnHJ))  

## Exercises

### Exercise 1

TBD
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
