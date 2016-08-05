## Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Profiling Commands

#### CPU Profiling
```
   go test -run none -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
   go tool pprof benchmarks.test cpu.out
   
   (pprof) list algorithmOne
   (pprof) web list algorithmOne
```

#### Memory Profiling
```
    go test -run none -bench . -benchtime 3s -benchmem -memprofile mem.out
    go tool pprof -alloc_space benchmarks.test mem.out

    (pprof) list algorithmOne
    (pprof) web list algorithmOne

    -inuse_space  : Display in-use memory size
    -inuse_objects: Display in-use object counts
    -alloc_space  : Display allocated memory size
    -alloc_objects: Display allocated object counts
```

## Code Review

[Profiling](stream.go) ([Go Playground](http://play.golang.org/p/gh9AUIn4Bt)) | 
[Profiling Test](stream_test.go) ([Go Playground](https://play.golang.org/p/2lP2kMQlir))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
