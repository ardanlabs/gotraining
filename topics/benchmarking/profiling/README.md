## Benchmark Profiling

Using benchmarks you can profile your programs and see exactly where your performance or memory is being taken.

## Profiling Commands

#### CPU Profiling
```
   go test -run none -bench . -cpuprofile cpu.out
   go tool pprof profiling.test cpu.out
```

#### Memory Profiling
```
   go test -run none -bench . -memprofile mem.out
   go tool pprof --alloc_space profiling.test mem.out

   -inuse_space  : Display in-use memory size
   -inuse_objects: Display in-use object counts
   -alloc_space  : Display allocated memory size
   -alloc_objects: Display allocated object counts
```

#### Profile Commands
```
   top --cum
   list profiling.getValue
```

## Profile Information

Look at how much time lines of code are taking:
```
      80ms      1.06s (flat, cum) 87.60% of Total
         .          .     32:func getValue(variable string, vars map[string]string) interface{} {
         .          .     33:
         .          .     34:	// variable: "#cmd:variable_name"
         .          .     35:
         .          .     36:	// Trim the # symbol from the string.
         .      380ms     37:	value := strings.TrimLeft(variable, "#")
         .          .     38:
         .          .     39:	// Find the first instance of the separator.
      10ms      130ms     40:	idx := strings.Index(value, ":")
         .          .     41:	if idx == -1 {
         .          .     42:		return nil
         .          .     43:	}
```

Look at how much memory lines of code are taking:
```
     125MB   437.01MB (flat, cum) 99.89% of Total
         .          .     34:func getValue(variable string, vars map[string]string) interface{} {
         .          .     35:
         .          .     36:	// variable: "#cmd:variable_name"
         .          .     37:
         .          .     38:	// Trim the # symbol from the string.
         .      136MB     39:	value := strings.TrimLeft(variable, "#")
         .          .     40:
         .          .     41:	// Find the first instance of the separator.
```

## Code Review

[Profiling](profiling.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
