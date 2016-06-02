## Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Profiling Commands

#### CPU Profiling
```
   go test -run none -bench . -cpuprofile cpu.out
   go tool pprof profiling.test cpu.out
   
      list getValue

      10ms       10ms     36:func getValue(variable string, vars map[string]string) interface{} {
         .          .     37:
         .          .     38: // variable: "#cmd:variable_name"
         .          .     39:
         .          .     40: // Trim the # symbol from the string.
      10ms      260ms     41: value := strings.TrimLeft(variable, "#")

   go tool pprof -pdf profiling.test cpu.out > cpu.pdf
   open -a "Adobe Acrobat" mem.pdf
```

#### Memory Profiling
```
   go test -run none -bench . -memprofile mem.out -memprofilerate 1
   go tool pprof -alloc_space profiling.test mem.out

      list getValue

      .          .     36:func getValue(variable string, vars map[string]string) interface{} {
      .          .     37:
      .          .     38: // variable: "#cmd:variable_name"
      .          .     39:
      .          .     40: // Trim the # symbol from the string.
      .     6.41MB     41: value := strings.TrimLeft(variable, "#")
   
   go tool pprof -pdf -alloc_space profiling.test mem.out > mem.pdf
   open -a "Adobe Acrobat" mem.pdf

   -inuse_space  : Display in-use memory size
   -inuse_objects: Display in-use object counts
   -alloc_space  : Display allocated memory size
   -alloc_objects: Display allocated object counts
```

## Code Review

[Profiling](profiling.go) ([Go Playground](http://play.golang.org/p/45RqOFR0Ms)) | 
[Profiling Test](profiling_test.go) ([Go Playground](http://play.golang.org/p/zY3Elhibcy))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
