## Testing and Benchmarking

Testing is built right into the go tools and the standard library. Testing needs to be a vital part of the development process because it can save you a tremendous amount of time throughout the life cycle of the project. Benchmarking is also a very powerful tool tied to the testing functionality. Aspect of your code can be setup to be benchmarked for performance reviews. Profiling provides a view of the interations between functions and which functions are most heavlily used.

## Notes

* The Go toolset has support for testing and benchmarking.
* The tools are very flexible and give you many options.
* Write tests in tandem with development.
* Example code serve as both a test and documentation.
* Benchmark throughout the dev, qa and release cycles.
* If performance problems are observed, profile your code to see what functions to focus on.
* The tools can interfere with each other. For example, precise memory profiling skews CPU profiles, goroutine blocking profiling affects scheduler trace, etc. Rerun tests for each needed profiling mode.

## Links

http://golang.org/pkg/testing/

http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

http://saml.rilspace.org/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof

http://golang.org/pkg/net/http/pprof/

https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs

## Code Review

[Basic Unit Test](example1/example1_test.go)

[Table Unit Test](example2/example2_test.go)

[Mocking Web Server Response](example3/example3_test.go)

[Testing Internal Endpoints](example4/example4_test.go)

[Example Test](example4/example4_example_test.go)

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
