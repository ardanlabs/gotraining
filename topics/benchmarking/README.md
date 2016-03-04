## Benchmarking

Go has support for both testing and debugging. This includes profiling and benchmarking Go programs.

## Package Review

[Basic Benchmarking](basic/basic_test.go) ([Go Playground](https://play.golang.org/p/VVcx4Jg5E6))

[Prediction](prediction/README.md)

[Caching](caching/README.md)

[Profiling](profiling/README.md)

## Links

http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

[Profiling & Optimizing in Go / Brad Fitzpatrick](https://www.youtube.com/watch?v=xxDZuPEgbBU)

[Benchstat computes and compares statistics about benchmarks](https://github.com/rsc/benchstat)

## Exercises

### Exercise 1
Write three benchmark tests for converting an integer into a string. First use the fmt.Sprintf function, then the strconv.FormatInt function and finally the strconv.Itoa. Identify which function performs the best.

[Template](exercises/template1/bench_test.go) ([Go Playground](https://play.golang.org/p/NzqLpYD3VT))
 
[Answer](exercises/exercise1/bench_test.go) ([Go Playground](https://play.golang.org/p/C0nEumC2Pz))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
