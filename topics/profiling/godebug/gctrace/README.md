## Memory Tracing

There is no way to identify specifically in the code where a leak is occurring. We can validate if a memory leak is present and which functions or methods are producing the most allocations.

## GODEBUG

To validate any sort of potential memory problems, including memory leaks, use the GODEBUG environmental variable. 

Setting **gctrace=1** causes the garbage collector to emit a single line to standard error at each collection, summarizing the amount of memory collected and the length of the pause. Setting gctrace=2 emits the same summary but also repeats each collection. The format of this line is subject to change:

    export GODEBUG=gctrace=1

    gc # @#s #%: #+...+# ms clock, #+...+# ms cpu, #->#-># MB, # MB goal, # P

Where the fields are as follows:

    gc #        the GC number, incremented at each GC
    @#s         time in seconds since program start
    #%          percentage of time spent in GC since program start
    #+...+#     wall-clock/CPU times for the phases of the GC
    #->#-># MB  heap size at GC start, at GC end, and live heap
    # MB goal   goal heap size
    # P         number of processors used

**wall-clock** time is a measure of the real time that elapses from start to end, including time that passes due to programmed (artificial) delays or waiting for resources to become available.
https://en.wikipedia.org/wiki/Wall-clock_time

**CPU time** (or process time) is the amount of time for which a central processing unit (CPU) was used for processing instructions of a computer program or operating system, as opposed to, for example, waiting for input/output (I/O) operations or entering low-power (idle) mode.
https://en.wikipedia.org/wiki/CPU_time

You can get more details by adding the **gcpacertrace=1** flag. This causes the garbage collector to print information about the internal state of the concurrent pacer.

    export GODEBUG=gctrace=1,gcpacertrace=1

Sample output:

    gc 5 @0.071s 0%: 0.018+0.46+0.071 ms clock, 0.14+0/0.38/0.14+0.56 ms cpu, 29->29->29 MB, 30 MB goal, 8 P
    pacer: sweep done at heap size 29MB; allocated 0MB of spans; swept 3752 pages at +6.183550e-004 pages/byte
    pacer: assist ratio=+1.232155e+000 (scan 1 MB in 70->71 MB) workers=2+0
    pacer: H_m_prev=30488736 h_t=+2.334071e-001 H_T=37605024 h_a=+1.409842e+000 H_a=73473040 h_g=+1.000000e+000 H_g=60977472 u_a=+2.500000e-001 u_g=+2.500000e-001 W_a=308200 goalΔ=+7.665929e-001 actualΔ=+1.176435e+000 u_a/u_g=+1.000000e+000

Notes:

In C++, a memory leak is memory you have lost a reference to.  
In Go, a memory leak is memory you retain a reference to.

## Running a GODEBUG GC Trace

    go build
    GODEBUG=gctrace=1 ./gctrace

    gc 1 @0.009s 1%: 0.059+0.17+0.005+0.24+0.12 ms clock, 0.17+0.17+0+0/0.36/0.067+0.38 ms cpu, 5->5->3 MB, 4 MB goal, 8 P
    gc 2 @0.017s 1%: 0.037+0.096+0.098+0.21+0.086 ms clock, 0.22+0.096+0+0.10/0.31/0.091+0.51 ms cpu, 8->8->7 MB, 7 MB goal, 8 P
    gc 3 @0.032s 1%: 0.020+0.16+0.007+0.25+0.090 ms clock, 0.14+0.16+0+0/0.20/0.27+0.63 ms cpu, 17->17->14 MB, 14 MB goal, 8 P
    gc 4 @0.066s 0%: 0.029+0.17+0.074+0.48+0.10 ms clock, 0.20+0.17+0+0/0.42/0.26+0.76 ms cpu, 35->35->29 MB, 29 MB goal, 8 P

    gc 1        : First GC run since program started.
    @0.009s     : Nine milliseconds since the program started.
    1%          : One percent of the programs time has been spent in GC.
    
    // wall-clock
    0.059ms     : **STW** Sweep termination.
    0.17ms      : Mark/Scan - Assist Time (GC performed in line with allocation).
    0.005ms     : Mark/Scan - Background GC time.
    0.24ms      : Mark/Scan - Idle GC time.
    0.12ms      : **STW** Mark termination.

    // CPU time
    0.17ms      : **STW** Sweep termination.
    0.17+0+0ms  : Mark/Scan - Assist Time (GC performed in line with allocation).
    0.36ms      : Mark/Scan - Background GC time.
    0.067ms     : Mark/Scan - Idle GC time.
    0.38ms      : **STW** Mark termination.

    5MB         : Heap size at GC start.
    5MB         : Heap size at GC end.
    3MB         : Live Heap.
    4MB         : Goal heap size.
    8P          : Number of logical processors. 

## Links

https://golang.org/pkg/runtime  
https://www.hakkalabs.co/articles/finding-memory-leaks-go-programs  
[http://golang.org/pkg/runtime/](http://golang.org/pkg/runtime/)  
[http://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector](Visualising the Go garbage collector)    
[http://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables](Tour of Go's env variables)    
[https://deferpanic.com/blog/understanding-golang-memory-usage](Understanding Go memory usage)  
[GC Runtime Source Code](https://golang.org/src/runtime/mgc.go)  

## Code Review

[Memory Trace](trace.go) ([Go Playground](https://play.golang.org/p/ty-4EwbuH_))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
