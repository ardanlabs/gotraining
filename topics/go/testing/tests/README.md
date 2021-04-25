## Testing

Testing is built right into the go tools and the standard library. Testing needs to be a vital part of the development process because it can save you a tremendous amount of time throughout the life cycle of the project. Benchmarking is also a very powerful tool tied to the testing functionality. Aspect of your code can be setup to be benchmarked for performance reviews. Profiling provides a view of the interations between functions and which functions are most heavily used.

## Notes

* The Go toolset has support for testing and benchmarking.
* The tools are very flexible and give you many options.
* Write tests in tandem with development.
* Example code serve as both a test and documentation.
* Benchmark throughout the dev, qa and release cycles.
* If performance problems are observed, profile your code to see what functions to focus on.
* The tools can interfere with each other. For example, precise memory profiling skews CPU profiles, goroutine blocking profiling affects scheduler trace, etc. Rerun tests for each needed profiling mode.

## Quotes

_"A unit test is a test of behavior whose success or failure is wholly determined by the correctness of the test and the correctness of the unit under test." - Kevlin Henney_

## Links

[The deep synergy between testability and good design](https://www.youtube.com/watch?reload=9&feature=share&v=4cVZvoFGJTU&app=desktop) - Michael Feathers  
[testing package](http://golang.org/pkg/testing/)    
[How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) - Dave Cheney    
[Profiling and creating call graphs for Go programs with "go tool pprof"](http://saml.rilspace.com/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof) - Samuel Lampa    
[pprof package](https://golang.org/pkg/net/http/pprof/)    
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov    
https://github.com/dvyukov/go-fuzz  
[Go Dynamic Tools](https://talks.golang.org/2015/dynamic-tools.slide#1) - Dmitry Vyukov    
[Automated Testing with go-fuzz](https://vimeo.com/141698770) - Filippo Valsorda    
[Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c#.b2m3nziyb) - Ben Johnson  
[Advanced Testing Concepts for Go 1.7](https://speakerdeck.com/mpvl/advanced-testing-concepts-for-go-1-dot-7) - Marcel van Lohuizen  
[Parallelize your table-driven tests](https://rakyll.org/parallelize-test-tables/) - JBD     
[Advanced Testing with Go - Video](https://www.youtube.com/shared?ci=LARb45o5TpA) - Mitchell Hashimoto  
[Advanced Testing with Go - Deck](https://speakerdeck.com/mitchellh/advanced-testing-with-go) - Mitchell Hashimoto  
[The tragedy of 100% code coverage](http://labs.ig.com/code-coverage-100-percent-tragedy) - Daniel Lebrero's  

## Code Review

[Basic Unit Test](example1/example1_test.go) ([Go Playground](https://play.golang.org/p/F7kXmSfr7AE))  
[Table Unit Test](example2/example2_test.go) ([Go Playground](https://play.golang.org/p/1a2u8omEqrX))  
[Mocking Web Server Response](example3/example3_test.go) ([Go Playground](https://play.golang.org/p/SILnu117hak))  
[Testing Internal Endpoints](example4/handlers/handlers_test.go) ([Go Playground](https://play.golang.org/p/CSK7SZEeWf3))  
[Example Test](example4/handlers/handlers_example_test.go) ([Go Playground](https://play.golang.org/p/rE0DRliZH9t))  
[Sub Tests](example5/example5_test.go) ([Go Playground](https://play.golang.org/p/7PrkFU-qVdY))  

_Look at the profiling topic to learn more about using test to [profile](../profiling) code._

## Coverage

Making sure your tests cover as much of your code as possible is critical. Go's testing tool allows you to create a profile for the code that is executed during all the tests and see a visual of what is and is not covered.

	go test -coverprofile cover.out
	go tool cover -html=cover.out

![figure1](testing_coverage.png)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
