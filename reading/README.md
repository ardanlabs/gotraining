## Important Reading

There is a ton of links to extra reading and videos through out the material. Occasionally we run into material that we feel is important and don't have an exact place to put it. This is a section place for that material.

### Books

[Go In Action](https://www.manning.com/books/go-in-action) - Kennedy, Ketelsen, St. Martin  
[Learn Go](https://www.manning.com/books/learn-go) - Nathan Youngman

### Blogs

[Going Go](http://www.goinggo.net) - William Kennedy  
[The acme of foolishness](http://dave.cheney.net) - Dave Cheney  
[The Go Blog](https://blog.golang.org) - Language Team  
[Ukiah Smith](https://ukiahsmith.com) - Ukiah Smith  
[Mechanical Sympathy](http://mechanical-sympathy.blogspot.com) - Martin Thompson  
[Go Documentation References](https://golang.org/doc/#references) - Language Team

### Interviews

[An Interview with Brian Kernighan](http://www.cs.cmu.edu/~mihaib/kernighan-interview/index.html)  
[Interview with Dennis Ritchie, Bjarne Stroustrup, and James Gosling](http://www.gotw.ca/publications/c_family_interview.htm)

### History

[The Evolution of Go](https://www.youtube.com/watch?v=0ReKdcpNyQg) - Robert Griesmer  
[Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article) - Rob Pike  

[A Very Brief History of Computing, 1948-2015](http://www.gresham.ac.uk/lectures-and-events/a-very-brief-history-of-computing-1948-2015) - Martyn Thomas  
[The Rise and Fall of Minicomputers](http://ethw.org/Rise_and_Fall_of_Minicomputers) - Gordon Bell  
[After Moore's Law - Economist](http://www.economist.com/technology-quarterly/2016-03-12/after-moores-law) - Tim Cross  
[A Crash Course in Modern Hardware](http://www.infoq.com/presentations/click-crash-course-modern-hardware#.VwoB63sl6no.twitter) -  Cliff Click  

[The Future of Programming](https://www.youtube.com/watch?v=9Xy3QC7yxJw) - Uncle Bob

### Vendoring

[How one developer just broke Node, Babel and thousands of projects in 11 lines of JavaScript](http://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos)

### Profiling and Debugging

[Profiling & Optimizing in Go](https://www.youtube.com/watch?v=xxDZuPEgbBU) - Brad Fitzpatrick  
[Go Dynamic Tools](https://www.youtube.com/watch?v=a9xrxRsIbSU) - Dmitry Vyukov  
[How NOT to Measure Latency](https://www.youtube.com/watch?v=lJ8ydIuPFeU&feature=youtu.be) - Gil Tene  
[Go Performance Tales](http://jmoiron.net/blog/go-performance-tales) - Jason Moiron  
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov  
[Reduce allocation in Go code](https://methane.github.io/2015/02/reduce-allocation-in-go-code) - Python Bytes

### Interfaces and Composition

[Design Guidelines](design_guidelines.md)  
[Composition](http://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy  
[Interface Pollution](https://medium.com/@rakyll/interface-pollution-in-go-7d58bccec275) - Burcu Dogan

### Buffer Bloat - 2011

[Bufferbloat: Dark Buffers in the Internet](https://www.youtube.com/watch?v=qbIozKVz73g) - Jim Gettys  
[Buffer Bloat Videos](http://www.bufferbloat.net/projects/cerowrt/wiki/Bloat-videos)

### Linux

[The Definitive Guide to Linux System Calls](http://blog.packagecloud.io/eng/2016/04/05/the-definitive-guide-to-linux-system-calls/#hardware-and-software)  
[Linux Performance Analysis in 60,000 Milliseconds](http://techblog.netflix.com/2015/11/linux-performance-analysis-in-60s.html) - Netflix  

### Docker

[How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Kulshekhar Kabra  

### Misc

[Shrink Your Go Binaries With This One Weird Trick](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick) - Filippo Valsorda  
[Write High Performance Go](http://go-talks.appspot.com/github.com/davecheney/presentations/writing-high-performance-go.slide) - Dave Cheney  

## Mechanical Sympathy

Much of this content can be found under [arrays](../topics/arrays) and [data races](../topics/data_race). We feel this content is so important it has been moved out to this general materials page.

#### CPU Caches

[CPU Caches and Why You Care - Video](https://www.youtube.com/watch?v=WDIkqP4JbkE) - Scott Meyers  
[CPU Caches and Why You Care - Deck](http://www.aristeia.com/TalkNotes/codedive-CPUCachesHandouts.pdf) - Scott Meyers  
[Mythbusting Modern Hardware to Gain 'Mechanical Sympathy`](https://www.youtube.com/watch?v=MC1EKLQ2Wmg) - Martin Thompson  
[What Every Programmer Should Know About Memory](http://www.akkadia.org/drepper/cpumemory.pdf) - Ulrich Drepper  
[How CPU Caches Work and Why](http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips) - Joel Hruska  
[Modern Microprocessors A 90 Minute Guide](http://www.lighterra.com/papers/modernmicroprocessors) - Jason Robert Carey Patterson  
[Memory part 2: CPU caches](http://lwn.net/Articles/252125) - Ulrich Drepper  
[The Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm) - Herb Sutter  
[Data Center Computers: Modern Challenges in CPU Design](https://m.youtube.com/watch?feature=youtu.be&v=QBu2Ae8-8LM) - Dick Sites  
[Wirth's Law](https://en.wikipedia.org/wiki/Wirth%27s_law) - Wikipedia  
[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206) - Herb Sutter  

#### Data-Oriented Design

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton  
[Pitfalls of OOP](http://harmful.cat-v.org/software/OO_programming/_pdf/Pitfalls_of_Object_Oriented_Programming_GCAP_09.pdf) - Tony Albrecht  
[Why you should avoid Linked Lists](https://www.youtube.com/watch?v=YQs6IC-vgmo) - Bjarne Stroustrup  
[Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth

#### Operating Systems and Virtualization

[The Linux Scheduler: a Decade of Wasted Cores](http://www.ece.ubc.ca/~sasha/papers/eurosys16-final29.pdf)  
[The Cost of Virtualization](http://queue.acm.org/detail.cfm?id=1348591) - Ulrich Drepper  

## Runtime

This content can be found under [pointers](../topics/pointers). We feel this content is so important it has been moved out to this general materials page.

#### Stacks

[Contiguous Stack Proposal](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)

#### Escape Analysis and Inlining

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw) - Dmitry Vyukov  
[Compiler Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)

#### Garbage Collection

[Tracing Garbage Collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection) - Wikipedia  
[Go Blog - 1.5 GC](https://blog.golang.org/go15gc)  
[Go GC: Solving the Latency Problem](https://www.youtube.com/watch?v=aiv1JOfMjm0&index=16&list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm) - Rick Hudson  

#### Single Static Assignment Optimizations

[GopherCon 2015: Static Code Analysis Using SSA](https://www.youtube.com/watch?v=D2-gaMvWfQY) - Ben Johnson  
https://github.com/golang/go/blob/dev.ssa/src/cmd/compile/internal/ssa/compile.go#L83  
https://godoc.org/golang.org/x/tools/go/ssa  
[Understanding Compiler Optimization](https://www.youtube.com/watch?v=FnGCDLhaxKU) - Chandler Carruth  

## Release Notes

[Go 1.6 Release Notes](https://golang.org/doc/go1.6) - 2016/02/17  
[Go 1.5 Release Notes](https://golang.org/doc/go1.5) - 2015/08/19  
[Go 1.4 Release Notes](https://golang.org/doc/go1.4) - 2014/12/10  
[Go 1.3 Release Notes](https://golang.org/doc/go1.3) - 2014/06/18  
[Go 1.2 Release Notes](https://golang.org/doc/go1.2) - 2013/12/01  
[Go 1.1 Release Notes](https://golang.org/doc/go1.1) - 2013/05/13  
[Go 1.0 Release Notes](https://golang.org/doc/go1.0) - 2012/03/28
