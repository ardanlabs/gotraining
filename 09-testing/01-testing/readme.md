## Testing and Benchmarking

Testing is built right into the go tools and the standard library. Testing needs to be a vital part of the development process because it can save you a tremendous amount of time throughout the life cycle of the project. Benchmarking is also a very powerful tool tied to the testing functionality. Aspect of your code can be setup to be benchmarked for performance reviews. Profiling provides a view of the interations between functions and which functions are most heavlily used.

## Notes

* The Go toolset has support for testing and benchmarking.
* The tools are very flexible and give you many options.
* Write tests in tandem with development.
* "Example" tests serve as both a test and documentation.
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

The sample program implements support for accessing a MongoDB database from MongoLab. The program implements two different find calls that return documents that are unmarshaled into user defined types.

### Sample App

[example1.go](example1/example1.go)

[buoy/buoy.go](example1/buoy/buoy.go)

[mongodb/mongodb.go](example1/mongodb/mongodb.go)

### Tests

[Standard tests for testing calls to MongoDB](example1/tests/example1_test.go)

[Table tests that perform multiple calls with different ids](example1/tests/example1_table_test.go)

## Exercises

### Exercise 1
Write three benchmark tests for converting an integer into a string. First use the fmt.Sprintf function, then the strconv.FormatInt function and finally the strconv.Itoa. Identify which function performs the best.

[Template](exercises/template1/bench_test.go) | 
[Answer](exercises/exercise1/bench_test.go)

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
