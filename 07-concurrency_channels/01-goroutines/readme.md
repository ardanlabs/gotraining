## Goroutines - Concurrency and Channels

Goroutines are functions that are created and scheduled to be run indenpently. Goroutines are multiplexed against a shared thread that is own by context. The scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independently.
* The scheduler uses a context that owns an OS thread and goroutine run queue.
* We must always maintain an account of running goroutines and shutdown cleanly.
* Concurrency is not parallelism.
	* Concurrency is about dealing with lots of things at once.
	* Parallelism is about doing lots of things at once.

## Documentation

[Scheduler Diagrams](documentation/scheduler.md)

## Links

http://blog.golang.org/advanced-go-concurrency-patterns

http://blog.golang.org/context

http://blog.golang.org/concurrency-is-not-parallelism

http://talks.golang.org/2013/distsys.slide

http://www.goinggo.net/2014/01/concurrency-goroutines-and-gomaxprocs.html

## Code Review

[Goroutines and concurrency](example1/example1.go) ([Go Playground](http://play.golang.org/p/LI0TEB6PMf))

[Goroutine time slicing](example2/example2.go) ([Go Playground](https://play.golang.org/p/AzDUqwkM58))

[Goroutines and parallelism](example3/example3.go) ([Go Playground](http://play.golang.org/p/cqsHoPD30n))

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. Once that counts up to 100 from 0 and one that counts down to 0 from 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/H-h1cbBW3B)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/mB4QslSNoA))

___
[![Ardan Labs](../../00-slides/images/ggt_logo.png)](http://www.ardanlabs.com)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).
