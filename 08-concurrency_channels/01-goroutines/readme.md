## Goroutines - Concurrency and Channels

Goroutines are functions that are created and scheduled to be run independently. Goroutines are multiplexed against a shared thread that is own by context. The scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independently.
* The scheduler uses a context that owns an OS thread and goroutine run queue.
* We must always maintain an account of running goroutines and shutdown cleanly.
* Concurrency is not parallelism.
	* Concurrency is about dealing with lots of things at once.
	* Parallelism is about doing lots of things at once.

## Documentation

[Scheduler Diagrams](documentation/scheduler.md)

## Multi Processor and Thread Memory Access Issues
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (32:46-35:40)](https://youtu.be/WDIkqP4JbkE?t=1966)

## Notes

* Cache lines (64 bytes) are moved in and out of the caches.
* Two or more processors or cores may cache the same memory, cores need to know about activity.
* Cache lines can be flagged as dirty and reloaded between the caches for a processor or core.
* Leveraging the stack can give us exclusivity.
* When a thread is migrated from one processor to another, all the cache lines have to be moved.
* Swapping threads between processors can mean the swapping of cache.
* Leveraging one thread against a consistent data set can provide better performance.
* Hyper-Threading lets the processor work more than one thread at a time.
* Even with HT, only one thread is executing at a time.

## Links

http://blog.golang.org/advanced-go-concurrency-patterns

http://blog.golang.org/context

http://blog.golang.org/concurrency-is-not-parallelism

http://talks.golang.org/2013/distsys.slide

http://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html

http://www.akkadia.org/drepper/cpumemory.pdf

http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips

## Code Review

[Goroutines and concurrency](example1/example1.go) ([Go Playground](https://play.golang.org/p/ki1woWvmzW))

[Goroutine time slicing](example2/example2.go) ([Go Playground](https://play.golang.org/p/SvJj6T4Jhi))

[Goroutines and parallelism](example3/example3.go) ([Go Playground](https://play.golang.org/p/kz65m4PHmC))

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. Once that counts up to 100 from 0 and one that counts down to 0 from 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/KlKIYq9s_3)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/pzIjQhIJ5J))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
