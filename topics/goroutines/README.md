## Goroutines

Goroutines are functions that are created and scheduled to be run independently. Goroutines are multiplexed against a shared thread that is owned by logical processor. The scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independently.
* The scheduler uses a logical processor that owns an OS thread and goroutine run queue.
* We must always maintain an account of running goroutines and shutdown cleanly.
* Concurrency is not parallelism.
	* Concurrency is about dealing with lots of things at once.
	* Parallelism is about doing lots of things at once.

## Diagrams

### How the scheduler works.

![Ardan Labs](scheduler.png)

### Difference between concurrency and parallelism.

![Ardan Labs](parallel.png)

## Links

http://blog.golang.org/advanced-go-concurrency-patterns  
http://blog.golang.org/context  
http://blog.golang.org/concurrency-is-not-parallelism  
http://talks.golang.org/2013/distsys.slide  
[Go 1.5 GOMAXPROCS Default](https://docs.google.com/document/d/1At2Ls5_fhJQ59kDK2DFVhFu3g5mATSXqqV5QrxinasI/edit)  
http://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html  
[The Linux Scheduler: a Decade of Wasted Cores](http://www.ece.ubc.ca/~sasha/papers/eurosys16-final29.pdf)

## Code Review

[Goroutines and concurrency](example1/example1.go) ([Go Playground](http://play.golang.org/p/eV4l2JqLZL))  
[Goroutine time slicing](example2/example2.go) ([Go Playground](http://play.golang.org/p/8NwVeZG6IB))  
[Goroutines and parallelism](example3/example3.go) ([Go Playground](http://play.golang.org/p/Bvk_ZwxJVh))

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. One that counts up to 100 from 0 and one that counts down to 0 from 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/xv15yVM2yE)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/87OanzQcg4))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
