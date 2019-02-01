## Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## Escape Analysis

* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.
* When a value is decoupled through the use of function or interface values.

## Garbage Collection History

The design of the Go GC has changed over the years:
* Go 1.0, Stop the world mark sweep collector based heavily on tcmalloc.
* Go 1.2, Precise collector, wouldn't mistake big numbers (or big strings of text) for pointers.
* Go 1.3, Fully precise tracking of all stack values.
* Go 1.4, Mark and sweep now parallel, but still stop the world.
* Go 1.5, New GC design, focusing on latency over throughput.
* Go 1.6, GC improvements, handling larger heaps with lower latency.
* Go 1.7, GC improvements, handling larger number of idle goroutines, substantial stack size fluctuation, or large package-level variables.
* Go 1.8, GC improvements, collection pauses should be significantly shorter than they were in Go 1.7, usually under 100 microseconds and often as low as 10 microseconds.

## Garbage Collection Semantics

The GC has a pacing algorithm which is used to determine when a garbage collection is to start. The algorithm depends on a feedback loop that the Pacer uses to collect information about the running application and the stress the application is putting on the heap. Stress can be defined as how fast the application is allocating all available memory on the heap within a given amount of time. Once the Pacer decides to start a collection, the amount of time to finish the collection is predetermined. This predetermined time is based on the current size of the heap, the live heap, and timing calculations about future heap usage while the collection is running.

During garbage collection there will be times when the Pacer must stop all application goroutines from running. This is called a Stop The World (STW) and there are many reasons for this. There is an initial STW that occurs at the beginning of a collection to turn the Write Barrier on. The purpose of the Write Barrier is to allow the Pacer to maintain data integrity on the heap during a collection since both GC and application goroutines will be running concurrently. In order to turn the Write Barrier on, every application goroutine running must be stopped. This STW is usually very quick, within a very small number of microseconds.

Once the write barrier is turned on, the GC goroutines begin to perform their Mark work. The Mark work consists of identifying values on the heap that are still in use. This work requires the application goroutines to share the existing CPU capacity with the GC goroutines. This is the part of the GC that makes it concurrent. The idea is to allow application work to get done at the same time GC work is getting done so the impact of the GC is minimal. During this time, the Pacer will do its best to minimize the amount of CPU the GC goroutines need to get the collection work done. There are lots of factors that go into determining how much CPU capacity the GC will use during any given collection.

During the Mark phase of the GC, there may be times when the Pacer decides to stop all the application goroutines and take all of the CPU capacity for itself. This constitutes more STW time during the collection. Reasons why the Pacer might do this could be because the application goroutines want to allocate memory on the heap at a time where the Pacer determines its not the best thing to do at that moment. Maybe too much allocation is going on and it needs to be slowed down. In these cases, the Pacer might context switch a different application goroutine to run that doesn't need the heap. It may also ask that application goroutine that wants to allocate to momentarily help with the processing of the Mark work to get it done faster.

In the end, the Pacer’s job is to start and finish a collection within the predetermined amount of time and with the least amount of STW time possible. The Pacer is constrained by the size of the heap, the size of the live heap and the number of values on the heap at the time the collection starts and the concurrent work being performed by the application goroutines. These factors play a big role in the Pacer’s ability to minimize its impact on the running program. You have to be sympathetic with these factors to help the Pacer do its job effectively.

## Links

### Pointer Mechanics

[Pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values)    
[Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html) - William Kennedy    
[Using Pointers In Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html) - William Kennedy    
[Understanding Pointers and Memory Allocation](https://www.ardanlabs.com/blog/2013/07/understanding-pointers-and-memory.html) - William Kennedy    

### Stacks

[Contiguous Stack Proposal](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)  

### Escape Analysis and Inlining

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)  
[Compiler Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)

### Garbage Collection

[The Garbage Collection Handbook](http://gchandbook.org/)  
[Tracing Garbage Collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection)  
[Go Blog - 1.5 GC](https://blog.golang.org/go15gc)  
[Go GC: Solving the Latency Problem](https://www.youtube.com/watch?v=aiv1JOfMjm0&index=16&list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm)  
[Concurrent garbage collection](http://rubinius.com/2013/06/22/concurrent-garbage-collection)  
[Go 1.5 concurrent garbage collector pacing](https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit)  
[Eliminating Stack Re-Scanning](https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md)  
[Why golang garbage-collector not implement Generational and Compact gc?](https://groups.google.com/forum/m/#!topic/golang-nuts/KJiyv2mV2pU) - Ian Lance Taylor  
[Getting to Go: The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote) - Rick Hudson  

### Static Single Assignment Optimizations

[GopherCon 2015: Ben Johnson - Static Code Analysis Using SSA](https://www.youtube.com/watch?v=D2-gaMvWfQY)  
[package ssa](https://godoc.org/golang.org/x/tools/go/ssa)    
[Understanding Compiler Optimization](https://www.youtube.com/watch?v=FnGCDLhaxKU)

### Debugging code generation

[Debugging code generation in Go](https://rakyll.org/codegen/) - JBD    

## Code Review

[Pass by Value](example1/example1.go) ([Go Playground](https://play.golang.org/p/9kxh18hd_BT))  
[Sharing data I](example2/example2.go) ([Go Playground](https://play.golang.org/p/mJz5RINaimn))  
[Sharing data II](example3/example3.go) ([Go Playground](https://play.golang.org/p/GpmPICMGMre))  
[Escape Analysis](example4/example4.go) ([Go Playground](https://play.golang.org/p/n9HijcdZ3pT))  
[Stack grow](example5/example5.go) ([Go Playground](https://play.golang.org/p/BgIKcFcZ4PO))  

### Escape Analysis Flaws

[Indirect Assignment](flaws/example1/example1_test.go)  
[Indirection Execution](flaws/example2/example2_test.go)  
[Assignment Slices Maps](flaws/example3/example3_test.go)  
[Indirection Level Interfaces](flaws/example4/example4_test.go)  
[Unknown](flaws/example5/example5_test.go)  

## Exercises

### Exercise 1

**Part A** Declare and initialize a variable of type int with the value of 20. Display the _address of_ and _value of_ the variable.

**Part B** Declare and initialize a pointer variable of type int that points to the last variable you just created. Display the _address of_ , _value of_ and the _value that the pointer points to_.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/6QYTKWzF8s8)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/qq5P9gRDHKc))

### Exercise 2

Declare a struct type and create a value of this type. Declare a function that can change the value of some field in this struct type. Display the value before and after the call to your function.

[Template](exercises/template2/template2.go) ([Go Playground](https://play.golang.org/p/nolKjrgBX-X)) |
[Answer](exercises/exercise2/exercise2.go) ([Go Playground](https://play.golang.org/p/i6utWhgDUH4))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
