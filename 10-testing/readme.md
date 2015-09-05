## Testing and Debugging

Go has support for both testing and debugging. This includes profiling and benchmarking Go programs.

## Package Review

[Testing](../10-testing/01-testing/readme.md)

[Prediction](../10-testing/02-prediction/readme.md)

[Caching](../10-testing/03-caching/readme.md)

[Godebug](../10-testing/04-godebug/readme.md)

[Profiling](../10-testing/05-profiling/readme.md)

[Stack Traces](../10-testing/06-stack_trace/readme.md)

## Exercises

### Exercise 1
Write three benchmark tests for converting an integer into a string. First use the fmt.Sprintf function, then the strconv.FormatInt function and finally the strconv.Itoa. Identify which function performs the best.

[Template](exercises/template1/bench_test.go) | 
[Answer](exercises/exercise1/bench_test.go)

___
[![Ardan Labs](../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).