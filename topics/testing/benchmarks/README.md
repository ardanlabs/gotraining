## Benchmarking

Go has support for testing the performance of your code.

## Package Review

[Basic Benchmarking](basic/basic_test.go) ([Go Playground](https://play.golang.org/p/Kwqe4zYpMi))  
[Sub Benchmarks](sub/sub_test.go) ([Go Playground](https://play.golang.org/p/q4qV9BdH4R))  
[Prediction](prediction/README.md)  
[Caching](caching/README.md)  

_Look at the profiling topic to learn more about using benchmarks to [profile](../profiling) code._

## Links

http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go  
[Profiling & Optimizing in Go / Brad Fitzpatrick](https://www.youtube.com/watch?v=xxDZuPEgbBU)  
[Benchstat computes and compares statistics about benchmarks](https://github.com/rsc/benchstat)  
[Advanced Testing Concepts for Go 1.7](https://speakerdeck.com/mpvl/advanced-testing-concepts-for-go-1-dot-7) - Marcel van Lohuizen  

## Exercises

### Exercise 1
Write three benchmark tests for converting an integer into a string. First use the fmt.Sprintf function, then the strconv.FormatInt function and finally the strconv.Itoa. Identify which function performs the best.

[Template](exercises/template1/bench_test.go) ([Go Playground](http://play.golang.org/p/do3XfkNqRt)) | 
[Answer](exercises/exercise1/bench_test.go) ([Go Playground](http://play.golang.org/p/ttqLnSM2q_))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
